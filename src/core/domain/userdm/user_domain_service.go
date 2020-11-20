package userdm

type UserDomainService struct {
	userRepo UserRepository
}

func NewUserDomainService(userRepo UserRepository) *UserDomainService {
	return &UserDomainService{
		userRepo: userRepo,
	}
}

func (service *UserDomainService) Exists(userID UserID) bool {
	user, err := service.userRepo.FindByID(userID)
	return !(err != nil || user == nil)
}
