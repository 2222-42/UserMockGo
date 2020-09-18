package infrainterface

type IUserIdGenerator interface {
	Generate() int64
}
