package handlers

import (
	"net/http"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/gin-gonic/gin"
	sessiontypes "github.com/sentinel-official/hub/x/session/types"
	subscriptiontypes "github.com/sentinel-official/hub/x/subscription/types"

	"github.com/sentinel-official/api-client/context"
	"github.com/sentinel-official/api-client/requests"
	"github.com/sentinel-official/api-client/types"
	"github.com/sentinel-official/api-client/utils"
)

func HandlerTxBankSend(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := requests.NewRequestTxBankSend(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewResponseError(1, err))
			return
		}

		kr, key, err := utils.NewInMemoryKey(req.Body.Mnemonic, req.Query.CoinType, req.Query.Account, req.Query.Index, req.Body.BIP39Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(2, err))
			return
		}

		message := banktypes.NewMsgSend(key.GetAddress(), req.ToAccAddress, req.Amount)

		result, err := ctx.Tx(kr, key.GetName(), req.Query.Gas, req.Body.Memo, req.GasPrices, req.Query.ChainID, req.Query.RPCAddress, req.Query.BroadcastMode, message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(3, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(result))
	}
}

func HandlerTxSubscribeToNode(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := requests.NewRequestTxSubscribeToNode(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewResponseError(1, err))
			return
		}

		kr, key, err := utils.NewInMemoryKey(req.Body.Mnemonic, req.Query.CoinType, req.Query.Account, req.Query.Index, req.Body.BIP39Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(2, err))
			return
		}

		message := subscriptiontypes.NewMsgSubscribeToNodeRequest(key.GetAddress(), req.NodeAddress, req.Deposit)

		result, err := ctx.Tx(kr, key.GetName(), req.Query.Gas, req.Body.Memo, req.GasPrices, req.Query.ChainID, req.Query.RPCAddress, req.Query.BroadcastMode, message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(3, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(result))
	}
}

func HandlerTxSubscribeToPlan(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := requests.NewRequestTxSubscribeToPlan(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewResponseError(1, err))
			return
		}

		kr, key, err := utils.NewInMemoryKey(req.Body.Mnemonic, req.Query.CoinType, req.Query.Account, req.Query.Index, req.Body.BIP39Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(2, err))
			return
		}

		message := subscriptiontypes.NewMsgSubscribeToPlanRequest(key.GetAddress(), req.URI.ID, req.Body.Denom)

		result, err := ctx.Tx(kr, key.GetName(), req.Query.Gas, req.Body.Memo, req.GasPrices, req.Query.ChainID, req.Query.RPCAddress, req.Query.BroadcastMode, message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(3, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(result))
	}
}

func HandlerTxStartSession(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := requests.NewRequestTxStartSession(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewResponseError(1, err))
			return
		}

		kr, key, err := utils.NewInMemoryKey(req.Body.Mnemonic, req.Query.CoinType, req.Query.Account, req.Query.Index, req.Body.BIP39Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(2, err))
			return
		}

		message := sessiontypes.NewMsgStartRequest(key.GetAddress(), req.URI.ID, req.NodeAddress)

		result, err := ctx.Tx(kr, key.GetName(), req.Query.Gas, req.Body.Memo, req.GasPrices, req.Query.ChainID, req.Query.RPCAddress, req.Query.BroadcastMode, message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(3, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(result))
	}
}
