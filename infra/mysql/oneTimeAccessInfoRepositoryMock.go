package mysql

import (
	"UserMockGo/domain/infrainterface"
	"UserMockGo/domain/model"
	"UserMockGo/domain/model/errors"
	"UserMockGo/infra/table"
	"net/http"
	"time"
)

type OneTimeAccessInfoRepositoryMock struct {
	OneTimeInfos map[model.UserID]table.OneTimeAccessInfo
}

func NewOneTimeAccessInfoRepositoryMock() infrainterface.IOneTimeAccessInfoRepository {
	return OneTimeAccessInfoRepositoryMock{OneTimeInfos: map[model.UserID]table.OneTimeAccessInfo{}}
}

func (mock OneTimeAccessInfoRepositoryMock) CreateOneTimeAccessInfo(userId model.UserID, code string) error {

	// TODO: ここのcode生成はinfra層の責務にする
	info := table.OneTimeAccessInfo{
		OneTimeAccessCode: code,
		UserId:            int64(userId),
		ExpiresAt:         time.Now().Add(time.Minute * 10).Unix(),
		RetryCount:        0,
	}

	mock.OneTimeInfos[userId] = info
	return nil
}

// 一時的なコードが存在するかを確認
// そのコードが
func (mock OneTimeAccessInfoRepositoryMock) CheckWithCode(userId model.UserID, code string) error {

	info, ok := mock.OneTimeInfos[userId]
	if !ok {
		return errors.MyError{
			StatusCode: http.StatusNotFound,
			Message:    "one time info not found",
			ErrorType:  "one_time_info_not_found",
		}
	}

	if info.RetryCount > 3 {
		delete(mock.OneTimeInfos, userId)
		return errors.MyError{
			StatusCode: http.StatusInternalServerError,
			Message:    "one time info exceeds retry_count",
			ErrorType:  "one_time_info_not_valid",
		}
	}

	if info.OneTimeAccessCode != code {
		info.RetryCount += 1
		return errors.MyError{
			StatusCode: http.StatusBadRequest,
			Message:    "code does not match",
			ErrorType:  "code_is_not_valid",
		}
	}

	return nil
}
