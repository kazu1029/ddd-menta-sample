package plandm

type PlanDomainService struct {
	planRepo PlanRepository
}

func NewPlanDomainService(planRepo PlanRepository) *PlanDomainService {
	return &PlanDomainService{
		planRepo: planRepo,
	}
}

func (service *PlanDomainService) Exists(planID PlanID) bool {
	plan, err := service.planRepo.FindByID(planID)
	return !(err != nil || plan == nil)
}
