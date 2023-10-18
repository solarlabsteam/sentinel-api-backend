package requests

import (
	"fmt"
	"time"

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
		AuthzGranter  string `json:"authz_granter"`
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
	AuthzGranter   sdk.AccAddress
	FeeGranter     sdk.AccAddress
	GasPrices      sdk.DecCoins
	ToAccAddresses []sdk.AccAddress
	Amounts        []sdk.Coins

	Query TxQuery
	Body  struct {
		TxBody
		ToAccAddresses []string `json:"to_acc_addresses" binding:"required"`
		Amounts        []string `json:"amounts" binding:"required"`
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

	if req.Body.AuthzGranter != "" {
		req.AuthzGranter, err = sdk.AccAddressFromBech32(req.Body.AuthzGranter)
		if err != nil {
			return nil, err
		}
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

	if len(req.Body.ToAccAddresses) != len(req.Body.Amounts) {
		return nil, fmt.Errorf("to_acc_addresses length must be equal to the amounts length")
	}

	for _, s := range req.Body.ToAccAddresses {
		v, err := sdk.AccAddressFromBech32(s)
		if err != nil {
			return nil, err
		}

		req.ToAccAddresses = append(req.ToAccAddresses, v)
	}

	for _, s := range req.Body.Amounts {
		v, err := sdk.ParseCoinsNormalized(s)
		if err != nil {
			return nil, err
		}

		req.Amounts = append(req.Amounts, v)
	}

	return req, err
}

type RequestTxNodeSubscribe struct {
	AuthzGranter sdk.AccAddress
	FeeGranter   sdk.AccAddress
	GasPrices    sdk.DecCoins
	NodeAddress  hubtypes.NodeAddress

	URI struct {
		NodeAddress string `uri:"node_address"`
	}
	Query TxQuery
	Body  struct {
		TxBody
		Gigabytes int64  `json:"gigabytes"`
		Hours     int64  `json:"hours"`
		Denom     string `json:"denom" binding:"required"`
	}
}

func NewRequestTxNodeSubscribe(c *gin.Context) (req *RequestTxNodeSubscribe, err error) {
	req = &RequestTxNodeSubscribe{}
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

	if req.Body.AuthzGranter != "" {
		req.AuthzGranter, err = sdk.AccAddressFromBech32(req.Body.AuthzGranter)
		if err != nil {
			return nil, err
		}
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

type RequestTxPlanCreate struct {
	AuthzGranter sdk.AccAddress
	FeeGranter   sdk.AccAddress
	GasPrices    sdk.DecCoins
	Prices       sdk.Coins

	Query TxQuery
	Body  struct {
		TxBody
		Duration  time.Duration `json:"duration" binding:"required"`
		Gigabytes int64         `json:"gigabytes" binding:"required"`
		Prices    string        `json:"prices" binding:"required"`
	}
}

func NewRequestTxPlanCreate(c *gin.Context) (req *RequestTxPlanCreate, err error) {
	req = &RequestTxPlanCreate{}
	if err = c.ShouldBindQuery(&req.Query); err != nil {
		return nil, err
	}
	if err = c.ShouldBindJSON(&req.Body); err != nil {
		return nil, err
	}

	if req.Body.AuthzGranter != "" {
		req.AuthzGranter, err = sdk.AccAddressFromBech32(req.Body.AuthzGranter)
		if err != nil {
			return nil, err
		}
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

	req.Prices, err = sdk.ParseCoinsNormalized(req.Body.Prices)
	if err != nil {
		return nil, err
	}

	return req, err
}

type RequestTxPlanUpdateStatus struct {
	AuthzGranter sdk.AccAddress
	FeeGranter   sdk.AccAddress
	GasPrices    sdk.DecCoins
	Status       hubtypes.Status

	URI struct {
		ID uint64 `uri:"id" binding:"gt=0"`
	}
	Query TxQuery
	Body  struct {
		TxBody
		Status string `json:"status" binding:"required"`
	}
}

func NewRequestTxPlanUpdateStatus(c *gin.Context) (req *RequestTxPlanUpdateStatus, err error) {
	req = &RequestTxPlanUpdateStatus{}
	if err = c.ShouldBindUri(&req.URI); err != nil {
		return nil, err
	}
	if err = c.ShouldBindQuery(&req.Query); err != nil {
		return nil, err
	}
	if err = c.ShouldBindJSON(&req.Body); err != nil {
		return nil, err
	}

	if req.Body.AuthzGranter != "" {
		req.AuthzGranter, err = sdk.AccAddressFromBech32(req.Body.AuthzGranter)
		if err != nil {
			return nil, err
		}
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

	req.Status = hubtypes.StatusFromString(req.Body.Status)

	return req, err
}

type RequestTxPlanLinkNode struct {
	AuthzGranter  sdk.AccAddress
	FeeGranter    sdk.AccAddress
	GasPrices     sdk.DecCoins
	NodeAddresses []hubtypes.NodeAddress

	URI struct {
		ID uint64 `uri:"id" binding:"gt=0"`
	}
	Query TxQuery
	Body  struct {
		TxBody
		NodeAddresses []string `json:"node_addresses" binding:"required"`
	}
}

func NewRequestTxPlanLinkNode(c *gin.Context) (req *RequestTxPlanLinkNode, err error) {
	req = &RequestTxPlanLinkNode{}
	if err = c.ShouldBindUri(&req.URI); err != nil {
		return nil, err
	}
	if err = c.ShouldBindQuery(&req.Query); err != nil {
		return nil, err
	}
	if err = c.ShouldBindJSON(&req.Body); err != nil {
		return nil, err
	}

	if req.Body.AuthzGranter != "" {
		req.AuthzGranter, err = sdk.AccAddressFromBech32(req.Body.AuthzGranter)
		if err != nil {
			return nil, err
		}
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

	for _, s := range req.Body.NodeAddresses {
		v, err := hubtypes.NodeAddressFromBech32(s)
		if err != nil {
			return nil, err
		}

		req.NodeAddresses = append(req.NodeAddresses, v)
	}

	return req, err
}

type RequestTxPlanUnlinkNode struct {
	AuthzGranter sdk.AccAddress
	FeeGranter   sdk.AccAddress
	GasPrices    sdk.DecCoins
	NodeAddress  hubtypes.NodeAddress

	URI struct {
		ID          uint64 `uri:"id" binding:"gt=0"`
		NodeAddress string `uri:"node_address"`
	}
	Query TxQuery
	Body  struct {
		TxBody
	}
}

func NewRequestTxPlanUnlinkNode(c *gin.Context) (req *RequestTxPlanUnlinkNode, err error) {
	req = &RequestTxPlanUnlinkNode{}
	if err = c.ShouldBindUri(&req.URI); err != nil {
		return nil, err
	}
	if err = c.ShouldBindQuery(&req.Query); err != nil {
		return nil, err
	}
	if err = c.ShouldBindJSON(&req.Body); err != nil {
		return nil, err
	}

	if req.Body.AuthzGranter != "" {
		req.AuthzGranter, err = sdk.AccAddressFromBech32(req.Body.AuthzGranter)
		if err != nil {
			return nil, err
		}
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

	req.NodeAddress, err = hubtypes.NodeAddressFromBech32(req.URI.NodeAddress)
	if err != nil {
		return nil, err
	}

	return req, err
}

type RequestTxPlanSubscribe struct {
	AuthzGranter sdk.AccAddress
	FeeGranter   sdk.AccAddress
	GasPrices    sdk.DecCoins

	URI struct {
		ID uint64 `uri:"id" binding:"gt=0"`
	}
	Query TxQuery
	Body  struct {
		TxBody
		Denom string `json:"denom" binding:"required"`
	}
}

func NewRequestTxPlanSubscribe(c *gin.Context) (req *RequestTxPlanSubscribe, err error) {
	req = &RequestTxPlanSubscribe{}
	if err = c.ShouldBindUri(&req.URI); err != nil {
		return nil, err
	}
	if err = c.ShouldBindQuery(&req.Query); err != nil {
		return nil, err
	}
	if err = c.ShouldBindJSON(&req.Body); err != nil {
		return nil, err
	}

	if req.Body.AuthzGranter != "" {
		req.AuthzGranter, err = sdk.AccAddressFromBech32(req.Body.AuthzGranter)
		if err != nil {
			return nil, err
		}
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

type RequestTxSubscriptionAllocate struct {
	AuthzGranter sdk.AccAddress
	FeeGranter   sdk.AccAddress
	GasPrices    sdk.DecCoins
	AccAddresses []sdk.AccAddress
	Bytes        []sdk.Int

	URI struct {
		ID uint64 `uri:"id" binding:"gt=0"`
	}
	Query TxQuery
	Body  struct {
		TxBody
		AccAddresses []string `json:"acc_addresses" binding:"required"`
		Bytes        []int64  `json:"bytes" binding:"required"`
	}
}

func NewRequestTxSubscriptionAllocate(c *gin.Context) (req *RequestTxSubscriptionAllocate, err error) {
	req = &RequestTxSubscriptionAllocate{}
	if err = c.ShouldBindUri(&req.URI); err != nil {
		return nil, err
	}
	if err = c.ShouldBindQuery(&req.Query); err != nil {
		return nil, err
	}
	if err = c.ShouldBindJSON(&req.Body); err != nil {
		return nil, err
	}

	if req.Body.AuthzGranter != "" {
		req.AuthzGranter, err = sdk.AccAddressFromBech32(req.Body.AuthzGranter)
		if err != nil {
			return nil, err
		}
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

	if len(req.Body.AccAddresses) != len(req.Body.Bytes) {
		return nil, fmt.Errorf("acc_addresses length must be equal to the bytes length")
	}

	for _, s := range req.Body.AccAddresses {
		v, err := sdk.AccAddressFromBech32(s)
		if err != nil {
			return nil, err
		}

		req.AccAddresses = append(req.AccAddresses, v)
	}

	for _, i := range req.Body.Bytes {
		req.Bytes = append(req.Bytes, sdk.NewInt(i))
	}

	return req, err
}

type RequestTxSessionStart struct {
	AuthzGranter sdk.AccAddress
	FeeGranter   sdk.AccAddress
	GasPrices    sdk.DecCoins
	NodeAddress  hubtypes.NodeAddress

	URI struct {
		ID          uint64 `uri:"id" binding:"gt=0"`
		NodeAddress string `uri:"node_address"`
	}
	Query TxQuery
	Body  struct {
		TxBody
	}
}

func NewRequestTxSessionStart(c *gin.Context) (req *RequestTxSessionStart, err error) {
	req = &RequestTxSessionStart{}
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

	if req.Body.AuthzGranter != "" {
		req.AuthzGranter, err = sdk.AccAddressFromBech32(req.Body.AuthzGranter)
		if err != nil {
			return nil, err
		}
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

type RequestTxSubscribe struct {
	AuthzGranter  sdk.AccAddress
	FeeGranter    sdk.AccAddress
	GasPrices     sdk.DecCoins
	NodeAddresses []hubtypes.NodeAddress

	Query TxQuery
	Body  struct {
		TxBody
		Denoms        []string `json:"denoms"`
		Gigabytes     []int64  `json:"gigabytes"`
		Hours         []int64  `json:"hours"`
		NodeAddresses []string `json:"node_addresses"`
		PlanIDs       []uint64 `json:"plan_ids"`
	}
}

func NewRequestTxSubscribe(c *gin.Context) (req *RequestTxSubscribe, err error) {
	req = &RequestTxSubscribe{}
	if err = c.ShouldBindQuery(&req.Query); err != nil {
		return nil, err
	}
	if err = c.ShouldBindJSON(&req.Body); err != nil {
		return nil, err
	}

	if len(req.Body.Denoms) != len(req.Body.NodeAddresses)+len(req.Body.PlanIDs) {
		return nil, fmt.Errorf("invalid denoms length")
	}
	if len(req.Body.Gigabytes) != len(req.Body.NodeAddresses) {
		return nil, fmt.Errorf("invalid gigabytes length")
	}
	if len(req.Body.Hours) != len(req.Body.NodeAddresses) {
		return nil, fmt.Errorf("invalid hours length")
	}

	for _, s := range req.Body.NodeAddresses {
		addr, err := hubtypes.NodeAddressFromBech32(s)
		if err != nil {
			return nil, err
		}

		req.NodeAddresses = append(req.NodeAddresses, addr)
	}

	if req.Body.AuthzGranter != "" {
		req.AuthzGranter, err = sdk.AccAddressFromBech32(req.Body.AuthzGranter)
		if err != nil {
			return nil, err
		}
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
