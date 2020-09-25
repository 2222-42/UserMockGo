package token

import (
	"math/rand"
	"time"
)

type UserTokenGeneratorMock struct {
}

const rs2Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func (generator UserTokenGeneratorMock) GenerateTokenAndExpiresAt() (string, int64) {
	b := make([]byte, 24)
	for i := range b {
		b[i] = rs2Letters[rand.Intn(len(rs2Letters))]
	}
	return string(b), time.Now().Unix() + 60*60
}
