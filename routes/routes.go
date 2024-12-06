package routes

import (
	"cbo-api/controller"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("/api/v1")
	{
		cboGroup := main.Group("/cbos")
		{
			cboGroup.GET("", controller.ListarCBO)
			cboGroup.GET("/:id", controller.VisualizarCBO)
			cboGroup.POST("/tipo", controller.FiltrarCBOTipo)
			cboGroup.POST("/nome", controller.FiltrarCBONome)
		}
	}
	return router
}
