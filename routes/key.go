package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/solarlabsteam/sentinel-api-backend/context"
	"github.com/solarlabsteam/sentinel-api-backend/handlers"
)

func RegisterKeyRoutes(router gin.IRouter, ctx context.Context) {
	router.POST("/nodes/:node_address/sessions/:id/keys", handlers.HandlerAddSessionKey(ctx))
}
