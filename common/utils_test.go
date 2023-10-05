package common

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestGetBytesOflen(t *testing.T) {
	bytes := GetBytesOflenWithZero(32)
	if len(bytes) != 32 {
		t.Fatalf("GetBytesOflen err %v", len(bytes))
	}
	fmt.Println("GetBytesOflen", bytes)
}

func TestGetBytesOflenWithRandomVal(t *testing.T) {
	bytes := GetBytesOflenWithRandomVal(32)
	if len(bytes) != 32 {
		t.Fatalf("GetBytesOflenWithRandomVal err %v", len(bytes))
	}
	strOfBase64 := base64.StdEncoding.EncodeToString(bytes)
	strOfHex := hex.EncodeToString(bytes)
	fmt.Println("GetBytesOflenWithRandomVal", bytes, strOfBase64, strOfHex)
}

func TestGetBytesOflenWithRandomValInBatch(t *testing.T) {
	for i := 0; i < 10; i++ {
		bytes := GetBytesOflenWithRandomVal(32)
		if len(bytes) != 32 {
			t.Fatalf("GetBytesOflenWithRandomVal err %v", len(bytes))
		}
		strOfHex := hex.EncodeToString(bytes)
		fmt.Println(i, strOfHex)
	}
}
