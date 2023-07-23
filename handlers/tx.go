package handlers

import (
	"net/http"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/gin-gonic/gin"
	nodetypes "github.com/sentinel-official/hub/x/node/types"
	plantypes "github.com/sentinel-official/hub/x/plan/types"
	sessiontypes "github.com/sentinel-official/hub/x/session/types"

	"github.com/solarlabsteam/sentinel-api-backend/context"
	"github.com/solarlabsteam/sentinel-api-backend/requests"
	"github.com/solarlabsteam/sentinel-api-backend/types"
	"github.com/solarlabsteam/sentinel-api-backend/utils"
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

		result, err := ctx.Tx(
			kr, key.GetName(), req.Query.Gas, req.Query.GasAdjustment, req.Query.GasPrices,
			req.Body.Fees, req.FeeGranter, req.Body.Memo, req.Body.SignMode, req.Query.ChainID, req.Query.RPCAddress,
			req.Body.TimeoutHeight, req.Query.SimulateAndExecute, req.Query.BroadcastMode, message,
		)
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

		message := nodetypes.NewMsgSubscribeRequest(key.GetAddress(), req.NodeAddress, req.Body.Gigabytes, req.Body.Hours, req.Body.Denom)

		result, err := ctx.Tx(
			kr, key.GetName(), req.Query.Gas, req.Query.GasAdjustment, req.Query.GasPrices,
			req.Body.Fees, req.FeeGranter, req.Body.Memo, req.Body.SignMode, req.Query.ChainID, req.Query.RPCAddress,
			req.Body.TimeoutHeight, req.Query.SimulateAndExecute, req.Query.BroadcastMode, message,
		)
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

		message := plantypes.NewMsgSubscribeRequest(key.GetAddress(), req.URI.ID, req.Body.Denom)

		result, err := ctx.Tx(
			kr, key.GetName(), req.Query.Gas, req.Query.GasAdjustment, req.Query.GasPrices,
			req.Body.Fees, req.FeeGranter, req.Body.Memo, req.Body.SignMode, req.Query.ChainID, req.Query.RPCAddress,
			req.Body.TimeoutHeight, req.Query.SimulateAndExecute, req.Query.BroadcastMode, message,
		)
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

		result, err := ctx.Tx(
			kr, key.GetName(), req.Query.Gas, req.Query.GasAdjustment, req.Query.GasPrices,
			req.Body.Fees, req.FeeGranter, req.Body.Memo, req.Body.SignMode, req.Query.ChainID, req.Query.RPCAddress,
			req.Body.TimeoutHeight, req.Query.SimulateAndExecute, req.Query.BroadcastMode, message,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(3, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(result))
	}
}
