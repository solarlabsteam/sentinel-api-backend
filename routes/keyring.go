package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/solarlabsteam/sentinel-api-backend/context"
	"github.com/solarlabsteam/sentinel-api-backend/handlers"
)

func RegisterKeyringRoutes(router gin.IRouter, ctx context.Context) {
	router.POST("/signatures", handlers.HandlerGenerateSignature(ctx))
}
