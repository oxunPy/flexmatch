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

func BuyWeapon(pool *pgxpool.Pool, weaponID string, playerID int64) (*models.PlayerWeapon, error) {
	ctx := context.Background()

	exists, err := weaponExists(ctx, pool, weaponID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("weapon not found: %s", weaponID)
	}

	var playerWeapon models.PlayerWeapon
	err = pool.QueryRow(ctx, `
		INSERT INTO player_weapons (player_id, weapon_id)
		VALUES ($1, $2::uuid)
		ON CONFLICT (player_id, weapon_id) DO UPDATE SET weapon_id = EXCLUDED.weapon_id
		RETURNING player_id, weapon_id::text
	`, playerID, weaponID).Scan(&playerWeapon.PlayerID, &playerWeapon.WeaponID)
	if err != nil {
		return nil, fmt.Errorf("buy weapon: %w", err)
	}

	return &playerWeapon, nil
}

func SellWeapon(pool *pgxpool.Pool, weaponID string, playerID int64) (bool, error) {
	commandTag, err := pool.Exec(context.Background(), `
		DELETE FROM player_weapons
		WHERE player_id = $1 AND weapon_id = $2::uuid
	`, playerID, weaponID)
	if err != nil {
		return false, fmt.Errorf("sell weapon: %w", err)
	}

	return commandTag.RowsAffected() > 0, nil
}

func CreateWeapon(pool *pgxpool.Pool, weapon models.Weapon) (*models.Weapon, error) {
	attr, err := json.Marshal(weapon.Attr)
	if err != nil {
		return nil, fmt.Errorf("encode weapon attr: %w", err)
	}

	var created models.Weapon
	var rawAttr []byte
	err = pool.QueryRow(context.Background(), `
		INSERT INTO weapons (name, description, weapon_type, cost, attr)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id::text, name, description, weapon_type, cost, attr, created, updated
	`, weapon.Name, weapon.Desc, weapon.Type, weapon.Cost, attr).Scan(
		&created.ID,
		&created.Name,
		&created.Desc,
		&created.Type,
		&created.Cost,
		&rawAttr,
		&created.Created,
		&created.Updated,
	)
	if err != nil {
		return nil, fmt.Errorf("create weapon: %w", err)
	}

	if len(rawAttr) > 0 {
		if err := json.Unmarshal(rawAttr, &created.Attr); err != nil {
			return nil, fmt.Errorf("decode weapon attr: %w", err)
		}
	}

	return &created, nil
}

func GetAllWeapons(pool *pgxpool.Pool) ([]*models.Weapon, error) {
	rows, err := pool.Query(context.Background(), `
		SELECT id::text, name, description, weapon_type, cost, attr, created, updated
		FROM weapons
		ORDER BY created DESC
	`)
	if err != nil {
		return nil, fmt.Errorf("get weapons: %w", err)
	}
	defer rows.Close()

	weapons := make([]*models.Weapon, 0)
	for rows.Next() {
		var weapon models.Weapon
		var rawAttr []byte
		if err := rows.Scan(
			&weapon.ID,
			&weapon.Name,
			&weapon.Desc,
			&weapon.Type,
			&weapon.Cost,
			&rawAttr,
			&weapon.Created,
			&weapon.Updated,
		); err != nil {
			return nil, fmt.Errorf("scan weapon: %w", err)
		}

		if len(rawAttr) > 0 {
			if err := json.Unmarshal(rawAttr, &weapon.Attr); err != nil {
				return nil, fmt.Errorf("decode weapon attr: %w", err)
			}
		}

		weapons = append(weapons, &weapon)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate weapons: %w", err)
	}

	return weapons, nil
}

func weaponExists(ctx context.Context, pool *pgxpool.Pool, weaponID string) (bool, error) {
	var exists bool
	err := pool.QueryRow(ctx, `
		SELECT EXISTS(SELECT 1 FROM weapons WHERE id = $1::uuid)
	`, weaponID).Scan(&exists)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}

		return false, fmt.Errorf("check weapon exists: %w", err)
	}

	return exists, nil
}
