package chain

import (
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/liufeihe/address-in-chain/common"
)

type EthscanRes struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

type EthAddrMgr struct {
	KeyPair *common.Secp256k1KeyPair
}

func NewEthAddrMgr() *EthAddrMgr {
	return &EthAddrMgr{}
}

func (e *EthAddrMgr) SetKeyPairWithRandom() {
	priKey := common.GetBytesOflenWithRandomVal(32)
	secp := common.GetSecp256k1KeyPairByPriKeyByte(priKey)
	e.KeyPair = secp
}

func (e *EthAddrMgr) GetAddress() (string, error) {
	if e.KeyPair == nil {
		return "", errors.New("no key pair")
	}
	ethAddress := crypto.PubkeyToAddress(*(*ecdsa.PublicKey)(e.KeyPair.PubKey.ToECDSA()))
	return ethAddress.Hex(), nil
}

func (e *EthAddrMgr) GetBalanceFromEtherScan(addr string, apiKey string) (string, error) {
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
	ethScanRes := EthscanRes{}
	err = json.Unmarshal(body, &ethScanRes)
	if err != nil {
		return "", err
	}
	if ethScanRes.Message != "OK" {
		return "", errors.New(ethScanRes.Message)
	}
	return ethScanRes.Result, nil
}
