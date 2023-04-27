package model

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ToggleCryptoUpvote(context *gin.Context) {
	id := context.Param("id")
	crypto, err := GetCryptoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "crypto not found"})
	}

	crypto.Upvotes = crypto.Upvotes + 1

	context.IndentedJSON(http.StatusOK, crypto)

}
