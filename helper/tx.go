package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	if err := recover(); err != nil {
		errorRollback := tx.Rollback()
		PanicIfError(errorRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		PanicIfError(errorCommit)
	}
}
