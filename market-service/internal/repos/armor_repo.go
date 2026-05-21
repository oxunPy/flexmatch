package repos

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"market-service/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func BuyArmor(pool *pgxpool.Pool, armorID string, playerID int64) (*models.PlayerArmor, error) {
	ctx := context.Background()

	exists, err := armorExists(ctx, pool, armorID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("armor not found: %s", armorID)
	}

	var playerArmor models.PlayerArmor
	err = pool.QueryRow(ctx, `
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

func SellArmor(pool *pgxpool.Pool, armorID string, playerID int64) (bool, error) {
	commandTag, err := pool.Exec(context.Background(), `
		DELETE FROM player_armors
		WHERE player_id = $1 AND armor_id = $2::uuid
	`, playerID, armorID)
	if err != nil {
		return false, fmt.Errorf("sell armor: %w", err)
	}

	return commandTag.RowsAffected() > 0, nil
}

func CreateArmor(pool *pgxpool.Pool, armor models.Armor) (*models.Armor, error) {
	attr, err := json.Marshal(armor.Attr)
	if err != nil {
		return nil, fmt.Errorf("encode armor attr: %w", err)
	}

	var created models.Armor
	var rawAttr []byte
	err = pool.QueryRow(context.Background(), `
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

func GetAllArmors(pool *pgxpool.Pool) ([]*models.Armor, error) {
	rows, err := pool.Query(context.Background(), `
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

func armorExists(ctx context.Context, pool *pgxpool.Pool, armorID string) (bool, error) {
	var exists bool
	err := pool.QueryRow(ctx, `
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
