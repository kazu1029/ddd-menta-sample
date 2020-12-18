package repoimpl

import "github.com/kazu1029/ddd-menta-sample/src/core/domain/planapplicationdm"

type PlanApplicationRepoImpl struct{}

func NewPlanApplicationRepoImpl() *PlanApplicationRepoImpl {
	return &PlanApplicationRepoImpl{}
}

var (
	planApplications []*planapplicationdm.PlanApplication = []*planapplicationdm.PlanApplication{}
)

func (repo *PlanApplicationRepoImpl) Create(planApplication *planapplicationdm.PlanApplication) (*planapplicationdm.PlanApplication, error) {
	planApplications = append(planApplications, planApplication)
	lastInsertedPlanApplication := planApplications[len(planApplications)-1]
	return lastInsertedPlanApplication, nil
}
