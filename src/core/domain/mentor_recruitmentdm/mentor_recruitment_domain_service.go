package mentor_recruitmentdm

type MentorRecruitmentDomainService struct {
	mrRepo MentorRecruitmentRepository
}

func NewMentorRecruitmentDomainService(mrRepo MentorRecruitmentRepository) *MentorRecruitmentDomainService {
	return &MentorRecruitmentDomainService{
		mrRepo: mrRepo,
	}
}

func (s *MentorRecruitmentDomainService) FindAll() ([]*MentorRecruitment, error) {
	return s.mrRepo.FindAll()
}
