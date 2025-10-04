package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomInt(min, max int64) int64 {
	return min + rng.Int63n(max-min+1)
}

func Randomstring(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i <= n; i++ {
		randomschar := alphabet[rng.Intn(k)]
		sb.WriteByte(randomschar)
	}
	return sb.String()
}

func RandomName() string {
	return Randomstring(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currency := []string{"USD", "CAD", "INR"}
	k := len(currency)
	return currency[rng.Intn(k)]
}
