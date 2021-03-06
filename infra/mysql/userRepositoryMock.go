package mysql

import (
	"UserMockGo/domain/infrainterface"
	"UserMockGo/domain/model"
	"UserMockGo/domain/model/userModel"
	"UserMockGo/infra/myBcryption"
	"UserMockGo/infra/table"
	"UserMockGo/lib/valueObjects/userValues"
	"fmt"
	"strconv"
	"time"
)

type UserRepositoryMock struct {
	Users       *[]table.User
	Activations *[]table.Activation
	Passwords   *[]table.Password
}

func NewUserRepositoryMock() infrainterface.IUserRepository {
	users := []table.User{}
	users = append(users, table.User{
		ID:        1,
		Email:     "test1@test.com",
		IsActive:  false,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	})
	users = append(users, table.User{
		ID:        3,
		Email:     "test3@test.com",
		IsActive:  true,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	})

	activations := []table.Activation{}
	activations = append(activations, table.Activation{
		ID:                       1,
		ActivationToken:          "aaa",
		ActivationTokenExpiresAt: 2145884400,
	})
	passwords := []table.Password{}
	hashedPass, _ := myBcryption.HashPassString(userValues.PassString("test123456"))
	passwords = append(passwords, table.Password{
		ID:       3,
		Password: hashedPass,
	})

	return UserRepositoryMock{
		Users:       &users,
		Activations: &activations,
		Passwords:   &passwords,
	}
}

func (repo UserRepositoryMock) CreateUserTransactional(user userModel.User, pass userModel.Password, activation userModel.Activation) error {
	u, err := table.MapFromUserModel(user)
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", u)

	p, err := table.MapFromUserPasswordModel(pass)
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", p)

	a, err := table.MapFromUserActivationModel(activation)
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", a)

	*repo.Users = append(*repo.Users, u)
	*repo.Activations = append(*repo.Activations, a)
	*repo.Passwords = append(*repo.Passwords, p)

	return nil
}

//　Userを更新して、それのActivationを消すのをTransactionalにやる
func (repo UserRepositoryMock) ActivateUserTransactional(user userModel.User, activation userModel.Activation) error {
	users := []table.User{}
	for _, u := range *repo.Users {
		if u.ID != user.ID.ConvertUserIdToInt64() {
			users = append(users, u)
		} else {
			user.IsActive = true
			updateUser, _ := table.MapFromUserModel(user)
			users = append(users, updateUser)
		}
	}
	*repo.Users = users

	activations := []table.Activation{}
	for _, a := range *repo.Activations {
		if a.ID != activation.ID.ConvertUserIdToInt64() {
			activations = append(activations, a)
		}
	}
	*repo.Activations = activations

	return nil
}

func (repo UserRepositoryMock) FindByEmail(email userValues.Email) (userModel.User, error) {
	switch email {
	case "test2@test.com":
		return userModel.User{
			ID:        2,
			Email:     "test2@test.com",
			IsActive:  false,
			CreatedAt: time.Now().Unix() - 60*30,
			UpdatedAt: time.Now().Unix() - 60*30,
		}, nil
	case "test3@test.com":
		return userModel.User{
			ID:        3,
			Email:     "test3@test.com",
			IsActive:  true,
			CreatedAt: time.Now().Unix() - 60*30,
			UpdatedAt: time.Now().Unix() - 60*30,
		}, nil
	default:
		for _, u := range *repo.Users {
			if u.Email == string(email) {
				return u.MapToUserModel()
			}
		}
		return userModel.User{}, userModel.UserNotFound(string(email))
	}
}

func (repo UserRepositoryMock) FindById(id model.UserID) (userModel.User, error) {
	switch id {
	case 2:
		return userModel.User{
			ID:        2,
			Email:     "test2@test.com",
			IsActive:  false,
			CreatedAt: time.Now().Unix() - 60*30,
			UpdatedAt: time.Now().Unix() - 60*30,
		}, nil
	case 3:
		return userModel.User{
			ID:        3,
			Email:     "test3@test.com",
			IsActive:  true,
			CreatedAt: time.Now().Unix() - 60*30,
			UpdatedAt: time.Now().Unix() - 60*30,
		}, nil
	default:
		for _, u := range *repo.Users {
			if u.ID == id.ConvertUserIdToInt64() {
				return u.MapToUserModel()
			}
		}
		return userModel.User{}, userModel.UserNotFound(string(id))
	}
}

func (repo UserRepositoryMock) FindByUserIdAndToken(userId model.UserID, token string) (userModel.Activation, error) {
	if token != "" {
		switch userId {
		case 1:
			return userModel.Activation{
				ID:                       1,
				ActivationToken:          "aaa",
				ActivationTokenExpiresAt: 2145884400,
			}, nil
		case 2:
			return userModel.Activation{
				ID:                       2,
				ActivationToken:          "bbb",
				ActivationTokenExpiresAt: 0,
			}, nil
		default:
			for _, a := range *repo.Activations {
				if a.ID == userId.ConvertUserIdToInt64() && a.ActivationToken == token {
					return a.MapToActivationModel()
				}
			}
			return userModel.Activation{}, userModel.ActivationNotFound(strconv.Itoa(int(userId)))
		}
	}
	return userModel.Activation{}, userModel.ActivationNotFound(strconv.Itoa(int(userId)))
}

// 既存のactivationを消して作るのをTransactionalに実施する
func (repo UserRepositoryMock) ReissueOfActivationTransactional(activation userModel.Activation) error {
	activations := []table.Activation{}
	newActivation, err := table.MapFromUserActivationModel(activation)
	if err != nil {
		return err
	}

	for _, a := range *repo.Activations {
		if a.ID != activation.ID.ConvertUserIdToInt64() {
			activations = append(activations, a)
		} else {
			activations = append(activations, newActivation)
		}
	}
	*repo.Activations = activations

	return nil
}

func (repo UserRepositoryMock) GetHashedPassword(id model.UserID) (string, error) {
	for _, p := range *repo.Passwords {
		if p.ID == id.ConvertUserIdToInt64() {
			return p.MapToHashedString(), nil
		}
	}
	return "", userModel.UserPassNotFound("")
}
