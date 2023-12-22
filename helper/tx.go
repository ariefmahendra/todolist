package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx, err error) {
	// err := recover()
	if err != nil {
		errRollback := tx.Rollback()
		PanicIfError(errRollback)
	} else {
		errCommit := tx.Commit()
		PanicIfError(errCommit)
	}
}
