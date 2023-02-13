package requests

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gin-gonic/gin"
	hubtypes "github.com/sentinel-official/hub/types"
)

type RequestTxBankSend struct {
	GasPrices    sdk.DecCoins
	ToAccAddress sdk.AccAddress
	Amount       sdk.Coins

	Query struct {
		BroadcastMode string `form:"broadcast_mode,default=block" binding:"oneof=async block sync"`
		RPCAddress    string `form:"rpc_address" binding:"required"`
		ChainID       string `form:"chain_id,default=sentinelhub-2" binding:"required"`
		CoinType      uint32 `form:"coin_type,default=118"`
		Account       uint32 `form:"account"`
		Index         uint32 `form:"index"`
		Gas           uint64 `form:"gas,default=200000"`
		GasPrices     string `form:"gas_prices,default=0.1udvpn"`
	}
	Body struct {
		Mnemonic      string `json:"mnemonic" binding:"required"`
		BIP39Password string `json:"bip39_password"`
		Memo          string `json:"memo"`
		ToAccAddress  string `json:"to_acc_address" binding:"required"`
		Amount        string `json:"amount" binding:"required"`
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
	GasPrices   sdk.DecCoins
	NodeAddress hubtypes.NodeAddress
	Deposit     sdk.Coin

	URI struct {
		NodeAddress string `uri:"node_address"`
	}
	Query struct {
		BroadcastMode string `form:"broadcast_mode,default=block" binding:"oneof=async block sync"`
		RPCAddress    string `form:"rpc_address" binding:"required"`
		ChainID       string `form:"chain_id,default=sentinelhub-2" binding:"required"`
		CoinType      uint32 `form:"coin_type,default=118"`
		Account       uint32 `form:"account"`
		Index         uint32 `form:"index"`
		Gas           uint64 `form:"gas,default=200000"`
		GasPrices     string `form:"gas_prices,default=0.1udvpn"`
	}
	Body struct {
		Mnemonic      string `json:"mnemonic" binding:"required"`
		BIP39Password string `json:"bip39_password"`
		Memo          string `json:"memo"`
		Deposit       string `json:"deposit" binding:"required"`
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
	GasPrices sdk.DecCoins

	URI struct {
		ID uint64 `uri:"id" binding:"gt=0"`
	}
	Query struct {
		BroadcastMode string `form:"broadcast_mode,default=block" binding:"oneof=async block sync"`
		RPCAddress    string `form:"rpc_address" binding:"required"`
		ChainID       string `form:"chain_id,default=sentinelhub-2" binding:"required"`
		CoinType      uint32 `form:"coin_type,default=118"`
		Account       uint32 `form:"account"`
		Index         uint32 `form:"index"`
		Gas           uint64 `form:"gas,default=200000"`
		GasPrices     string `form:"gas_prices,default=0.1udvpn"`
	}
	Body struct {
		Mnemonic      string `json:"mnemonic" binding:"required"`
		BIP39Password string `json:"bip39_password"`
		Memo          string `json:"memo"`
		Denom         string `json:"denom" binding:"required"`
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

	req.GasPrices, err = sdk.ParseDecCoins(req.Query.GasPrices)
	if err != nil {
		return nil, err
	}

	return req, err
}

type RequestTxStartSession struct {
	GasPrices   sdk.DecCoins
	NodeAddress hubtypes.NodeAddress

	URI struct {
		ID          uint64 `uri:"id" binding:"gt=0"`
		NodeAddress string `uri:"node_address"`
	}
	Query struct {
		BroadcastMode string `form:"broadcast_mode,default=block" binding:"oneof=async block sync"`
		RPCAddress    string `form:"rpc_address" binding:"required"`
		ChainID       string `form:"chain_id,default=sentinelhub-2" binding:"required"`
		CoinType      uint32 `form:"coin_type,default=118"`
		Account       uint32 `form:"account"`
		Index         uint32 `form:"index"`
		Gas           uint64 `form:"gas,default=200000"`
		GasPrices     string `form:"gas_prices,default=0.1udvpn"`
	}
	Body struct {
		Mnemonic      string `json:"mnemonic" binding:"required"`
		BIP39Password string `json:"bip39_password"`
		Memo          string `json:"memo"`
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

	req.GasPrices, err = sdk.ParseDecCoins(req.Query.GasPrices)
	if err != nil {
		return nil, err
	}

	return req, err
}
