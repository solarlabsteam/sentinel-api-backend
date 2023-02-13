package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/solarlabsteam/sentinel-api-backend/context"
	"github.com/solarlabsteam/sentinel-api-backend/handlers"
)

func RegisterTxRoutes(router gin.IRouter, ctx context.Context) {
	router.POST("/balances", handlers.HandlerTxBankSend(ctx))
	router.POST("/nodes/:node_address/subscriptions", handlers.HandlerTxSubscribeToNode(ctx))
	router.POST("/plans/:id/subscriptions", handlers.HandlerTxSubscribeToPlan(ctx))
	router.POST("/subscriptions/:id/nodes/:node_address/sessions", handlers.HandlerTxStartSession(ctx))
}
