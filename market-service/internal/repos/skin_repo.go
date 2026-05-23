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

type SkinRepo struct {
	storage *database.PostgresStorage
}

func NewSkinRepo(storage *database.PostgresStorage) *SkinRepo {
	return &SkinRepo{
		storage: storage,
	}
}

func (r *SkinRepo) BuySkin(skinID string, playerID int64) (*models.PlayerSkin, error) {
	ctx := context.Background()

	exists, err := r.skinExists(ctx, skinID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("skin not found: %s", skinID)
	}

	var playerSkin models.PlayerSkin
	err = r.storage.QueryRow(ctx, `
		INSERT INTO player_skins (player_id, skin_id)
		VALUES ($1, $2::uuid)
		ON CONFLICT (player_id, skin_id) DO UPDATE SET skin_id = EXCLUDED.skin_id
		RETURNING player_id, skin_id::text
	`, playerID, skinID).Scan(&playerSkin.PlayerID, &playerSkin.SkinID)
	if err != nil {
		return nil, fmt.Errorf("buy skin: %w", err)
	}

	return &playerSkin, nil
}

func (r *SkinRepo) SellSkin(skinID string, playerID int64) (bool, error) {
	commandTag, err := r.storage.Exec(context.Background(), `
		DELETE FROM player_skins
		WHERE player_id = $1 AND skin_id = $2::uuid
	`, playerID, skinID)
	if err != nil {
		return false, fmt.Errorf("sell skin: %w", err)
	}

	return commandTag.RowsAffected() > 0, nil
}

func (r *SkinRepo) CreateSkin(skin models.Skin) (*models.Skin, error) {
	attr, err := json.Marshal(skin.Attr)
	if err != nil {
		return nil, fmt.Errorf("encode skin attr: %w", err)
	}

	var created models.Skin
	var rawAttr []byte
	err = r.storage.QueryRow(context.Background(), `
		INSERT INTO skins (name, cost, attr)
		VALUES ($1, $2, $3)
		RETURNING id::text, name, cost, attr, created, updated
	`, skin.Name, skin.Cost, attr).Scan(
		&created.ID,
		&created.Name,
		&created.Cost,
		&rawAttr,
		&created.Created,
		&created.Updated,
	)
	if err != nil {
		return nil, fmt.Errorf("create skin: %w", err)
	}

	if len(rawAttr) > 0 {
		if err := json.Unmarshal(rawAttr, &created.Attr); err != nil {
			return nil, fmt.Errorf("decode skin attr: %w", err)
		}
	}

	return &created, nil
}

func (r *SkinRepo) GetAllSkins() ([]*models.Skin, error) {
	rows, err := r.storage.Query(context.Background(), `
		SELECT id::text, name, cost, attr, created, updated
		FROM skins
		ORDER BY created DESC
	`)
	if err != nil {
		return nil, fmt.Errorf("get skins: %w", err)
	}
	defer rows.Close()

	skins := make([]*models.Skin, 0)
	for rows.Next() {
		var skin models.Skin
		var rawAttr []byte
		if err := rows.Scan(
			&skin.ID,
			&skin.Name,
			&skin.Cost,
			&rawAttr,
			&skin.Created,
			&skin.Updated,
		); err != nil {
			return nil, fmt.Errorf("scan skin: %w", err)
		}

		if len(rawAttr) > 0 {
			if err := json.Unmarshal(rawAttr, &skin.Attr); err != nil {
				return nil, fmt.Errorf("decode skin attr: %w", err)
			}
		}

		skins = append(skins, &skin)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate skins: %w", err)
	}

	return skins, nil
}

func (r *SkinRepo) skinExists(ctx context.Context, skinID string) (bool, error) {
	var exists bool
	err := r.storage.QueryRow(ctx, `
		SELECT EXISTS(SELECT 1 FROM skins WHERE id = $1::uuid)
	`, skinID).Scan(&exists)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}

		return false, fmt.Errorf("check skin exists: %w", err)
	}

	return exists, nil
}
