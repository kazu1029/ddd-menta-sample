package userapp

import (
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/tagdm"
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/userdm"
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/vo"
)

type CreateUserApp struct {
	userRepo userdm.UserRepository
	tagRepo  tagdm.TagRepository
}

func NewCreateUserApp(userRepo userdm.UserRepository, tagRepo tagdm.TagRepository) *CreateUserApp {
	return &CreateUserApp{
		userRepo: userRepo,
		tagRepo:  tagRepo,
	}
}

type CreateUserRequest struct {
	UserName         string
	Email            string
	Password         string
	SelfIntroduction string
	Skills           []CreateUserSkillRequset
}

type CreateUserSkillRequset struct {
	ID                string
	YearsOfExperience int
}

type CreateUserResponse struct {
	ID string
}

func (app *CreateUserApp) Exec(req *CreateUserRequest) (*CreateUserResponse, error) {
	email, err := vo.NewEmail(req.Email)
	if err != nil {
		return nil, err
	}
	password, err := vo.NewPassword(req.Password)
	if err != nil {
		return nil, err
	}
	userID := vo.NewUserID()

	var userSkillIDs []string
	for _, skill := range req.Skills {
		userSkillIDs = append(userSkillIDs, skill.ID)
	}
	skillIDs, err := vo.NewTagIDs(userSkillIDs)
	if err != nil {
		return nil, err
	}
	tagDomainService := tagdm.NewTagDomainService(app.tagRepo)
	if ok := tagDomainService.ExistsWithIDs(skillIDs); !ok {
		return nil, err
	}

	user, err := userdm.NewUser(userID, req.UserName, email, password, req.SelfIntroduction, skillIDs)
	if err != nil {
		return nil, err
	}

	insertedUser, err := app.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return &CreateUserResponse{ID: insertedUser.ID().Value()}, nil
}
