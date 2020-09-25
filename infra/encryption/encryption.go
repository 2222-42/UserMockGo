package encryption

import (
	"UserMockGo/domain/infrainterface"
	"UserMockGo/lib/valueObjects/userValues"
	"golang.org/x/crypto/bcrypt"
)

func PassEncryption(pass userValues.PassString) (string, error) {
	hp, err := bcrypt.GenerateFromPassword([]byte(string(pass)), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hp), nil
}

func ComparePass(hp []byte, pass userValues.PassString) bool {
	return bcrypt.CompareHashAndPassword(hp, []byte(string(pass))) == nil
}

type LoginInfraMock struct {
}

func NewLoginInfraMock() infrainterface.ILogin {
	return LoginInfraMock{}
}

func (login LoginInfraMock) CheckPassAndHash(hp string, passString userValues.PassString) bool {
	return ComparePass([]byte(hp), passString)
}
