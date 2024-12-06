package handler

import (
	"cbo-api/api/routes"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Handler(_ http.ResponseWriter, _ *http.Request) {
	router := gin.Default()
	router = routes.SetupRoutes(router)
	err := router.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
