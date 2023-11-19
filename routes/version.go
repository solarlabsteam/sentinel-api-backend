package routes

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/version"
	"github.com/gin-gonic/gin"

	"github.com/solarlabsteam/sentinel-api-backend/context"
)

func RegisterVersionRoutes(router gin.IRouter, _ context.Context) {
	router.GET("/version.json", func(c *gin.Context) {
		c.JSON(http.StatusOK, &struct {
			Commit string `json:"commit"`
		}{
			Commit: version.Commit,
		})
	})
}
