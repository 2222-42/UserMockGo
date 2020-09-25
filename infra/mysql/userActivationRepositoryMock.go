package mysql

import (
	"UserMockGo/domain/infrainterface"
	"UserMockGo/domain/model"
	"UserMockGo/domain/model/user"
	"strconv"
)

type ActivationRepositoryMock struct {
}

func NewActivationRepositoryMock() infrainterface.IActivationRepository {
	return ActivationRepositoryMock{}
}

func (repo ActivationRepositoryMock) FindByUserIdAndToken(userId model.UserID, token string) (user.Activation, error) {
	if token != "" {
		switch userId {
		case 1:
			return user.Activation{
				ID:                       1,
				ActivationToken:          "aaa",
				ActivationTokenExpiresAt: 2145884400,
			}, nil
		case 2:
			return user.Activation{
				ID:                       2,
				ActivationToken:          "bbb",
				ActivationTokenExpiresAt: 0,
			}, nil
		default:
			return user.Activation{}, user.ActivationNotFound(strconv.Itoa(int(userId)))
		}
	}
	return user.Activation{}, user.ActivationNotFound(strconv.Itoa(int(userId)))

}
