package queryserviceimpl

import "github.com/kazu1029/ddd-menta-sample/src/core/domain/plandm"

type MentorPlanQueryServiceImpl struct{}

func NewMentorPlanQueryService() *MentorPlanQueryServiceImpl {
	return &MentorPlanQueryServiceImpl{}
}

var (
	mentorPlans []*plandm.Plan = []*plandm.Plan{}
)

func (repo *MentorPlanQueryServiceImpl) FindAll(page, limit uint) ([]*plandm.Plan, error) {
	return mentorPlans, nil
}
