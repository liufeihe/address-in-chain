package chain

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/liufeihe/address-in-chain/common"
)

type EthMgr struct {
	KeyPair *common.Secp256k1KeyPair
}

func NewEthMgr() *EthMgr {
	return &EthMgr{}
}

func (e *EthMgr) SetKeyPairWithRandom() {
	priKey := common.GetBytesOflenWithRandomVal(32)
	secp := common.GetSecp256k1KeyPairByPriKeyByte(priKey)
	e.KeyPair = secp
}

func (e *EthMgr) GetAddress() (string, error) {
	if e.KeyPair == nil {
		return "", errors.New("no key pair")
	}
	ethAddress := crypto.PubkeyToAddress(*(*ecdsa.PublicKey)(e.KeyPair.PubKey.ToECDSA()))
	return ethAddress.Hex(), nil
}

func (e *EthMgr) GetBalanceFromEtherScan(addr string, apiKey string) (string, error) {
	url := fmt.Sprintf("https://api.etherscan.io/api?module=account&action=balance&address=%s&tag=latest&apikey=%s", addr, apiKey)

	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	fmt.Println("Balance (Wei):", string(body))
	return "", nil
}
