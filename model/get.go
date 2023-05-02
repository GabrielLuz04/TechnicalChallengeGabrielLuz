package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"std/db"
	"std/entities"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetCryptosByCMC(context *gin.Context) {

	client := &http.Client{
		Timeout: time.Second * 30,
	}

	request, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
	request.Header.Add("X-CMC_PRO_API_KEY", "1537756d-0b5e-446a-bd98-1b5855cc57ae")

	if err != nil {
		fmt.Println("[Main] Erro ao realizar o método GET da página 1", err.Error())
	}

	resposta, err := client.Do(request)

	if err != nil {
		fmt.Println("[Main] Erro ao realizar o método GET da página 1", err.Error())
	}

	if resposta.StatusCode == 200 {
		fmt.Println("ENTROU")
		corpo, err := ioutil.ReadAll(resposta.Body)

		if err != nil {
			fmt.Println("[Main] Erro ao ler o método GET da página 1", err.Error())
		}

		post := entities.CryptoRequest{}

		err = json.Unmarshal(corpo, &post)
		if err != nil {
			fmt.Println("[Main] Erro ao fazer o marshal do GET da página 1", err.Error())
		}

		entities.CryptosCMC = post

		context.IndentedJSON(http.StatusOK, post)

		fmt.Println(len(entities.CryptosCMC.Data))

		// controller.Clean(context)
		// controller.Create(context)

	}
}

func GetOne(id string) (crypto *entities.Crypto, err error) {

	// conn, err := db.OpenConn()
	if err != nil {
		return
	}

	intId, err := strconv.Atoi(id)

	if err != nil {
		return nil, err
	}

	for _, result := range entities.CryptosCMC.Data {
		// fmt.Println(result)
		if result.ID == intId {
			return &result, nil
		}
	}

	return nil, errors.New("crypto not found")
}

func GetOneDb(id string) (crypto *entities.Crypto, err error) {
	conn, err := db.OpenConn()
	intId, err := strconv.Atoi(id)

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
		return nil, err
	}

	defer conn.Close()

	resposta, err := conn.Query(`SELECT * FROM cryptos where id=$1;`, intId)

	if err != nil {
		return nil, err
	}

	for resposta.Next() {
		err = resposta.Scan(&cryptoSelected.Upvotes, &cryptoSelected.ID, &cryptoSelected.Name, &cryptoSelected.Symbol, &cryptoSelected.Slug, &cryptoSelected.Quote.USD.Price)
	}

	if err != nil {
		return nil, err
	}

	crypto = (*entities.Crypto)(&cryptoSelected)

	return crypto, nil

}

func GetAllInDB() (cryptos []entities.Crypto, err error) {
	conn, err := db.OpenConn()

	crypto := entities.Crypto{}

	if err != nil {
		return
	}

	log.Printf("ERRO LINHA 95")

	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM cryptos;`)

	if err != nil {
		return
	}

	log.Printf("ERRO LINHA 102")

	for rows.Next() {

		cryptos = append(cryptos, crypto)

		err = rows.Scan(&crypto.Upvotes, &crypto.ID, &crypto.Name, &crypto.Symbol, &crypto.Slug, &crypto.Quote.USD.Price)

		if err != nil {
			continue
		}

		log.Printf("ID: %v, Name: %v", crypto.ID, crypto.Name)

	}

	log.Print(cryptos)
	return cryptos, nil

}
