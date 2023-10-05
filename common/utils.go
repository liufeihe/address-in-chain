package common

import (
	"math/rand"
	"time"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
)

func GetTimestampInMs() int64 {
	return time.Now().UnixMilli()
}

func GetTimestampInNano() int64 {
	return time.Now().UnixNano()
}

func GetBytesOflenWithZero(len int64) []byte {
	return make([]byte, len)
}

func GetBytesOflenWithRandomVal(len int64) []byte {
	rand.Seed(GetTimestampInNano())
	bytes := make([]byte, len)
	for i := int64(0); i < len; i++ {
		// bytes[i] = 255
		bytes[i] = byte(rand.Intn(256))
	}
	return bytes
}

type Secp256k1KeyPair struct {
	PriKey *secp256k1.PrivateKey
	PubKey *secp256k1.PublicKey
}

func GetSecp256k1KeyPairByPriKeyByte(priKey []byte) *Secp256k1KeyPair {
	privateKey, publicKey := btcec.PrivKeyFromBytes(priKey)
	return &Secp256k1KeyPair{
		PriKey: privateKey,
		PubKey: publicKey,
	}
}
