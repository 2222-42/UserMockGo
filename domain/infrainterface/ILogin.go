package infrainterface

import "UserMockGo/lib/valueObjects/userValues"

type ILogin interface {
	CheckPassAndHash(hp string, passString userValues.PassString) bool
}
