package routes

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func Run() {
	v1 := router.Group("/v1")
	InitializeMovieRouter(v1)
	InitializeTicketRouter(v1)

	router.Run(":5000")
}
