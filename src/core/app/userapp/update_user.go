package userapp

import (
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/tagdm"
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/userdm"
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/vo"
)

type UpdateUserApp struct {
	userRepo userdm.UserRepository
	tagRepo  tagdm.TagRepository
}

func NewUpdateUserApp(userRepo userdm.UserRepository, tagRepo tagdm.TagRepository) *UpdateUserApp {
	return &UpdateUserApp{
		userRepo: userRepo,
		tagRepo:  tagRepo,
	}
}

type UpdateUserRequest struct {
	ID               string
	UserName         string
	Email            string
	Password         string
	SelfIntroduction string
	SkillIDs         []UpdateUserSkillRequest
	WorkExperiences  []UpdateUserWorkExperienceRequest
}

type UpdateUserSkillRequest struct{}

type UpdateUserWorkExperienceRequest struct{}

type UpdateUserResponse struct {
	ID               string
	UserName         string
	Email            string
	SelfIntroduction string
	SkillIDs         []string
}

func (app *UpdateUserApp) Exec(req *UpdateUserRequest) (*UpdateUserResponse, error) {
	userID, err := vo.NewUserIDWithStr(req.ID)
	if err != nil {
		return nil, err
	}
	user, err := app.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}
	email, err := vo.NewEmail(req.Email)
	if err != nil {
		return nil, err
	}
	password, err := vo.NewPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// TODO: Update skills

	// TODO: Update work experiences

	// TODO: Change each field

	return nil, nil
}
