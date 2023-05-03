package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func OpenConn() (*sqlx.DB, error) {
	Db, err := sqlx.Open("postgres", "postgres://postgres:1234@0.0.0.0:5432/apitech?sslmode=disable")

	if err != nil {
		panic(err)
	}

	err = Db.Ping()

	if err != nil {
		log.Printf("ERRO AO PINGAR O BD")
		panic(err)
	}

	return Db, err

}
