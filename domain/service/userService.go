package service

import (
	user2 "UserMockGo/domain/model/user"
	"math/rand"
	"time"
)

//Passwordはこの時点ではいらないかも？
func CreateUser(email string, password string, passwordConfirmation string) error {
	// TODO: idの生成
	rand.Seed(time.Now().Unix())
	id := rand.Int63n(10000)
	// TODO: timerを導入する
	now := time.Now().Unix()

	//user :=
	user2.NewUser(id, email, password, passwordConfirmation, now)

	//infrainterface使ってInsertする

	return nil
}
