package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/solarlabsteam/sentinel-api-backend/context"
	"github.com/solarlabsteam/sentinel-api-backend/requests"
	"github.com/solarlabsteam/sentinel-api-backend/responses"
	"github.com/solarlabsteam/sentinel-api-backend/types"
	"github.com/solarlabsteam/sentinel-api-backend/utils"
)

func HandlerGenerateSignature(_ context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := requests.NewRequestGenerateSignature(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewResponseError(1, err))
			return
		}

		kr, key, err := utils.NewInMemoryKey(req.Body.Mnemonic, req.Query.CoinType, req.Query.Account, req.Query.Index, req.Body.BIP39Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(2, err))
			return
		}

		signature, _, err := kr.Sign(key.GetName(), req.Message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(3, err))
			return
		}

		result := &responses.ResponseGenerateSignature{
			Signature: signature,
		}

		c.JSON(http.StatusOK, types.NewResponseResult(result))
	}
}
