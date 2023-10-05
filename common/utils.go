package common

import (
	"math/rand"
	"time"
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
