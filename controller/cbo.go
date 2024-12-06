package controller

import (
	"cbo-api/dataprovider"
	"cbo-api/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func ListarCBO(c *gin.Context) {
	cbos, err := dataprovider.GetListCBO()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "data": nil})
	}

	c.JSON(http.StatusOK, gin.H{"error": nil, "data": cbos})
}

func VisualizarCBO(c *gin.Context) {
	id := c.Param("id")

	cbos, err := dataprovider.GetListCBO()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "data": nil})
	}

	var foundCBO *domain.CBO
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
}

func FiltrarCBOTipo(c *gin.Context) {
	var request domain.CBO

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	cbos, err := dataprovider.GetListCBO()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "data": nil})
	}

	var filteredCBOs []domain.CBO
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
}

func FiltrarCBONome(c *gin.Context) {
	var request domain.CBO

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	cbos, err := dataprovider.GetListCBO()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "data": nil})
	}

	var filteredCBOs []domain.CBO
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
}
