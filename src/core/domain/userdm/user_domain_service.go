package userdm

import (
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/skilldm"
	"golang.org/x/xerrors"
)

type UserDomainService struct {
	userRepo UserRepository
}

func NewUserDomainService(userRepo UserRepository) *UserDomainService {
	return &UserDomainService{
		userRepo: userRepo,
	}
}

func (service *UserDomainService) IsExists(userID UserID) bool {
	user, err := service.userRepo.FindByID(userID)
	return !(err != nil || user == nil)
}

// CanRegisterUser method checks if user has one or more SkillID.
func (service *UserDomainService) CanRegisterUser(skillIDs []skilldm.SkillID) error {
	if len(skillIDs) == 0 {
		return xerrors.New("skill must be set one or more")
	}

	return nil
}
