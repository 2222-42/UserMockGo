package jwtManager

import (
	"UserMockGo/domain/infrainterface"
	"UserMockGo/domain/model"
	"UserMockGo/domain/model/authorization"
	"UserMockGo/domain/model/errors"
	"UserMockGo/domain/model/user"
	"UserMockGo/lib/valueObjects/userValues"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strconv"
	"time"
)

type TokenManager struct {
}

func NewTokenManagerMock() infrainterface.ITokenManager {
	return TokenManager{}
}

func (manager TokenManager) GenerateToken(u user.User) (string, error) {

	// TODO: HS256を使わなくする
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

func (manager TokenManager) Parse(tokenString string) (authorization.Authorization, error) {
	if tokenString == "" {
		return authorization.Authorization{}, errors.MyError{
			StatusCode: http.StatusForbidden,
			Message:    "no token",
			ErrorType:  "invalid_token",
		}
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SIGNINGKEY")), nil
	})

	if err != nil {
		return authorization.Authorization{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["sub"], claims["email"])
		id, err := strconv.ParseInt(claims["sub"].(string), 10, 64)
		if err != nil {
			return authorization.Authorization{}, err
		}

		return authorization.Authorization{
			UserId: model.UserID(id),
			Email:  userValues.Email(claims["email"].(string)),
		}, nil
	} else {
		fmt.Println(err)
		return authorization.Authorization{}, err
	}
}
