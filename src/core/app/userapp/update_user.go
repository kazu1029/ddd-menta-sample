package userapp

import (
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/userdm"
)

type UpdateUserApp struct {
	userRepo userdm.UserRepository
}

func NewUpdateUserApp(userRepo userdm.UserRepository) *UpdateUserApp {
	return &UpdateUserApp{
		userRepo: userRepo,
	}
}

type UpdateUserRequest struct {
	ID               uint32
	UserName         string
	Email            string
	Password         string
	SelfIntroduction string
	SkillIDs         []uint32
}

type UpdateUserResponse struct {
	ID               uint32
	UserName         string
	Email            string
	SelfIntroduction string
	SkillIDs         []uint32
}

func (app *UpdateUserApp) Exec(req *UpdateUserRequest) (*UpdateUserResponse, error) {
	// userID, err := vo.NewUserID(req.ID)
	// if err != nil {
	// 	return nil, err
	// }
	// user, err := app.userRepo.FindByID(userID)
	// if err != nil {
	// 	return nil, err
	// }
	// email, err := vo.NewEmail(req.Email)
	// if err != nil {
	// 	return nil, err
	// }
	// password, err := vo.NewPassword(req.Password)
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}
