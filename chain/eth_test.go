package chain

import (
	"fmt"
	"testing"
)

func TestEthGetAddress(t *testing.T) {
	mgr := NewEthMgr()
	if mgr == nil {
		t.Fatalf("NewEthMgr err")
	}
	mgr.SetKeyPairWithRandom()
	addr, err := mgr.GetAddress()
	if err != nil {
		t.Fatalf("GetAddress err %v", err)
	}
	fmt.Println("eth address", addr)
}

func TestGetBalanceFromEtherScan(t *testing.T) {
	mgr := NewEthMgr()
	if mgr == nil {
		t.Fatalf("NewEthMgr err")
	}
	mgr.SetKeyPairWithRandom()
	addr, err := mgr.GetAddress()
	if err != nil {
		t.Fatalf("GetAddress err %v", err)
	}
	apiKey := ""
	balance, err := mgr.GetBalanceFromEtherScan(addr, apiKey)
	if err != nil {
		t.Fatalf("GetBalanceFromEtherScan err %v", err)
	}
	fmt.Println("GetBalanceFromEtherScan", balance)
}
