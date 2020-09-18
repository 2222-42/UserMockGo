package infrainterface

type IUserTokenGenerator interface {
	GenerateTokenAndExpiresAt() (string, int64)
}
