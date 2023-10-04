package handlers

import (
	"net/http"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/gin-gonic/gin"
	nodetypes "github.com/sentinel-official/hub/x/node/types"
	plantypes "github.com/sentinel-official/hub/x/plan/types"
	sessiontypes "github.com/sentinel-official/hub/x/session/types"
	subscriptiontypes "github.com/sentinel-official/hub/x/subscription/types"

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

		var messages []sdk.Msg
		for i := 0; i < len(req.ToAccAddresses); i++ {
			messages = append(messages, banktypes.NewMsgSend(key.GetAddress(), req.ToAccAddresses[i], req.Amounts[i]))
		}

		result, err := ctx.Tx(
			kr, key.GetName(), req.Query.Gas, req.Query.GasAdjustment, req.Query.GasPrices,
			req.Body.Fees, req.FeeGranter, req.Body.Memo, req.Body.SignMode, req.Query.ChainID, req.Query.RPCAddress,
			req.Body.TimeoutHeight, req.Query.SimulateAndExecute, req.Query.BroadcastMode, messages...,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(3, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(result))
	}
}

func HandlerTxPlanCreate(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := requests.NewRequestTxPlanCreate(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewResponseError(1, err))
			return
		}

		kr, key, err := utils.NewInMemoryKey(req.Body.Mnemonic, req.Query.CoinType, req.Query.Account, req.Query.Index, req.Body.BIP39Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(2, err))
			return
		}

		message := plantypes.NewMsgCreateRequest(key.GetAddress().Bytes(), req.Body.Duration, req.Body.Gigabytes, req.Prices)

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

func HandlerTxPlanUpdateStatus(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := requests.NewRequestTxPlanUpdateStatus(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewResponseError(1, err))
			return
		}

		kr, key, err := utils.NewInMemoryKey(req.Body.Mnemonic, req.Query.CoinType, req.Query.Account, req.Query.Index, req.Body.BIP39Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(2, err))
			return
		}

		message := plantypes.NewMsgUpdateStatusRequest(key.GetAddress().Bytes(), req.URI.ID, req.Status)

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

func HandlerTxPlanLinkNode(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := requests.NewRequestTxPlanLinkNode(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewResponseError(1, err))
			return
		}

		kr, key, err := utils.NewInMemoryKey(req.Body.Mnemonic, req.Query.CoinType, req.Query.Account, req.Query.Index, req.Body.BIP39Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(2, err))
			return
		}

		var messages []sdk.Msg
		for i := 0; i < len(req.NodeAddresses); i++ {
			messages = append(messages, plantypes.NewMsgLinkNodeRequest(key.GetAddress().Bytes(), req.URI.ID, req.NodeAddresses[i]))
		}

		result, err := ctx.Tx(
			kr, key.GetName(), req.Query.Gas, req.Query.GasAdjustment, req.Query.GasPrices,
			req.Body.Fees, req.FeeGranter, req.Body.Memo, req.Body.SignMode, req.Query.ChainID, req.Query.RPCAddress,
			req.Body.TimeoutHeight, req.Query.SimulateAndExecute, req.Query.BroadcastMode, messages...,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(3, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(result))
	}
}

func HandlerTxPlanUnlinkNode(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := requests.NewRequestTxPlanUnlinkNode(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewResponseError(1, err))
			return
		}

		kr, key, err := utils.NewInMemoryKey(req.Body.Mnemonic, req.Query.CoinType, req.Query.Account, req.Query.Index, req.Body.BIP39Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(2, err))
			return
		}

		message := plantypes.NewMsgUnlinkNodeRequest(key.GetAddress().Bytes(), req.URI.ID, req.NodeAddress)

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

func HandlerTxNodeSubscribe(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := requests.NewRequestTxNodeSubscribe(c)
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

func HandlerTxPlanSubscribe(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := requests.NewRequestTxPlanSubscribe(c)
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

func HandlerTxSubscriptionAllocate(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := requests.NewRequestTxSubscriptionAllocate(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewResponseError(1, err))
			return
		}

		kr, key, err := utils.NewInMemoryKey(req.Body.Mnemonic, req.Query.CoinType, req.Query.Account, req.Query.Index, req.Body.BIP39Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(2, err))
			return
		}

		var messages []sdk.Msg
		for i := 0; i < len(req.AccAddresses); i++ {
			messages = append(messages, subscriptiontypes.NewMsgAllocateRequest(key.GetAddress(), req.URI.ID, req.AccAddresses[i], req.Bytes[i]))
		}

		result, err := ctx.Tx(
			kr, key.GetName(), req.Query.Gas, req.Query.GasAdjustment, req.Query.GasPrices,
			req.Body.Fees, req.FeeGranter, req.Body.Memo, req.Body.SignMode, req.Query.ChainID, req.Query.RPCAddress,
			req.Body.TimeoutHeight, req.Query.SimulateAndExecute, req.Query.BroadcastMode, messages...,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(3, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(result))
	}
}

func HandlerTxSessionStart(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := requests.NewRequestTxSessionStart(c)
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

func HandlerTxSubscribe(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := requests.NewRequestTxSubscribe(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewResponseError(1, err))
			return
		}

		kr, key, err := utils.NewInMemoryKey(req.Body.Mnemonic, req.Query.CoinType, req.Query.Account, req.Query.Index, req.Body.BIP39Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(2, err))
			return
		}

		var (
			dIndex   = 0
			messages []sdk.Msg
		)

		for i := 0; i < len(req.NodeAddresses) && dIndex < len(req.Body.Denoms); dIndex, i = dIndex+1, i+1 {
			messages = append(
				messages,
				nodetypes.NewMsgSubscribeRequest(key.GetAddress(), req.NodeAddresses[i], req.Body.Gigabytes[i], req.Body.Hours[i], req.Body.Denoms[dIndex]),
			)
		}
		for i := 0; i < len(req.Body.PlanIDs) && dIndex < len(req.Body.Denoms); dIndex, i = dIndex+1, i+1 {
			messages = append(
				messages,
				plantypes.NewMsgSubscribeRequest(key.GetAddress(), req.Body.PlanIDs[i], req.Body.Denoms[dIndex]),
			)
		}

		result, err := ctx.Tx(
			kr, key.GetName(), req.Query.Gas, req.Query.GasAdjustment, req.Query.GasPrices,
			req.Body.Fees, req.FeeGranter, req.Body.Memo, req.Body.SignMode, req.Query.ChainID, req.Query.RPCAddress,
			req.Body.TimeoutHeight, req.Query.SimulateAndExecute, req.Query.BroadcastMode, messages...,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(3, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(result))
	}
}
