package mentorrecruitmentapp

import "github.com/kazu1029/ddd-menta-sample/src/core/domain/mentorrecruitmentdm"

type ListMentorRecruitmentItem struct {
	ID             string
	MenteeID       string
	Title          string
	Fee            int
	IsSubscription bool
	Description    string
	Status         string
}

type MentorRecruitmentQueryService interface {
	FindAll() ([]*mentorrecruitmentdm.MentorRecruitment, error)
}
