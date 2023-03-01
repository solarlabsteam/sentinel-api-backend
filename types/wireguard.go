package types

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/curve25519"
)

const (
	KeyLength = 32
)

type (
	Key [KeyLength]byte
)

func NewPresharedKey() (*Key, error) {
	var key Key

	_, err := rand.Read(key[:])
	if err != nil {
		return nil, err
	}

	return &key, nil
}

func NewPrivateKey() (*Key, error) {
	key, err := NewPresharedKey()
	if err != nil {
		return nil, err
	}

	key[0] &= 248
	key[31] = (key[31] & 127) | 64
	return key, nil
}

func (k *Key) String() string {
	return base64.StdEncoding.EncodeToString(k[:])
}

func (k *Key) Public() *Key {
	var p [KeyLength]byte
	curve25519.ScalarBaseMult(&p, (*[KeyLength]byte)(k))
	return (*Key)(&p)
}
