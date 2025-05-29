package helpers

import (
	"database/sql"
	customErrors "dr.agenda/errors"
)

func CloseStmt(rows *sql.Stmt) (any, *error) {
	err := rows.Close()
	if err != nil {
		return nil, &customErrors.StmtCloseError
	}
	return nil, nil
}

func CloseRows(rows *sql.Rows) (any, *error) {
	err := rows.Close()
	if err != nil {
		return nil, &customErrors.RowsCloseError
	}
	return nil, nil
}
