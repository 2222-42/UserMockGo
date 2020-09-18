package randomintgenerator

import (
	"math/rand"
	"time"
)

type UserIdGeneratorMock struct {
}

func (generator UserIdGeneratorMock) Generate() int64 {
	rand.Seed(time.Now().Unix())
	return rand.Int63n(10000)
}
