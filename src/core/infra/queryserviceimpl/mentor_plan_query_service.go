package queryserviceimpl

import (
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/categorydm"
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/plandm"
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/tagdm"
)

type MentorPlanQueryServiceImpl struct{}

func NewMentorPlanQueryService() *MentorPlanQueryServiceImpl {
	return &MentorPlanQueryServiceImpl{}
}

var (
	mentorPlanList []*ListMentorPlanItem = []*ListMentorPlanItem{}
)

type ListMentorPlanItem struct {
	*plandm.Plan
	Categories []*categorydm.Category
	Skills     []*tagdm.Tag
}

func (repo *MentorPlanQueryServiceImpl) FindAllByStatus(status int, page, limit uint) ([]*ListMentorPlanItem, error) {
	return mentorPlanList, nil
}
