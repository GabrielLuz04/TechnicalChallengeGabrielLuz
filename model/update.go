package model

import (
	"log"
	"std/db"
	"std/entities"
)

func Update(id int64, crypto *entities.Crypto) (int64, error) {
	conn, err := db.OpenConn()

	type cryptoUpdate struct {
		Upvotes int    `json:"upvotes"`
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Symbol  string `json:"symbol"`
		Slug    string `json:"slug"`
		Quote   struct {
			USD struct {
				Price float64 `json:"price"`
			} `json:"USD"`
		} `json:"quote"`
	}

	cryptoSelected := cryptoUpdate{}

	if err != nil {
		return 0, err
	}

	defer conn.Close()

	resposta, err := conn.Query(`select upvotes from cryptos where id=$1;`, crypto.ID)

	if err != nil {
		return 0, err
	}

	for resposta.Next() {
		err = resposta.Scan(&cryptoSelected.Upvotes)
	}

	log.Print(cryptoSelected.Upvotes)

	if err != nil {
		return 0, err
	}

	cryptoSelected.Upvotes = cryptoSelected.Upvotes + 1

	res, err := conn.Exec(`UPDATE cryptos SET upvotes=$1 where id=$2`,
		cryptoSelected.Upvotes, crypto.ID)

	if err != nil {
		return 0, err
	}

	crypto.Upvotes = cryptoSelected.Upvotes

	return res.RowsAffected()

}
