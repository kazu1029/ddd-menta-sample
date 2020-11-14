package userdm

import (
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/skilldm"
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/vo"
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

func (service *UserDomainService) IsExists(userID vo.UserID) bool {
	user, err := service.userRepo.FindByID(userID)
	return !(err != nil || user == nil)
}

// CanRegisterWithRequiredItems method checks if user can register with required params.
// required items:
// 		at least one SkillID
func (service *UserDomainService) CanRegisterWithRequiredItems(skillIDs []skilldm.SkillID) error {
	if len(skillIDs) == 0 {
		return xerrors.New("skill must be set one or more")
	}

	return nil
}
