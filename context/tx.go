package context

import (
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txsigning "github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
)

func (c Context) Tx(kr keyring.Keyring, from string, gas uint64, memo string, gasPrices sdk.DecCoins, chainID, rpcAddress, broadcastMode string, messages ...sdk.Msg) (result *sdk.TxResponse, err error) {
	c.ChainID = chainID
	c.BroadcastMode = broadcastMode

	c.Client, err = rpchttp.New(rpcAddress, "/websocket")
	if err != nil {
		return nil, err
	}

	key, err := kr.Key(from)
	if err != nil {
		return nil, err
	}

	account, err := c.QueryAccount(rpcAddress, key.GetAddress())
	if err != nil {
		return nil, err
	}

	txb := c.TxConfig.NewTxBuilder()
	if err = txb.SetMsgs(messages...); err != nil {
		return nil, err
	}

	txb.SetGasLimit(gas)
	txb.SetMemo(memo)

	if !gasPrices.IsZero() {
		var (
			gas  = sdk.NewDec(int64(gas))
			fees = make(sdk.Coins, len(gasPrices))
		)

		for i, price := range gasPrices {
			fee := price.Amount.Mul(gas)
			fees[i] = sdk.NewCoin(price.Denom, fee.Ceil().RoundInt())
		}

		txb.SetFeeAmount(fees)
	}

	txSignature := txsigning.SignatureV2{
		PubKey: &secp256k1.PubKey{
			Key: key.GetPubKey().Bytes(),
		},
		Data: &txsigning.SingleSignatureData{
			SignMode:  c.TxConfig.SignModeHandler().DefaultMode(),
			Signature: nil,
		},
		Sequence: account.GetSequence(),
	}

	if err = txb.SetSignatures(txSignature); err != nil {
		return nil, err
	}

	message, err := c.TxConfig.SignModeHandler().GetSignBytes(
		c.TxConfig.SignModeHandler().DefaultMode(),
		authsigning.SignerData{
			ChainID:       c.ChainID,
			AccountNumber: account.GetAccountNumber(),
			Sequence:      account.GetSequence(),
		},
		txb.GetTx(),
	)
	if err != nil {
		return nil, err
	}

	signature, _, err := kr.Sign(from, message)
	if err != nil {
		return nil, err
	}

	txSignature.Data = &txsigning.SingleSignatureData{
		SignMode:  c.TxConfig.SignModeHandler().DefaultMode(),
		Signature: signature,
	}

	if err = txb.SetSignatures(txSignature); err != nil {
		return nil, err
	}

	txBytes, err := c.TxConfig.TxEncoder()(txb.GetTx())
	if err != nil {
		return nil, err
	}

	return c.BroadcastTx(txBytes)
}
