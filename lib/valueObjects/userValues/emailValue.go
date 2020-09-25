package userValues

import "regexp"

type Email string

const emailRex string = `[\w\-._]+@[\w\-._]+\.[A-Za-z]+`

func (email Email) IsValidForm() bool {
	result, err := regexp.MatchString(emailRex, string(email))
	if err != nil {
		return false
	}
	return result
}
