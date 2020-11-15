package userdm

import (
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/vo"
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
