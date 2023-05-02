package controller

import (
	"fmt"
	"net/http"
	"std/entities"
	"std/model"

	"github.com/gin-gonic/gin"
)

func Create(context *gin.Context) {

	var crypto entities.Crypto
	op := context.Param("op")
	var Resp map[string]any

	if op == "all" {
		_, err := model.InsertCryptosInDB(&crypto)

		if err != nil {
			Resp = map[string]any{
				"Error":   true,
				"Message": fmt.Sprintf("Ocorreu um erro ao inserir as Cryptos: %v\n", err),
			}
		} else {
			Resp = map[string]any{
				"Error":   false,
				"Message": fmt.Sprintf("Cryptos inseridas com sucesso: %v", len(entities.CryptosCMC.Data)),
			}
		}
	} else {
		Resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("OPÇÃO INVÁLIDA: %v", crypto),
		}
	}

	context.IndentedJSON(http.StatusOK, Resp)

}
