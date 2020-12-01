package mentordm

type MentorDomainService struct {
	mentorRepo MentorRepository
}

func NewMentorDomainService(mentorRepo MentorRepository) *MentorDomainService {
	return &MentorDomainService{
		mentorRepo: mentorRepo,
	}
}

func (service *MentorDomainService) Exists(mentorID MentorID) bool {
	mentor, err := service.mentorRepo.FindByID(mentorID)
	return !(err != nil || mentor == nil)
}
