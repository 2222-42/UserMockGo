package encryption

import (
	"UserMockGo/domain/model/valueObjects"
	"golang.org/x/crypto/bcrypt"
)

func PassEncryption(pass valueObjects.PassString) (string, error) {
	hp, err := bcrypt.GenerateFromPassword([]byte(string(pass)), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hp), nil
}

func ComparePass(hp []byte, pass valueObjects.PassString) bool {
	return bcrypt.CompareHashAndPassword(hp, []byte(string(pass))) == nil
}
