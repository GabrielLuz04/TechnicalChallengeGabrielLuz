package controller

import (
	"net/http"
	"std/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ToggleCryptoUpvoteCMC(context *gin.Context) {
	id := context.Param("id")
	crypto, err := model.GetOne(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "crypto not found"})
	}

	intId, _ := strconv.Atoi(id)
	model.Update(int64(intId), crypto)

	context.IndentedJSON(http.StatusOK, crypto)

}
