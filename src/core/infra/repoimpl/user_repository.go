package repoimpl

import (
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/tagdm"
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/userdm"
)

type UserRepoImpl struct{}

func NewUserRepoImpl() *UserRepoImpl {
	return &UserRepoImpl{}
}

var (
	users []*userdm.User = []*userdm.User{}
)

func (repo *UserRepoImpl) FindByID(userID userdm.UserID) (*userdm.User, error) {
	return nil, nil
}

func (repo *UserRepoImpl) FindWorkExperienceByWorkExperienceID(workExperienceID userdm.WorkExperienceID) (*userdm.UserWorkExperience, error) {
	return nil, nil
}

func (repo *UserRepoImpl) FindSkillBySkillID(userID userdm.UserID, skillID tagdm.TagID) (*userdm.UserSkill, error) {
	return nil, nil
}

func (repo *UserRepoImpl) Create(user userdm.User) (*userdm.User, error) {
	// This is sample implementation
	u, err := userdm.NewUser(user.ID(), user.UserName(), user.Email(), user.Password(), user.SelfIntroduction(), user.Skills(), user.WorkExperiences())
	if err != nil {
		return nil, err
	}
	users = append(users, u)

	lastInsertedUser := users[len(users)-1]
	return lastInsertedUser, nil
}

func (repo *UserRepoImpl) Update(user userdm.User) (*userdm.User, error) {
	// TODO: Get target user from users

	// TODO: Update user

	// TODO: Update user skills

	// TODO: Update user work experiences
	return &user, nil
}
