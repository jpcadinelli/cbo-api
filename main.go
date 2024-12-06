package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

type CBO struct {
	Codigo string `json:"Codigo"`
	Nome   string `json:"Nome"`
	Tipo   string `json:"Tipo"`
}

func main() {
	r := gin.Default()

	r.GET("/cbos", func(c *gin.Context) {
		file, err := os.Open("cbo.json")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read cbo.json", "data": nil})
			return
		}
		defer func(file *os.File) {
			err = file.Close()
			if err != nil {
				log.Println("Error closing file:", err)
			}
		}(file)

		var cbos []CBO

		err = json.NewDecoder(file).Decode(&cbos)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing cbo.json", "data": nil})
			return
		}

		c.JSON(http.StatusOK, gin.H{"error": nil, "data": cbos})
	})

	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
