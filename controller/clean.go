package controller

import (
	"fmt"
	"log"
	"net/http"
	"std/model"

	"github.com/gin-gonic/gin"
)

func DeleteDataInDB(context *gin.Context) {

	rows, err := model.Clean()

	if err != nil {
		log.Printf("[delete] Erro ao remover o registro: %v", err)
	}

	if rows > 1 {
		log.Printf("[delete] Error: Foram deletados %d registros", rows)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": "Dados removidos com sucesso!",
	}

	if rows == 0 {
		log.Print("[delete] O banco não possui dados para serem deletados.")
		resp = map[string]any{
			"Error":   false,
			"Message": "O banco não possui dados para serem deletados.",
		}
	}
	fmt.Println(resp)
	context.IndentedJSON(http.StatusOK, resp)

}
