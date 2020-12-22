package repoimpl

import "github.com/kazu1029/ddd-menta-sample/src/core/domain/plandm"

type PlanRepoImpl struct{}

func NewPlanRepoImpl() *PlanRepoImpl {
	return &PlanRepoImpl{}
}

var (
	plans []*plandm.Plan = []*plandm.Plan{}
)

func (repo *PlanRepoImpl) Create(plan *plandm.Plan) (*plandm.Plan, error) {
	plans = append(plans, plan)
	lastInsertedPlan := plans[len(plans)-1]
	return lastInsertedPlan, nil
}

func (repo *PlanRepoImpl) FindByID(planID plandm.PlanID) (*plandm.Plan, error) {
	return nil, nil
}
