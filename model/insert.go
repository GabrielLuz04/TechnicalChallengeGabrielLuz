package model

import (
	"fmt"
	"std/db"
	"std/entities"
)

func InsertCryptosInDB(*entities.Crypto) (int64, error) {

	// crypto := entities.Crypto{}

	conn, err := db.OpenConn()

	if err != nil {
		return 0, err
	}

	defer conn.Close()

	SizeOfCryptosSlice := len(entities.CryptosCMC.Data)

	for index, result := range entities.CryptosCMC.Data {
		// fmt.Println(result)

		if index < SizeOfCryptosSlice {
			index++
			fmt.Println("\n tamanho do index", index)
			_, err := conn.Exec(`INSERT INTO cryptos (
				upvotes, id, name, symbol, slug, price
				) VALUES ($1, $2, $3, $4, $5, $6);`,
				result.Upvotes, result.ID,
				result.Name, result.Symbol,
				result.Slug, result.Quote.USD.Price,
			)
			if err != nil {
				fmt.Println("[insert] ERRO AO INSERIR CRYPTO NO BD")
			}

		}

		// return res.RowsAffected()

	}

	return 0, err

}
