package mysql

import (
	"UserMockGo/domain/infrainterface"
	"UserMockGo/domain/model"
	"UserMockGo/domain/model/errors"
	"UserMockGo/infra/table"
	"math/rand"
	"net/http"
	"time"
)

type OneTimeAccessInfoRepositoryMock struct {
	OneTimeInfos map[string]table.OneTimeAccessInfo
}

func NewOneTimeAccessInfoRepositoryMock() infrainterface.IOneTimeAccessInfoRepository {
	return OneTimeAccessInfoRepositoryMock{OneTimeInfos: map[string]table.OneTimeAccessInfo{}}
}

const rs2Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func (mock OneTimeAccessInfoRepositoryMock) CreateOneTimeAccessInfo(userId model.UserID) string {

	b := make([]byte, 10)
	for i := range b {
		b[i] = rs2Letters[rand.Intn(len(rs2Letters))]
	}

	info := table.OneTimeAccessInfo{
		OneTimeAccessCode: string(b),
		UserId:            userId,
		ExpiresAt:         time.Now().Add(time.Minute * 10).Unix(),
		RetryCount:        0,
	}

	mock.OneTimeInfos[string(b)] = info
	return string(b)
}

func (mock OneTimeAccessInfoRepositoryMock) GetUserIdByOneTimeCode(code string) (model.UserID, error) {
	info, ok := mock.OneTimeInfos[code]
	if !ok {
		return 0, errors.MyError{
			StatusCode: http.StatusNotFound,
			Message:    "one time info not found",
			ErrorType:  "one_time_info_not_found",
		}
	}

	if info.RetryCount > 3 {
		delete(mock.OneTimeInfos, code)
		return 0, errors.MyError{
			StatusCode: http.StatusInternalServerError,
			Message:    "one time info exceeds retry_count",
			ErrorType:  "one_time_info_not_valid",
		}
	}

	return info.UserId, nil
}

func (mock OneTimeAccessInfoRepositoryMock) RemoveAccessInfo(code string) {
	_, ok := mock.OneTimeInfos[code]
	if !ok {
		return
	}

	delete(mock.OneTimeInfos, code)
	return
}

func (mock OneTimeAccessInfoRepositoryMock) IncrementRetryCount(code string) {

	info, ok := mock.OneTimeInfos[code]
	if !ok {
		return
	}

	info.RetryCount += 1
	mock.OneTimeInfos[code] = info
	return
}
