package main

import (
	"cbo-api/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	r = routes.SetupRoutes(r)
	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
