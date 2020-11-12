package repoimpl

import "github.com/kazu1029/ddd-menta-sample/src/core/domain/userdm"

type UserRepoImpl struct{}

func NewUserRepoImpl() *UserRepoImpl {
	return &UserRepoImpl{}
}

func (repo *UserRepoImpl) FindByID(userID userdm.UserID) (*userdm.User, error) {
	return nil, nil
}

func (repo *UserRepoImpl) Create(user userdm.User) error {
	return nil
}
