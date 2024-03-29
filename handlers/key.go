package handlers

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/transport/http/jsonrpc"
	"github.com/hashicorp/go-uuid"
	sessiontypes "github.com/sentinel-official/hub/x/session/types"

	"github.com/solarlabsteam/sentinel-api-backend/context"
	"github.com/solarlabsteam/sentinel-api-backend/requests"
	"github.com/solarlabsteam/sentinel-api-backend/responses"
	"github.com/solarlabsteam/sentinel-api-backend/types"
	"github.com/solarlabsteam/sentinel-api-backend/utils"
	eventutils "github.com/solarlabsteam/sentinel-api-backend/utils/event"
)

func fetchNodeInfo(remoteURL string) (map[string]interface{}, error) {
	endpoint, err := url.JoinPath(remoteURL, "status")
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: 15 * time.Second,
	}

	resp, err := client.Get(endpoint)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err = resp.Body.Close(); err != nil {
			panic(err)
		}
	}()

	var body types.Response
	if err = json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, err
	}

	return body.Result.(map[string]interface{}), nil
}

func HandlerAddSessionKey(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := requests.NewRequestAddSessionKey(c)
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
			accAddress = key.GetAddress()
			messages   []sdk.Msg
		)

		rSession, err := ctx.QueryActiveSession(req.Query.RPCAddress, accAddress)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(3, err))
			return
		}

		if rSession != nil {
			messages = append(
				messages,
				sessiontypes.NewMsgEndRequest(
					accAddress,
					rSession.ID,
					0,
				),
			)
		}

		messages = append(
			messages,
			sessiontypes.NewMsgStartRequest(
				accAddress,
				req.URI.ID,
				req.NodeAddress,
			),
		)

		txResp, err := ctx.Tx(
			kr, key.GetName(), req.Query.Gas, req.Query.GasAdjustment, req.Query.GasPrices,
			req.Body.Fees, req.FeeGranter, req.Body.Memo, req.Body.SignMode, req.Query.ChainID, req.Query.RPCAddress,
			req.Body.TimeoutHeight, req.Query.SimulateAndExecute, req.Query.BroadcastMode, messages...,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(4, err))
			return
		}

		txRes, err := ctx.QueryTxWithRetry(req.Query.RPCAddress, txResp.TxHash, req.Query.MaxQueryTries)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(5, err))
			return
		}
		if txRes == nil {
			err := fmt.Errorf("query result is nil for the transaction %s", txResp.TxHash)
			c.JSON(http.StatusInternalServerError, types.NewResponseError(5, err))
			return
		}
		if !txRes.TxResult.IsOK() {
			err := fmt.Errorf("transaction %s failed with the code %d", txResp.TxHash, txRes.TxResult.Code)
			c.JSON(http.StatusInternalServerError, types.NewResponseError(5, err))
			return
		}

		sessionID, err := eventutils.GetSessionIDFromABCIEvents(txRes.TxResult.Events)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(6, err))
			return
		}

		rNode, err := ctx.QueryNode(req.Query.RPCAddress, req.NodeAddress)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(7, err))
			return
		}

		rNodeInfo, err := fetchNodeInfo(rNode.RemoteURL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(8, err))
			return
		}

		var (
			nodeType     = rNodeInfo["type"].(float64)
			clientKey    string
			wgPrivateKey *types.Key
			uid          []byte
		)

		if nodeType == 1 {
			wgPrivateKey, err = types.NewPrivateKey()
			if err != nil {
				return
			}

			clientKey = wgPrivateKey.Public().String()
		} else if nodeType == 2 {
			uid, err = uuid.GenerateRandomBytes(16)
			if err != nil {
				return
			}

			clientKey = base64.StdEncoding.EncodeToString(append([]byte{0x01}, uid...))
		} else {
			err := fmt.Errorf("unknown node type %f", nodeType)
			c.JSON(http.StatusBadRequest, types.NewResponseError(9, err))
			return
		}

		signature, _, err := kr.Sign(key.GetName(), sdk.Uint64ToBigEndian(sessionID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(10, err))
			return
		}

		nReq, err := json.Marshal(
			map[string]interface{}{
				"key":       clientKey,
				"signature": signature,
			},
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(11, err))
			return
		}

		endpoint, err := url.JoinPath(rNode.RemoteURL, fmt.Sprintf("/accounts/%s/sessions/%d", accAddress, sessionID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(12, err))
			return
		}

		var (
			body   types.Response
			client = &http.Client{
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{
						InsecureSkipVerify: true,
					},
				},
				Timeout: 15 * time.Second,
			}
		)

		tStart := time.Now()

		resp, err := client.Post(endpoint, jsonrpc.ContentType, bytes.NewBuffer(nReq))
		if err != nil {
			err := fmt.Errorf("error %s; time took %s", err, time.Since(tStart))
			c.JSON(http.StatusInternalServerError, types.NewResponseError(13, err))
			return
		}

		defer func() {
			if err = resp.Body.Close(); err != nil {
				panic(err)
			}
		}()

		if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
			c.JSON(http.StatusInternalServerError, types.NewResponseError(14, err))
			return
		}
		if body.Error != nil {
			err := fmt.Errorf("node responded with code %d and message %s", body.Error.Code, body.Error.Message)
			c.JSON(http.StatusInternalServerError, types.NewResponseError(15, err))
			return
		}

		result := &responses.ResponseAddSessionKey{
			NodeType: nodeType,
			Result:   body.Result.(string),
		}

		if nodeType == 1 {
			result.PrivateKey = wgPrivateKey.String()
		} else if nodeType == 2 {
			uid, _ := uuid.FormatUUID(uid)
			result.UID = uid
		}

		c.JSON(http.StatusOK, types.NewResponseResult(result))
	}
}
