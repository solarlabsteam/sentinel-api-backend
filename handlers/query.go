package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sentinel-official/api-client/context"
	"github.com/sentinel-official/api-client/requests"
	"github.com/sentinel-official/api-client/types"
)

func HandlerGetAccount(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := requests.NewRequestGetAccount(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewResponseError(1, err))
			return
		}

		result, err := ctx.QueryAccount(req.Query.RPCAddress, req.AccAddress)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(2, err))
			return
		}

		buf, err := ctx.Codec.MarshalJSON(result)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(3, err))
			return
		}

		var m json.RawMessage
		if err := json.Unmarshal(buf, &m); err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(4, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(m))
	}
}

func HandlerGetBalancesForAccount(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := requests.NewRequestGetBalancesForAccount(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewResponseError(1, err))
			return
		}

		result, err := ctx.QueryBalances(req.Query.RPCAddress, req.AccAddress, req.Pagination)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(2, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(result))
	}
}

func HandlerGetSessionsForAccount(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := requests.NewRequestGetSessionsForAccount(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewResponseError(1, err))
			return
		}

		result, err := ctx.QuerySessionsForAddress(req.Query.RPCAddress, req.AccAddress, req.Status, req.Pagination)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(2, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(result))
	}
}

func HandlerGetSubscriptionsForAccount(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := requests.NewRequestGetSubscriptionsForAccount(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewResponseError(1, err))
			return
		}

		result, err := ctx.QuerySubscriptionsForAddress(req.Query.RPCAddress, req.AccAddress, req.Status, req.Pagination)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(2, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(result))
	}
}

func HandlerGetDeposits(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := requests.NewRequestGetDeposits(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewResponseError(1, err))
			return
		}

		result, err := ctx.QueryDeposits(req.Query.RPCAddress, req.Pagination)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(2, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(result))
	}
}

func HandlerGetDeposit(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := requests.NewRequestGetDeposit(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewResponseError(1, err))
			return
		}

		result, err := ctx.QueryDeposit(req.Query.RPCAddress, req.AccAddress)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(2, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(result))
	}
}

func HandlerGetNodes(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := requests.NewRequestGetNodes(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewResponseError(1, err))
			return
		}

		result, err := ctx.QueryNodes(req.Query.RPCAddress, req.Status, req.Pagination)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(2, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(result))
	}
}

func HandlerGetNode(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := requests.NewRequestGetNode(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewResponseError(1, err))
			return
		}

		result, err := ctx.QueryNode(req.Query.RPCAddress, req.NodeAddress)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(2, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(result))
	}
}

func HandlerGetPlans(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := requests.NewRequestGetPlans(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewResponseError(1, err))
			return
		}

		result, err := ctx.QueryPlans(req.Query.RPCAddress, req.Status, req.Pagination)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(2, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(result))
	}
}

func HandlerGetPlan(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := requests.NewRequestGetPlan(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewResponseError(1, err))
			return
		}

		result, err := ctx.QueryPlan(req.Query.RPCAddress, req.URI.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(2, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(result))
	}
}

func HandlerGetProviders(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := requests.NewRequestGetProviders(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewResponseError(1, err))
			return
		}

		result, err := ctx.QueryProviders(req.Query.RPCAddress, req.Pagination)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(2, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(result))
	}
}

func HandlerGetProvider(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := requests.NewRequestGetProvider(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewResponseError(1, err))
			return
		}

		result, err := ctx.QueryProvider(req.Query.RPCAddress, req.ProvAddress)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(2, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(result))
	}
}

func HandlerGetNodesForProvider(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := requests.NewRequestGetNodesForProvider(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewResponseError(1, err))
			return
		}

		result, err := ctx.QueryNodesForProvider(req.Query.RPCAddress, req.ProvAddress, req.Status, req.Pagination)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(2, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(result))
	}
}

func HandlerGetPlansForProvider(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := requests.NewRequestGetPlansForProvider(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewResponseError(1, err))
			return
		}

		result, err := ctx.QueryPlansForProvider(req.Query.RPCAddress, req.ProvAddress, req.Status, req.Pagination)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(2, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(result))
	}
}

func HandlerGetSessions(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := requests.NewRequestGetSessions(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewResponseError(1, err))
			return
		}

		result, err := ctx.QuerySessions(req.Query.RPCAddress, req.Pagination)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(2, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(result))
	}
}

func HandlerGetSession(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := requests.NewRequestGetSession(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewResponseError(1, err))
			return
		}

		result, err := ctx.QuerySession(req.Query.RPCAddress, req.URI.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(2, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(result))
	}
}

func HandlerGetSubscriptions(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := requests.NewRequestGetSubscriptions(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewResponseError(1, err))
			return
		}

		result, err := ctx.QuerySubscriptions(req.Query.RPCAddress, req.Pagination)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(2, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(result))
	}
}

func HandlerGetSubscription(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := requests.NewRequestGetSubscription(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewResponseError(1, err))
			return
		}

		result, err := ctx.QuerySubscription(req.Query.RPCAddress, req.URI.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(2, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(result))
	}
}

func HandlerGetQuotasForSubscription(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := requests.NewRequestGetQuotasForSubscription(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewResponseError(1, err))
			return
		}

		result, err := ctx.QueryQuotas(req.Query.RPCAddress, req.URI.ID, req.Pagination)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(2, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(result))
	}
}

func HandlerGetQuotaForSubscription(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := requests.NewRequestGetQuotaForSubscription(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewResponseError(1, err))
			return
		}

		result, err := ctx.QueryQuota(req.Query.RPCAddress, req.URI.ID, req.AccAddress)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(2, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(result))
	}
}
