package menteedm

type MenteeDomainService struct {
	menteeRepo MenteeRepository
}

func NewMenteeDomainService(menteeRepo MenteeRepository) *MenteeDomainService {
	return &MenteeDomainService{
		menteeRepo: menteeRepo,
	}
}

func (service *MenteeDomainService) Exists(menteeID MenteeID) bool {
	mentee, err := service.menteeRepo.FindByID(menteeID)
	return !(err != nil || mentee == nil)
}
