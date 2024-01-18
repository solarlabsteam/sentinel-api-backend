package handlers

import (
	"fmt"
	"net/http"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
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

func HandlerTxFeegrantGrantAllowance(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := requests.NewRequestTxFeegrantGrantAllowance(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewResponseError(1, err))
			return
		}

		kr, key, err := utils.NewInMemoryKey(req.Body.Mnemonic, req.Query.CoinType, req.Query.Account, req.Query.Index, req.Body.BIP39Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(2, err))
			return
		}

		fromAddr := key.GetAddress()
		if !req.AuthzGranter.Empty() {
			fromAddr = req.AuthzGranter
		}

		var expiration *time.Time
		if !req.Body.Expiration.IsZero() {
			expiration = &req.Body.Expiration
		}

		var (
			messages       []sdk.Msg
			basicAllowance = &feegrant.BasicAllowance{
				SpendLimit: req.SpendLimit,
				Expiration: expiration,
			}
		)

		for i := 0; i < len(req.AccAddresses); i++ {
			message, err := feegrant.NewMsgGrantAllowance(basicAllowance, fromAddr, req.AccAddresses[i])
			if err != nil {
				c.JSON(http.StatusBadRequest, types.NewResponseError(3, err))
				return
			}

			messages = append(messages, message)
		}

		if !req.AuthzGranter.Empty() {
			execMsg := authz.NewMsgExec(key.GetAddress(), messages)
			messages = []sdk.Msg{&execMsg}
		}

		txResp, err := ctx.Tx(
			kr, key.GetName(), req.Query.Gas, req.Query.GasAdjustment, req.Query.GasPrices,
			req.Body.Fees, req.FeeGranter, req.Body.Memo, req.Body.SignMode, req.Query.ChainID, req.Query.RPCAddress,
			req.Body.TimeoutHeight, req.Query.SimulateAndExecute, req.Query.BroadcastMode, messages...,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(3, err))
			return
		}

		txRes, err := ctx.QueryTxWithRetry(req.Query.RPCAddress, txResp.TxHash, req.Query.MaxQueryTries)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(4, err))
			return
		}
		if txRes == nil {
			err := fmt.Errorf("query result is nil for the transaction %s", txResp.TxHash)
			c.JSON(http.StatusInternalServerError, types.NewResponseError(4, err))
			return
		}
		if !txRes.TxResult.IsOK() {
			err := fmt.Errorf("transaction %s failed with the code %d", txResp.TxHash, txRes.TxResult.Code)
			c.JSON(http.StatusInternalServerError, types.NewResponseError(4, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(txRes))
	}
}

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

		fromAddr := key.GetAddress()
		if !req.AuthzGranter.Empty() {
			fromAddr = req.AuthzGranter
		}

		var messages []sdk.Msg
		for i := 0; i < len(req.ToAccAddresses); i++ {
			messages = append(messages, banktypes.NewMsgSend(fromAddr, req.ToAccAddresses[i], req.Amounts[i]))
		}

		if !req.AuthzGranter.Empty() {
			execMsg := authz.NewMsgExec(key.GetAddress(), messages)
			messages = []sdk.Msg{&execMsg}
		}

		txResp, err := ctx.Tx(
			kr, key.GetName(), req.Query.Gas, req.Query.GasAdjustment, req.Query.GasPrices,
			req.Body.Fees, req.FeeGranter, req.Body.Memo, req.Body.SignMode, req.Query.ChainID, req.Query.RPCAddress,
			req.Body.TimeoutHeight, req.Query.SimulateAndExecute, req.Query.BroadcastMode, messages...,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(3, err))
			return
		}

		txRes, err := ctx.QueryTxWithRetry(req.Query.RPCAddress, txResp.TxHash, req.Query.MaxQueryTries)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(4, err))
			return
		}
		if txRes == nil {
			err := fmt.Errorf("query result is nil for the transaction %s", txResp.TxHash)
			c.JSON(http.StatusInternalServerError, types.NewResponseError(4, err))
			return
		}
		if !txRes.TxResult.IsOK() {
			err := fmt.Errorf("transaction %s failed with the code %d", txResp.TxHash, txRes.TxResult.Code)
			c.JSON(http.StatusInternalServerError, types.NewResponseError(4, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(txRes))
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

		fromAddr := key.GetAddress()
		if !req.AuthzGranter.Empty() {
			fromAddr = req.AuthzGranter
		}

		var messages []sdk.Msg
		messages = append(messages, plantypes.NewMsgCreateRequest(fromAddr.Bytes(), req.Body.Duration, req.Body.Gigabytes, req.Prices))

		if !req.AuthzGranter.Empty() {
			execMsg := authz.NewMsgExec(key.GetAddress(), messages)
			messages = []sdk.Msg{&execMsg}
		}

		txResp, err := ctx.Tx(
			kr, key.GetName(), req.Query.Gas, req.Query.GasAdjustment, req.Query.GasPrices,
			req.Body.Fees, req.FeeGranter, req.Body.Memo, req.Body.SignMode, req.Query.ChainID, req.Query.RPCAddress,
			req.Body.TimeoutHeight, req.Query.SimulateAndExecute, req.Query.BroadcastMode, messages...,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(3, err))
			return
		}

		txRes, err := ctx.QueryTxWithRetry(req.Query.RPCAddress, txResp.TxHash, req.Query.MaxQueryTries)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(4, err))
			return
		}
		if txRes == nil {
			err := fmt.Errorf("query result is nil for the transaction %s", txResp.TxHash)
			c.JSON(http.StatusInternalServerError, types.NewResponseError(4, err))
			return
		}
		if !txRes.TxResult.IsOK() {
			err := fmt.Errorf("transaction %s failed with the code %d", txResp.TxHash, txRes.TxResult.Code)
			c.JSON(http.StatusInternalServerError, types.NewResponseError(4, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(txRes))
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

		fromAddr := key.GetAddress()
		if !req.AuthzGranter.Empty() {
			fromAddr = req.AuthzGranter
		}

		var messages []sdk.Msg
		messages = append(messages, plantypes.NewMsgUpdateStatusRequest(fromAddr.Bytes(), req.URI.ID, req.Status))

		if !req.AuthzGranter.Empty() {
			execMsg := authz.NewMsgExec(key.GetAddress(), messages)
			messages = []sdk.Msg{&execMsg}
		}

		txResp, err := ctx.Tx(
			kr, key.GetName(), req.Query.Gas, req.Query.GasAdjustment, req.Query.GasPrices,
			req.Body.Fees, req.FeeGranter, req.Body.Memo, req.Body.SignMode, req.Query.ChainID, req.Query.RPCAddress,
			req.Body.TimeoutHeight, req.Query.SimulateAndExecute, req.Query.BroadcastMode, messages...,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(3, err))
			return
		}

		txRes, err := ctx.QueryTxWithRetry(req.Query.RPCAddress, txResp.TxHash, req.Query.MaxQueryTries)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(4, err))
			return
		}
		if txRes == nil {
			err := fmt.Errorf("query result is nil for the transaction %s", txResp.TxHash)
			c.JSON(http.StatusInternalServerError, types.NewResponseError(4, err))
			return
		}
		if !txRes.TxResult.IsOK() {
			err := fmt.Errorf("transaction %s failed with the code %d", txResp.TxHash, txRes.TxResult.Code)
			c.JSON(http.StatusInternalServerError, types.NewResponseError(4, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(txRes))
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

		fromAddr := key.GetAddress()
		if !req.AuthzGranter.Empty() {
			fromAddr = req.AuthzGranter
		}

		var messages []sdk.Msg
		for i := 0; i < len(req.NodeAddresses); i++ {
			messages = append(messages, plantypes.NewMsgLinkNodeRequest(fromAddr.Bytes(), req.URI.ID, req.NodeAddresses[i]))
		}

		if !req.AuthzGranter.Empty() {
			execMsg := authz.NewMsgExec(key.GetAddress(), messages)
			messages = []sdk.Msg{&execMsg}
		}

		txResp, err := ctx.Tx(
			kr, key.GetName(), req.Query.Gas, req.Query.GasAdjustment, req.Query.GasPrices,
			req.Body.Fees, req.FeeGranter, req.Body.Memo, req.Body.SignMode, req.Query.ChainID, req.Query.RPCAddress,
			req.Body.TimeoutHeight, req.Query.SimulateAndExecute, req.Query.BroadcastMode, messages...,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(3, err))
			return
		}

		txRes, err := ctx.QueryTxWithRetry(req.Query.RPCAddress, txResp.TxHash, req.Query.MaxQueryTries)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(4, err))
			return
		}
		if txRes == nil {
			err := fmt.Errorf("query result is nil for the transaction %s", txResp.TxHash)
			c.JSON(http.StatusInternalServerError, types.NewResponseError(4, err))
			return
		}
		if !txRes.TxResult.IsOK() {
			err := fmt.Errorf("transaction %s failed with the code %d", txResp.TxHash, txRes.TxResult.Code)
			c.JSON(http.StatusInternalServerError, types.NewResponseError(4, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(txRes))
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

		fromAddr := key.GetAddress()
		if !req.AuthzGranter.Empty() {
			fromAddr = req.AuthzGranter
		}

		var messages []sdk.Msg
		messages = append(messages, plantypes.NewMsgUnlinkNodeRequest(fromAddr.Bytes(), req.URI.ID, req.NodeAddress))

		if !req.AuthzGranter.Empty() {
			execMsg := authz.NewMsgExec(key.GetAddress(), messages)
			messages = []sdk.Msg{&execMsg}
		}

		txResp, err := ctx.Tx(
			kr, key.GetName(), req.Query.Gas, req.Query.GasAdjustment, req.Query.GasPrices,
			req.Body.Fees, req.FeeGranter, req.Body.Memo, req.Body.SignMode, req.Query.ChainID, req.Query.RPCAddress,
			req.Body.TimeoutHeight, req.Query.SimulateAndExecute, req.Query.BroadcastMode, messages...,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(3, err))
			return
		}

		txRes, err := ctx.QueryTxWithRetry(req.Query.RPCAddress, txResp.TxHash, req.Query.MaxQueryTries)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(4, err))
			return
		}
		if txRes == nil {
			err := fmt.Errorf("query result is nil for the transaction %s", txResp.TxHash)
			c.JSON(http.StatusInternalServerError, types.NewResponseError(4, err))
			return
		}
		if !txRes.TxResult.IsOK() {
			err := fmt.Errorf("transaction %s failed with the code %d", txResp.TxHash, txRes.TxResult.Code)
			c.JSON(http.StatusInternalServerError, types.NewResponseError(4, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(txRes))
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

		fromAddr := key.GetAddress()
		if !req.AuthzGranter.Empty() {
			fromAddr = req.AuthzGranter
		}

		var messages []sdk.Msg
		messages = append(messages, nodetypes.NewMsgSubscribeRequest(fromAddr, req.NodeAddress, req.Body.Gigabytes, req.Body.Hours, req.Body.Denom))

		if !req.AuthzGranter.Empty() {
			execMsg := authz.NewMsgExec(key.GetAddress(), messages)
			messages = []sdk.Msg{&execMsg}
		}

		txResp, err := ctx.Tx(
			kr, key.GetName(), req.Query.Gas, req.Query.GasAdjustment, req.Query.GasPrices,
			req.Body.Fees, req.FeeGranter, req.Body.Memo, req.Body.SignMode, req.Query.ChainID, req.Query.RPCAddress,
			req.Body.TimeoutHeight, req.Query.SimulateAndExecute, req.Query.BroadcastMode, messages...,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(3, err))
			return
		}

		txRes, err := ctx.QueryTxWithRetry(req.Query.RPCAddress, txResp.TxHash, req.Query.MaxQueryTries)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(4, err))
			return
		}
		if txRes == nil {
			err := fmt.Errorf("query result is nil for the transaction %s", txResp.TxHash)
			c.JSON(http.StatusInternalServerError, types.NewResponseError(4, err))
			return
		}
		if !txRes.TxResult.IsOK() {
			err := fmt.Errorf("transaction %s failed with the code %d", txResp.TxHash, txRes.TxResult.Code)
			c.JSON(http.StatusInternalServerError, types.NewResponseError(4, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(txRes))
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

		fromAddr := key.GetAddress()
		if !req.AuthzGranter.Empty() {
			fromAddr = req.AuthzGranter
		}

		var messages []sdk.Msg
		messages = append(messages, plantypes.NewMsgSubscribeRequest(fromAddr, req.URI.ID, req.Body.Denom))

		if !req.AuthzGranter.Empty() {
			execMsg := authz.NewMsgExec(key.GetAddress(), messages)
			messages = []sdk.Msg{&execMsg}
		}

		txResp, err := ctx.Tx(
			kr, key.GetName(), req.Query.Gas, req.Query.GasAdjustment, req.Query.GasPrices,
			req.Body.Fees, req.FeeGranter, req.Body.Memo, req.Body.SignMode, req.Query.ChainID, req.Query.RPCAddress,
			req.Body.TimeoutHeight, req.Query.SimulateAndExecute, req.Query.BroadcastMode, messages...,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(3, err))
			return
		}

		txRes, err := ctx.QueryTxWithRetry(req.Query.RPCAddress, txResp.TxHash, req.Query.MaxQueryTries)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(4, err))
			return
		}
		if txRes == nil {
			err := fmt.Errorf("query result is nil for the transaction %s", txResp.TxHash)
			c.JSON(http.StatusInternalServerError, types.NewResponseError(4, err))
			return
		}
		if !txRes.TxResult.IsOK() {
			err := fmt.Errorf("transaction %s failed with the code %d", txResp.TxHash, txRes.TxResult.Code)
			c.JSON(http.StatusInternalServerError, types.NewResponseError(4, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(txRes))
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

		fromAddr := key.GetAddress()
		if !req.AuthzGranter.Empty() {
			fromAddr = req.AuthzGranter
		}

		var messages []sdk.Msg
		for i := 0; i < len(req.AccAddresses); i++ {
			messages = append(messages, subscriptiontypes.NewMsgAllocateRequest(fromAddr, req.URI.ID, req.AccAddresses[i], req.Bytes[i]))
		}

		if !req.AuthzGranter.Empty() {
			execMsg := authz.NewMsgExec(key.GetAddress(), messages)
			messages = []sdk.Msg{&execMsg}
		}

		txResp, err := ctx.Tx(
			kr, key.GetName(), req.Query.Gas, req.Query.GasAdjustment, req.Query.GasPrices,
			req.Body.Fees, req.FeeGranter, req.Body.Memo, req.Body.SignMode, req.Query.ChainID, req.Query.RPCAddress,
			req.Body.TimeoutHeight, req.Query.SimulateAndExecute, req.Query.BroadcastMode, messages...,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(3, err))
			return
		}

		txRes, err := ctx.QueryTxWithRetry(req.Query.RPCAddress, txResp.TxHash, req.Query.MaxQueryTries)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(4, err))
			return
		}
		if txRes == nil {
			err := fmt.Errorf("query result is nil for the transaction %s", txResp.TxHash)
			c.JSON(http.StatusInternalServerError, types.NewResponseError(4, err))
			return
		}
		if !txRes.TxResult.IsOK() {
			err := fmt.Errorf("transaction %s failed with the code %d", txResp.TxHash, txRes.TxResult.Code)
			c.JSON(http.StatusInternalServerError, types.NewResponseError(4, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(txRes))
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

		fromAddr := key.GetAddress()
		if !req.AuthzGranter.Empty() {
			fromAddr = req.AuthzGranter
		}

		var messages []sdk.Msg
		messages = append(messages, sessiontypes.NewMsgStartRequest(fromAddr, req.URI.ID, req.NodeAddress))

		if !req.AuthzGranter.Empty() {
			execMsg := authz.NewMsgExec(key.GetAddress(), messages)
			messages = []sdk.Msg{&execMsg}
		}

		txResp, err := ctx.Tx(
			kr, key.GetName(), req.Query.Gas, req.Query.GasAdjustment, req.Query.GasPrices,
			req.Body.Fees, req.FeeGranter, req.Body.Memo, req.Body.SignMode, req.Query.ChainID, req.Query.RPCAddress,
			req.Body.TimeoutHeight, req.Query.SimulateAndExecute, req.Query.BroadcastMode, messages...,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(3, err))
			return
		}

		txRes, err := ctx.QueryTxWithRetry(req.Query.RPCAddress, txResp.TxHash, req.Query.MaxQueryTries)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(4, err))
			return
		}
		if txRes == nil {
			err := fmt.Errorf("query result is nil for the transaction %s", txResp.TxHash)
			c.JSON(http.StatusInternalServerError, types.NewResponseError(4, err))
			return
		}
		if !txRes.TxResult.IsOK() {
			err := fmt.Errorf("transaction %s failed with the code %d", txResp.TxHash, txRes.TxResult.Code)
			c.JSON(http.StatusInternalServerError, types.NewResponseError(4, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(txRes))
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

		fromAddr := key.GetAddress()
		if !req.AuthzGranter.Empty() {
			fromAddr = req.AuthzGranter
		}

		var (
			dIndex   = 0
			messages []sdk.Msg
		)

		for i := 0; i < len(req.NodeAddresses) && dIndex < len(req.Body.Denoms); dIndex, i = dIndex+1, i+1 {
			messages = append(
				messages,
				nodetypes.NewMsgSubscribeRequest(fromAddr, req.NodeAddresses[i], req.Body.Gigabytes[i], req.Body.Hours[i], req.Body.Denoms[dIndex]),
			)
		}
		for i := 0; i < len(req.Body.PlanIDs) && dIndex < len(req.Body.Denoms); dIndex, i = dIndex+1, i+1 {
			messages = append(
				messages,
				plantypes.NewMsgSubscribeRequest(fromAddr, req.Body.PlanIDs[i], req.Body.Denoms[dIndex]),
			)
		}

		if !req.AuthzGranter.Empty() {
			execMsg := authz.NewMsgExec(key.GetAddress(), messages)
			messages = []sdk.Msg{&execMsg}
		}

		txResp, err := ctx.Tx(
			kr, key.GetName(), req.Query.Gas, req.Query.GasAdjustment, req.Query.GasPrices,
			req.Body.Fees, req.FeeGranter, req.Body.Memo, req.Body.SignMode, req.Query.ChainID, req.Query.RPCAddress,
			req.Body.TimeoutHeight, req.Query.SimulateAndExecute, req.Query.BroadcastMode, messages...,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(3, err))
			return
		}

		txRes, err := ctx.QueryTxWithRetry(req.Query.RPCAddress, txResp.TxHash, req.Query.MaxQueryTries)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(4, err))
			return
		}
		if txRes == nil {
			err := fmt.Errorf("query result is nil for the transaction %s", txResp.TxHash)
			c.JSON(http.StatusInternalServerError, types.NewResponseError(4, err))
			return
		}
		if !txRes.TxResult.IsOK() {
			err := fmt.Errorf("transaction %s failed with the code %d", txResp.TxHash, txRes.TxResult.Code)
			c.JSON(http.StatusInternalServerError, types.NewResponseError(4, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(txRes))
	}
}

