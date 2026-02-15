package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alpabets = "qwertyuiopasdfghjklzxcvbnm"

func init() {
	source := rand.NewSource(time.Now().UnixMilli())
	rand.New(source)
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(size int) string {
	var sb strings.Builder
	k := len(alpabets)
	for i := 0; i < size; i++ {
		sb.WriteByte(alpabets[rand.Intn(k)])
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(7)
}

func RamdonBalnce() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"INR", "USD", "EUR", "RUB", "CY", "JY", "KY"}
	return currencies[rand.Intn(len(currencies))]
}

func RandomEmail() string {
	return fmt.Sprintf("%s@mail.com", RandomString(5))
}
