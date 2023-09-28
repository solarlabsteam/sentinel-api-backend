package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/solarlabsteam/sentinel-api-backend/context"
)

func RegisterHealthRoutes(router gin.IRouter, ctx context.Context) {
	router.GET("/robots933456.txt", func(c *gin.Context) {
		c.JSON(http.StatusOK, &struct{}{})
	})
}
