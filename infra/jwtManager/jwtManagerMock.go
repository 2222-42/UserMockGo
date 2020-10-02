package jwtManager

import (
	"UserMockGo/domain/infrainterface"
	"UserMockGo/domain/model/user"
	"github.com/dgrijalva/jwt-go"
	"os"
	"strconv"
	"time"
)

type TokenManagerMock struct {
}

func NewTokenManagerMock() infrainterface.ITokenManager {
	return TokenManagerMock{}
}

func (manager TokenManagerMock) GenerateToken(u user.User) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	// claimsのセット
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = strconv.Itoa(int(u.ID))
	claims["email"] = string(u.Email)
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// 電子署名
	tokenString, err := token.SignedString([]byte(os.Getenv("SIGNINGKEY")))

	// JWTを返却
	return tokenString, err
}

//func (manager TokenManagerMock) ParseToken(token string) bool {
//
//}
