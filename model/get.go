package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"std/entities"
	"time"

	"github.com/gin-gonic/gin"
)

func GetCryptos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, entities.Cryptos)
}

func GetCryptoById(id string) (*entities.Crypto, error) {
	for i, t := range entities.Cryptos {
		if t.Id == id {
			return &entities.Cryptos[i], nil
		}
	}

	return nil, errors.New("crypto not found")
}

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

		context.IndentedJSON(http.StatusOK, post)
	}
}

func GetCryptosCMCbyID(id int) (*entities.CryptoRequest, error) {

	var CryptosCMC = entities.CryptoRequest{}

	for index, result := range CryptosCMC.Data {
		if result.ID == id {
			return CryptosCMC, nil
		}
	}

	return nil, errors.New("crypto not found")
}
