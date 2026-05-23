package repos

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"market-service/internal/database"

	"market-service/internal/models"

	"github.com/jackc/pgx/v5"
)

type ArmorRepo struct {
	storage *database.PostgresStorage
}

func NewArmorRepo(storage *database.PostgresStorage) *ArmorRepo {
	return &ArmorRepo{
		storage: storage,
	}
}

func (r *ArmorRepo) BuyArmor(armorID string, playerID int64) (*models.PlayerArmor, error) {
	ctx := context.Background()

	exists, err := r.armorExists(ctx, armorID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("armor not found: %s", armorID)
	}

	var playerArmor models.PlayerArmor
	err = r.storage.QueryRow(ctx, `
		INSERT INTO player_armors (player_id, armor_id)
		VALUES ($1, $2::uuid)
		ON CONFLICT (player_id, armor_id) DO UPDATE SET armor_id = EXCLUDED.armor_id
		RETURNING player_id, armor_id::text
	`, playerID, armorID).Scan(&playerArmor.PlayerID, &playerArmor.ArmorID)
	if err != nil {
		return nil, fmt.Errorf("buy armor: %w", err)
	}

	return &playerArmor, nil
}

func (r *ArmorRepo) SellArmor(armorID string, playerID int64) (bool, error) {
	commandTag, err := r.storage.Exec(context.Background(), `
		DELETE FROM player_armors
		WHERE player_id = $1 AND armor_id = $2::uuid
	`, playerID, armorID)
	if err != nil {
		return false, fmt.Errorf("sell armor: %w", err)
	}

	return commandTag.RowsAffected() > 0, nil
}

func (r *ArmorRepo) CreateArmor(armor models.Armor) (*models.Armor, error) {
	attr, err := json.Marshal(armor.Attr)
	if err != nil {
		return nil, fmt.Errorf("encode armor attr: %w", err)
	}

	var created models.Armor
	var rawAttr []byte
	err = r.storage.QueryRow(context.Background(), `
		INSERT INTO armors (name, description, cost, attr)
		VALUES ($1, $2, $3, $4)
		RETURNING id::text, name, description, cost, attr, created, updated
	`, armor.Name, armor.Desc, armor.Cost, attr).Scan(
		&created.ID,
		&created.Name,
		&created.Desc,
		&created.Cost,
		&rawAttr,
		&created.Created,
		&created.Updated,
	)
	if err != nil {
		return nil, fmt.Errorf("create armor: %w", err)
	}

	if len(rawAttr) > 0 {
		if err := json.Unmarshal(rawAttr, &created.Attr); err != nil {
			return nil, fmt.Errorf("decode armor attr: %w", err)
		}
	}

	return &created, nil
}

func (r *ArmorRepo) GetAllArmors() ([]*models.Armor, error) {
	rows, err := r.storage.Query(context.Background(), `
		SELECT id::text, name, description, cost, attr, created, updated
		FROM armors
		ORDER BY created DESC
	`)
	if err != nil {
		return nil, fmt.Errorf("get armors: %w", err)
	}
	defer rows.Close()

	armors := make([]*models.Armor, 0)
	for rows.Next() {
		var armor models.Armor
		var rawAttr []byte
		if err := rows.Scan(
			&armor.ID,
			&armor.Name,
			&armor.Desc,
			&armor.Cost,
			&rawAttr,
			&armor.Created,
			&armor.Updated,
		); err != nil {
			return nil, fmt.Errorf("scan armor: %w", err)
		}

		if len(rawAttr) > 0 {
			if err := json.Unmarshal(rawAttr, &armor.Attr); err != nil {
				return nil, fmt.Errorf("decode armor attr: %w", err)
			}
		}

		armors = append(armors, &armor)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate armors: %w", err)
	}

	return armors, nil
}

func (r *ArmorRepo) armorExists(ctx context.Context, armorID string) (bool, error) {
	var exists bool
	err := r.storage.QueryRow(ctx, `
		SELECT EXISTS(SELECT 1 FROM armors WHERE id = $1::uuid)
	`, armorID).Scan(&exists)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}

		return false, fmt.Errorf("check armor exists: %w", err)
	}

	return exists, nil
}
