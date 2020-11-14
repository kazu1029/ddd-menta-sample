package userapp

import (
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/skilldm"
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/userdm"
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/vo"
)

type CreateUserApp struct {
	userRepo userdm.UserRepository
}

func NewCreateUserApp(userRepo userdm.UserRepository) *CreateUserApp {
	return &CreateUserApp{
		userRepo: userRepo,
	}
}

type CreateUserRequest struct {
	UserName         string
	Email            string
	Password         string
	SelfIntroduction string
	SkillIDs         []uint32
}

type CreateUserResponse struct {
	ID uint32
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
	lastUserID, err := app.userRepo.GetLastID()
	if err != nil {
		return nil, err
	}
	userID := lastUserID + 1

	user, err := userdm.NewUser(userID, req.UserName, email, password, req.SelfIntroduction)
	if err != nil {
		return nil, err
	}
	skillIDs, err := skilldm.NewSkillIDs(req.SkillIDs)
	if err != nil {
		return nil, err
	}

	userDomainService := userdm.NewUserDomainService(app.userRepo)
	if err := userDomainService.CanRegisterWithRequiredItems(skillIDs); err != nil {
		return nil, err
	}

	// TODO: Check if SkillIDs are valid

	createdUser, err := app.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return &CreateUserResponse{ID: createdUser.ID().Value()}, nil
}
