package main

import (
	"std/model"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/cryptos", model.GetCryptos)
	router.GET("/cryptos/:id", model.GetCryptosCMCbyID)
	router.GET("/cryptos/all", model.GetCryptosByCMC)
	router.PATCH("/cryptos/:id", model.ToggleCryptoUpvote)
	router.Run("localhost:9090")

}
