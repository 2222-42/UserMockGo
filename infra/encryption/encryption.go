package encryption

import (
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
