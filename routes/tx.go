package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/solarlabsteam/sentinel-api-backend/context"
	"github.com/solarlabsteam/sentinel-api-backend/handlers"
)

func RegisterTxRoutes(router gin.IRouter, ctx context.Context) {
	router.POST("/balances", handlers.HandlerTxBankSend(ctx))

	router.POST("/nodes/:node_address/subscriptions", handlers.HandlerTxNodeSubscribe(ctx))

	router.POST("/plans", handlers.HandlerTxPlanCreate(ctx))
	router.PUT("/plans/:id", handlers.HandlerTxPlanUpdateStatus(ctx))
	router.POST("/plans/:id/nodes", handlers.HandlerTxPlanLinkNode(ctx))
	router.DELETE("/plans/:id/nodes/:node_address", handlers.HandlerTxPlanUnlinkNode(ctx))
	router.POST("/plans/:id/subscriptions", handlers.HandlerTxPlanSubscribe(ctx))

	router.POST("/subscriptions/:id/allocations", handlers.HandlerTxSubscriptionAllocate(ctx))

	router.POST("/subscriptions/:id/nodes/:node_address/sessions", handlers.HandlerTxSessionStart(ctx))
}
