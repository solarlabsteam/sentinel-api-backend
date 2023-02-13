package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/sentinel-official/api-client/context"
	"github.com/sentinel-official/api-client/handlers"
)

func RegisterQueryRoutes(router gin.IRouter, ctx context.Context) {
	router.GET("/accounts/:acc_address", handlers.HandlerGetAccount(ctx))
	router.GET("/accounts/:acc_address/balances", handlers.HandlerGetBalancesForAccount(ctx))
	router.GET("/accounts/:acc_address/sessions", handlers.HandlerGetSessionsForAccount(ctx))
	router.GET("/accounts/:acc_address/subscriptions", handlers.HandlerGetSubscriptionsForAccount(ctx))

	router.GET("/deposits", handlers.HandlerGetDeposits(ctx))
	router.GET("/deposits/:acc_address", handlers.HandlerGetDeposit(ctx))

	router.GET("/nodes", handlers.HandlerGetNodes(ctx))
	router.GET("/nodes/:node_address", handlers.HandlerGetNode(ctx))

	router.GET("/plans", handlers.HandlerGetPlans(ctx))
	router.GET("/plans/:id", handlers.HandlerGetPlan(ctx))

	router.GET("/providers", handlers.HandlerGetProviders(ctx))
	router.GET("/providers/:prov_address", handlers.HandlerGetProvider(ctx))
	router.GET("/providers/:prov_address/nodes", handlers.HandlerGetNodesForProvider(ctx))
	router.GET("/providers/:prov_address/plans", handlers.HandlerGetPlansForProvider(ctx))

	router.GET("/sessions", handlers.HandlerGetSessions(ctx))
	router.GET("/sessions/:id", handlers.HandlerGetSession(ctx))

	router.GET("/subscriptions", handlers.HandlerGetSubscriptions(ctx))
	router.GET("/subscriptions/:id", handlers.HandlerGetSubscription(ctx))
	router.GET("/subscriptions/:id/quotas", handlers.HandlerGetQuotasForSubscription(ctx))
	router.GET("/subscriptions/:id/quotas/:acc_address", handlers.HandlerGetQuotaForSubscription(ctx))
}
