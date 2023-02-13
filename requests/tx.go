package requests

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gin-gonic/gin"
	hubtypes "github.com/sentinel-official/hub/types"
)

type (
	TxQuery struct {
		BroadcastMode      string  `form:"broadcast_mode,default=block" binding:"oneof=async block sync"`
		ChainID            string  `form:"chain_id,default=sentinelhub-2" binding:"required"`
		CoinType           uint32  `form:"coin_type,default=118"`
		Account            uint32  `form:"account"`
		Index              uint32  `form:"index"`
		GasAdjustment      float64 `form:"gas_adjustment,default=1.25" binding:"gt=0"`
		GasPrices          string  `form:"gas_prices,default=0.1udvpn"`
		Gas                uint64  `form:"gas,default=200000" binding:"gt=0"`
		RPCAddress         string  `form:"rpc_address,default=https://rpc.sentinel.co:443" binding:"required"`
		SimulateAndExecute bool    `form:"simulate_and_execute,default=true"`
	}
	TxBody struct {
		BIP39Password string `json:"bip39_password"`
		FeeGranter    string `json:"fee_granter"`
		Fees          string `json:"fees"`
		Memo          string `json:"memo"`
		SignMode      string `json:"sign_mode"`
		TimeoutHeight uint64 `json:"timeout_height"`
		Mnemonic      string `json:"mnemonic" binding:"required"`
	}
)

type RequestTxBankSend struct {
	FeeGranter   sdk.AccAddress
	GasPrices    sdk.DecCoins
	ToAccAddress sdk.AccAddress
	Amount       sdk.Coins

	Query TxQuery
	Body  struct {
		TxBody
		ToAccAddress string `json:"to_acc_address" binding:"required"`
		Amount       string `json:"amount" binding:"required"`
	}
}

func NewRequestTxBankSend(c *gin.Context) (req *RequestTxBankSend, err error) {
	req = &RequestTxBankSend{}
	if err = c.ShouldBindQuery(&req.Query); err != nil {
		return nil, err
	}
	if err = c.ShouldBindJSON(&req.Body); err != nil {
		return nil, err
	}

	if req.Body.FeeGranter != "" {
		req.FeeGranter, err = sdk.AccAddressFromBech32(req.Body.FeeGranter)
		if err != nil {
			return nil, err
		}
	}

	req.GasPrices, err = sdk.ParseDecCoins(req.Query.GasPrices)
	if err != nil {
		return nil, err
	}

	req.ToAccAddress, err = sdk.AccAddressFromBech32(req.Body.ToAccAddress)
	if err != nil {
		return nil, err
	}

	req.Amount, err = sdk.ParseCoinsNormalized(req.Body.Amount)
	if err != nil {
		return nil, err
	}

	return req, err
}

type RequestTxSubscribeToNode struct {
	FeeGranter  sdk.AccAddress
	GasPrices   sdk.DecCoins
	NodeAddress hubtypes.NodeAddress
	Deposit     sdk.Coin

	URI struct {
		NodeAddress string `uri:"node_address"`
	}
	Query TxQuery
	Body  struct {
		TxBody
		Deposit string `json:"deposit" binding:"required"`
	}
}

func NewRequestTxSubscribeToNode(c *gin.Context) (req *RequestTxSubscribeToNode, err error) {
	req = &RequestTxSubscribeToNode{}
	if err = c.ShouldBindUri(&req.URI); err != nil {
		return nil, err
	}
	if err = c.ShouldBindQuery(&req.Query); err != nil {
		return nil, err
	}
	if err = c.ShouldBindJSON(&req.Body); err != nil {
		return nil, err
	}

	req.NodeAddress, err = hubtypes.NodeAddressFromBech32(req.URI.NodeAddress)
	if err != nil {
		return nil, err
	}

	if req.Body.FeeGranter != "" {
		req.FeeGranter, err = sdk.AccAddressFromBech32(req.Body.FeeGranter)
		if err != nil {
			return nil, err
		}
	}

	req.GasPrices, err = sdk.ParseDecCoins(req.Query.GasPrices)
	if err != nil {
		return nil, err
	}

	req.Deposit, err = sdk.ParseCoinNormalized(req.Body.Deposit)
	if err != nil {
		return nil, err
	}

	return req, err
}

type RequestTxSubscribeToPlan struct {
	FeeGranter sdk.AccAddress
	GasPrices  sdk.DecCoins

	URI struct {
		ID uint64 `uri:"id" binding:"gt=0"`
	}
	Query TxQuery
	Body  struct {
		TxBody
		Denom string `json:"denom" binding:"required"`
	}
}

func NewRequestTxSubscribeToPlan(c *gin.Context) (req *RequestTxSubscribeToPlan, err error) {
	req = &RequestTxSubscribeToPlan{}
	if err = c.ShouldBindUri(&req.URI); err != nil {
		return nil, err
	}
	if err = c.ShouldBindQuery(&req.Query); err != nil {
		return nil, err
	}
	if err = c.ShouldBindJSON(&req.Body); err != nil {
		return nil, err
	}

	if req.Body.FeeGranter != "" {
		req.FeeGranter, err = sdk.AccAddressFromBech32(req.Body.FeeGranter)
		if err != nil {
			return nil, err
		}
	}

	req.GasPrices, err = sdk.ParseDecCoins(req.Query.GasPrices)
	if err != nil {
		return nil, err
	}

	return req, err
}

type RequestTxStartSession struct {
	FeeGranter  sdk.AccAddress
	GasPrices   sdk.DecCoins
	NodeAddress hubtypes.NodeAddress

	URI struct {
		ID          uint64 `uri:"id" binding:"gt=0"`
		NodeAddress string `uri:"node_address"`
	}
	Query TxQuery
	Body  struct {
		TxBody
	}
}

func NewRequestTxStartSession(c *gin.Context) (req *RequestTxStartSession, err error) {
	req = &RequestTxStartSession{}
	if err = c.ShouldBindUri(&req.URI); err != nil {
		return nil, err
	}
	if err = c.ShouldBindQuery(&req.Query); err != nil {
		return nil, err
	}
	if err = c.ShouldBindJSON(&req.Body); err != nil {
		return nil, err
	}

	req.NodeAddress, err = hubtypes.NodeAddressFromBech32(req.URI.NodeAddress)
	if err != nil {
		return nil, err
	}

	if req.Body.FeeGranter != "" {
		req.FeeGranter, err = sdk.AccAddressFromBech32(req.Body.FeeGranter)
		if err != nil {
			return nil, err
		}
	}

	req.GasPrices, err = sdk.ParseDecCoins(req.Query.GasPrices)
	if err != nil {
		return nil, err
	}

	return req, err
}
