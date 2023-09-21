package context

import (
	"sync"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
)

func (c Context) Tx(
	kr keyring.Keyring, from string, gas uint64, gasAdjustment float64, gasPrices string,
	fees string, feeGranter sdk.AccAddress, memo, signModeStr, chainID, rpcAddress string,
	timeoutHeight uint64, simulateAndExecute bool, broadcastMode string, messages ...sdk.Msg,
) (result *sdk.TxResponse, err error) {
	c.Client, err = rpchttp.New(rpcAddress, "/websocket")
	if err != nil {
		return nil, err
	}

	key, err := kr.Key(from)
	if err != nil {
		return nil, err
	}

	var (
		accAddr       = key.GetAddress()
		bech32AccAddr = accAddr.String()
	)

	if _, ok := c.mutex[bech32AccAddr]; !ok {
		c.mutex[bech32AccAddr] = &sync.Mutex{}
	}

	c.mutex[bech32AccAddr].Lock()
	defer c.mutex[bech32AccAddr].Unlock()

	account, err := c.QueryAccount(rpcAddress, accAddr)
	if err != nil {
		return nil, err
	}

	c.BroadcastMode = broadcastMode
	c.ChainID = chainID
	c.FeeGranter = feeGranter
	c.FromName = from
	c.Keyring = kr
	c.NodeURI = rpcAddress
	c.SignModeStr = signModeStr
	c.Simulate = false
	c.SkipConfirm = true

	signMode := signing.SignMode_SIGN_MODE_UNSPECIFIED
	switch signModeStr {
	case flags.SignModeDirect:
		signMode = signing.SignMode_SIGN_MODE_DIRECT
	case flags.SignModeLegacyAminoJSON:
		signMode = signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON
	}

	txf := tx.Factory{}.
		WithTxConfig(c.TxConfig).
		WithAccountRetriever(c.AccountRetriever).
		WithKeybase(kr).
		WithChainID(chainID).
		WithGas(gas).
		WithSimulateAndExecute(simulateAndExecute).
		WithAccountNumber(account.GetAccountNumber()).
		WithSequence(account.GetSequence()).
		WithTimeoutHeight(timeoutHeight).
		WithGasAdjustment(gasAdjustment).
		WithMemo(memo).
		WithSignMode(signMode).
		WithFees(fees).
		WithGasPrices(gasPrices)

	if txf.SimulateAndExecute() {
		_, adjusted, err := tx.CalculateGas(c, txf, messages...)
		if err != nil {
			return nil, err
		}

		txf = txf.WithGas(adjusted)
	}

	txb, err := tx.BuildUnsignedTx(txf, messages...)
	if err != nil {
		return nil, err
	}

	txb.SetFeeGranter(c.GetFeeGranterAddress())
	if err = tx.Sign(txf, c.GetFromName(), txb, true); err != nil {
		return nil, err
	}

	txBytes, err := c.TxConfig.TxEncoder()(txb.GetTx())
	if err != nil {
		return nil, err
	}

	return c.BroadcastTx(txBytes)
}
