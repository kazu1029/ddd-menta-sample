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

func (repo *UserRepoImpl) Create(user userdm.User) (*userdm.User, error) {
	// This is sample implementation
	u, err := userdm.NewUser(user.ID(), user.UserName(), user.Email(), user.Password(), user.SelfIntroduction(), user.SkillIDs(), user.WorkExperienceIDs())
	if err != nil {
		return nil, err
	}
	users = append(users, u)

	lastInsertedUser := users[len(users)-1]
	return lastInsertedUser, nil
}
