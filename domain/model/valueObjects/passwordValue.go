package valueObjects

import "unicode/utf8"

const minimumLength int = 12

type PassString string

func (pass PassString) IsValidLength() bool {
	return utf8.RuneCountInString(string(pass)) >= minimumLength
}
