package db

import (
	"context"
	"database/sql"

	"github.com/cyverse-de/permissions/models"
)

// LoadPermissionLevelPrecedence returns a map from permission level to precedence.
func LoadPermissionLevelPrecedence(ctx context.Context, tx *sql.Tx) (map[models.PermissionLevel]int, error) {

	// Query the database.
	query := `SELECT name, precedence FROM permission_levels`
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Build the map.
	precedenceLevelFor := make(map[models.PermissionLevel]int)
	for rows.Next() {
		var name models.PermissionLevel
		var precedence int
		if err := rows.Scan(&name, &precedence); err != nil {
			return nil, err
		}
		precedenceLevelFor[name] = precedence
	}

	return precedenceLevelFor, nil
}
