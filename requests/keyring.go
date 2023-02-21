package requests

import (
	"encoding/base64"

	"github.com/gin-gonic/gin"
)

type RequestGenerateSignature struct {
	Message []byte
	Query   struct {
		CoinType uint32 `form:"coin_type,default=118"`
		Account  uint32 `form:"account"`
		Index    uint32 `form:"index"`
	}
	Body struct {
		BIP39Password string `json:"bip39_password"`
		Mnemonic      string `json:"mnemonic" binding:"required"`
		Message       string `json:"message"`
	}
}

func NewRequestGenerateSignature(c *gin.Context) (req *RequestGenerateSignature, err error) {
	req = &RequestGenerateSignature{}
	if err = c.ShouldBindQuery(&req.Query); err != nil {
		return nil, err
	}
	if err = c.ShouldBindJSON(&req.Body); err != nil {
		return nil, err
	}

	req.Message, err = base64.StdEncoding.DecodeString(req.Body.Message)
	if err != nil {
		return nil, err
	}

	return req, nil
}
