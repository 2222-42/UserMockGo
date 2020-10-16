package table

type OneTimeAccessInfo struct {
	OneTimeAccessCode string
	UserId            int64
	ExpiresAt         int64
	RetryCount        int64
}
