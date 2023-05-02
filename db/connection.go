package db

import (
	"database/sql"
	"fmt"
	"log"
	"std/configs"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	stringConnection := fmt.Sprintf("host=%s port=%s user=%s pass=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	conn, err := sql.Open("postgres", stringConnection)

	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err

}

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
