package requests

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gin-gonic/gin"
	hubtypes "github.com/sentinel-official/hub/types"
)

type RequestAddSessionKey struct {
	FeeGranter  sdk.AccAddress
	GasPrices   sdk.DecCoins
	NodeAddress hubtypes.NodeAddress

	URI struct {
		ID          uint64 `uri:"id"`
		NodeAddress string `uri:"node_address"`
	}
	Query TxQuery
	Body  struct {
		TxBody
	}
}

func NewRequestAddSessionKey(c *gin.Context) (req *RequestAddSessionKey, err error) {
	req = &RequestAddSessionKey{}
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
