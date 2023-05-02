package controller

import (
	"log"
	"net/http"
	"std/model"

	"github.com/gin-gonic/gin"
)

func GetCryptosCMCbyID(context *gin.Context) {
	id := context.Param("id")
	crypto, err := model.GetOne(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "crypto not found"})
	}

	context.IndentedJSON(http.StatusOK, crypto)

}

func GetOneInDb(context *gin.Context) {
	id := context.Param("id")
	cryptos, err := model.GetOneDb(id)

	if err != nil {
		log.Printf("[list] Erro ao obter registros: %v\n", err)
	}

	context.Header("Content-Type", "application/json")
	context.JSON(200, cryptos)
}