func HandlerTxSubscriptionCancel(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := requests.NewRequestTxSubscriptionCancel(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewResponseError(1, err))
			return
		}

		kr, key, err := utils.NewInMemoryKey(req.Body.Mnemonic, req.Query.CoinType, req.Query.Account, req.Query.Index, req.Body.BIP39Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(2, err))
			return
		}

		fromAddr := key.GetAddress()
		if !req.AuthzGranter.Empty() {
			fromAddr = req.AuthzGranter
		}

		var messages []sdk.Msg
		for i := 0; i < len(req.Body.IDs); i++ {
			messages = append(messages, subscriptiontypes.NewMsgCancelRequest(fromAddr, req.Body.IDs[i]))
		}

		if !req.AuthzGranter.Empty() {
			execMsg := authz.NewMsgExec(key.GetAddress(), messages)
			messages = []sdk.Msg{&execMsg}
		}

		txResp, err := ctx.Tx(
			kr, key.GetName(), req.Query.Gas, req.Query.GasAdjustment, req.Query.GasPrices,
			req.Body.Fees, req.FeeGranter, req.Body.Memo, req.Body.SignMode, req.Query.ChainID, req.Query.RPCAddress,
			req.Body.TimeoutHeight, req.Query.SimulateAndExecute, req.Query.BroadcastMode, messages...,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(3, err))
			return
		}

		txRes, err := ctx.QueryTxWithRetry(req.Query.RPCAddress, txResp.TxHash, req.Query.MaxQueryTries)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(4, err))
			return
		}
		if txRes == nil {
			err := fmt.Errorf("query result is nil for the transaction %s", txResp.TxHash)
			c.JSON(http.StatusInternalServerError, types.NewResponseError(4, err))
			return
		}
		if !txRes.TxResult.IsOK() {
			err := fmt.Errorf("transaction %s failed with the code %d", txResp.TxHash, txRes.TxResult.Code)
			c.JSON(http.StatusInternalServerError, types.NewResponseError(4, err))
			return
		}

		c.JSON(http.StatusOK, types.NewResponseResult(txRes))
	}
}
