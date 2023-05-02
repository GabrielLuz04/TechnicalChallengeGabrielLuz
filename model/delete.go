package model

import (
	"std/db"
)

func Clean() (int64, error) {

	conn, _ := db.OpenConn()

	sqlStatement, err := conn.Exec(`DELETE FROM cryptos
	WHERE 1 = 1`)

	if err != nil {
		return 0, err
	}

	return sqlStatement.RowsAffected()

}
