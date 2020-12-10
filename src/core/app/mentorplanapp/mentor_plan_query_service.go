package mentorplanapp

import "github.com/kazu1029/ddd-menta-sample/src/core/domain/plandm"

type ListMentorPlanItem struct {
	ID             string
	MentorID       string
	Title          string
	Description    string
	IsSubscription bool
	Status         string
	Price          uint
	Categories     []ListMentorPlanCategoryItem
	Skills         []ListMentorPlanSkillItem
}

type ListMentorPlanCategoryItem struct {
	ID   string
	Name string
}

type ListMentorPlanSkillItem struct {
	ID   string
	Name string
}

type MentorPlanQueryService interface {
	FindAll(uint, uint) ([]*plandm.Plan, error)
}
