package utils

import (
	cryptohd "github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	uuid2 "github.com/hashicorp/go-uuid"
)

func NewInMemoryKey(mnemonic string, coinType, account, index uint32, bip39Password string) (keyring.Keyring, keyring.Info, error) {
	uuid, err := uuid2.GenerateUUID()
	if err != nil {
		return nil, nil, err
	}

	var (
		kr   = keyring.NewInMemory()
		path = cryptohd.CreateHDPath(coinType, account, index)
	)

	key, err := kr.NewAccount(uuid, mnemonic, bip39Password, path.String(), cryptohd.Secp256k1)
	if err != nil {
		return nil, nil, err
	}

	return kr, key, nil
}
