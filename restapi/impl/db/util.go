package db

import (
	"database/sql"

	"github.com/cyverse-de/permissions/logger"
)

// Attempts to roll back a transaction and logs any errors that are returned. ErrTxDone is never logged because that
// will happen if this function is called using `defer` and an update succeeds.
func RollbackTx(tx *sql.Tx) {
	err := tx.Rollback()
	if err != nil && err != sql.ErrTxDone {
		logger.Log.Errorf("rolling back transaction: %v", err)
	}
}
