package main

import (

	//"std/configs"
	"std/controller"
	"std/db"

	// "std/controller"
	"std/model"

	"github.com/gin-gonic/gin"
)

func main() {

	// err := configs.Load()
	_, err := db.OpenConn()

	if err != nil {
		panic(err)
	}

	// fmt.Println(model.GetCryptosCMCbyID("1"))

	router := gin.Default()

	// GET ONE CRYPTO FROM COIN MARKET CAP REQUEST BY ID
	router.GET("/cryptos/:id", controller.GetCryptosCMCbyID)

	// GET ONE CRYPTO FROM DATABASE BY ID
	router.GET("/cryptos/getindb/:id", controller.GetOneInDb)

	// GET ALL CRYPTOS FROM DATABASE
	router.GET("/cryptos/getindb", controller.Get)

	// GET ALL CRYPTOS FROM COIN MARKET CAP REQUEST
	router.GET("/cryptos", model.GetCryptosByCMC)

	// DELETE ALL DATA FROM TABLE CRYPTOS IN DATABASE
	router.DELETE("/cryptos", controller.DeleteDataInDB)

	// UPVOTE A CRYPTO BY ID
	router.PATCH("/cryptos/:id", controller.ToggleCryptoUpvoteCMC)

	// INSERT ALL CRYPTOS FROM COIN MARKET CAP REQUEST TO DATABASE
	router.POST("/cryptos/insert/:op", controller.Create)

	router.Run("localhost:9090")

}
