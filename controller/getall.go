package controller

import (
	"log"
	"std/model"

	"github.com/gin-gonic/gin"
)

func Get(w *gin.Context) {

	cryptos, err := model.GetAllInDB()

	if err != nil {
		log.Printf("[list] Erro ao obter registros: %v\n", err)
	}

	w.Header("Content-Type", "application/json")
	w.JSON(200, cryptos)
}
