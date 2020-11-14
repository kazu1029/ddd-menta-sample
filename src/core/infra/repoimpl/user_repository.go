package repoimpl

import (
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/userdm"
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/vo"
)

type UserRepoImpl struct{}

func NewUserRepoImpl() *UserRepoImpl {
	return &UserRepoImpl{}
}

var (
	users []*userdm.User = []*userdm.User{}
)

func (repo *UserRepoImpl) FindByID(userID vo.UserID) (*userdm.User, error) {
	return nil, nil
}

func (repo *UserRepoImpl) GetLastID() (vo.UserID, error) {
	usersCount := len(users)
	if usersCount == 0 {
		return 1, nil
	}
	lastUser := users[usersCount-1]
	return lastUser.ID(), nil
}

func (repo *UserRepoImpl) Create(user userdm.User) (*userdm.User, error) {
	return nil, nil
}
