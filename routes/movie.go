package routes

import (
	"net/http"

	"github.com/cacawildaandika/go-rest-cinema/dtos"
	"github.com/gin-gonic/gin"
)

func InitializeMovieRouter(rg *gin.RouterGroup) {

	router := rg.Group("/movie")

	router.GET("/", GetAllMovie)
}

func GetAllMovie(context *gin.Context) {

	response := dtos.Response{
		Status:  "Ok",
		Code:    http.StatusOK,
		Message: "Get all ticket",
	}

	context.JSON(http.StatusOK, response)
}
