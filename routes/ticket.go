package routes

import (
	"net/http"

	"github.com/cacawildaandika/go-rest-cinema/dtos"
	"github.com/gin-gonic/gin"
)

func InitializeTicketRouter(rg *gin.RouterGroup) {

	router := rg.Group("/ticket")

	router.GET("/", GetAllTicket)
}

func GetAllTicket(context *gin.Context) {
	response := dtos.Response{
		Status:  "Ok",
		Code:    http.StatusOK,
		Message: "Get all ticket",
	}

	context.JSON(http.StatusOK, response)
}
