package skilldm

type SkillDomainService struct {
	skillRepo SkillRepository
}

func NewSkillDomainService(skillRepo SkillRepository) *SkillDomainService {
	return &SkillDomainService{
		skillRepo: skillRepo,
	}
}

func (service *SkillDomainService) IsExists(skillID SkillID) bool {
	skill, err := service.skillRepo.FindByID(skillID)
	return !(err != nil || skill == nil)
}
