package app

import "database/sql"

func (a *App) transact(fn func(*sql.Tx) error) (err error) {
	var tx *sql.Tx
	tx, err = a.db.Begin()
	if err != nil {
		return err
	}

	rollback := true

	defer func() {
		if rollback {
			if rerr := tx.Rollback(); err == nil && rerr != nil {
				err = rerr
			}
		}
	}()

	err = fn(tx)
	rollback = false

	if err != nil {
		return tx.Rollback()
	}

	rollback = false
	return tx.Commit()
}
