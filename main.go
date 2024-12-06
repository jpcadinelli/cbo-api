package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
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

	r.GET("/cbos/:id", func(c *gin.Context) {
		id := c.Param("id")

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

		content, err := io.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading cbo.json", "data": nil})
			return
		}

		var cbos []CBO
		err = json.Unmarshal(content, &cbos)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing cbo.json", "data": nil})
			return
		}

		var foundCBO *CBO
		for _, cbo := range cbos {
			if cbo.Codigo == id {
				foundCBO = &cbo
				break
			}
		}

		if foundCBO == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "CBO not found", "data": nil})
			return
		}

		c.JSON(http.StatusOK, gin.H{"error": nil, "data": foundCBO})
	})

	r.POST("/cbos/tipo", func(c *gin.Context) {
		var request CBO

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
			return
		}

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

		content, err := io.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading cbo.json", "data": nil})
			return
		}

		var cbos []CBO
		err = json.Unmarshal(content, &cbos)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing cbo.json", "data": nil})
			return
		}

		var filteredCBOs []CBO
		for _, cbo := range cbos {
			if cbo.Tipo == request.Tipo {
				filteredCBOs = append(filteredCBOs, cbo)
			}
		}

		if len(filteredCBOs) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "No CBOs found for the given type", "data": nil})
			return
		}

		c.JSON(http.StatusOK, gin.H{"error": nil, "data": filteredCBOs})
	})

	r.POST("/cbos/nome", func(c *gin.Context) {
		var request CBO

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
			return
		}

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

		content, err := io.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading cbo.json", "data": nil})
			return
		}

		var cbos []CBO
		err = json.Unmarshal(content, &cbos)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing cbo.json", "data": nil})
			return
		}

		var filteredCBOs []CBO
		for _, cbo := range cbos {
			if strings.Contains(strings.ToLower(cbo.Nome), strings.ToLower(request.Nome)) {
				filteredCBOs = append(filteredCBOs, cbo)
			}
		}

		if len(filteredCBOs) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "No CBOs found for the given name", "data": nil})
			return
		}

		c.JSON(http.StatusOK, gin.H{"error": nil, "data": filteredCBOs})
	})

	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
