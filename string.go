package _mock

import (
	"math/rand"
	"time"
)

func Username() string {
	
	return ""
}

func String() string {
	return StringFixed(24)
}

func StringZero() string {
	return ""
}

func StringFixed(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < l; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

func StringRange(min, max int) string {
	return ""
}

func Password() string {
	return ""
}
