package mentor_recruitmentapp

import "github.com/kazu1029/ddd-menta-sample/src/core/domain/mentor_recruitmentdm"

type ListMentorRecruitmentApp struct {
	mrService mentor_recruitmentdm.MentorRecruitmentDomainService
}

func NewListMentorRecruitmentApp(mrService mentor_recruitmentdm.MentorRecruitmentDomainService) *ListMentorRecruitmentApp {
	return &ListMentorRecruitmentApp{
		mrService: mrService,
	}
}

type ListMentorRecruitmentItem struct {
	ID             string
	MenteeID       string
	Title          string
	Fee            int
	IsSubscription bool
	Description    string
	Status         string
}

func (app *ListMentorRecruitmentApp) Exec() ([]*ListMentorRecruitmentItem, error) {
	mentorRecruitments, err := app.mrService.FindAll()
	if err != nil {
		return nil, err
	}
	mrList := make([]*ListMentorRecruitmentItem, len(mentorRecruitments))
	for i, v := range mentorRecruitments {
		mrList[i] = &ListMentorRecruitmentItem{
			ID:             v.ID().Value(),
			MenteeID:       v.MenteeID().Value(),
			Title:          v.Title(),
			Fee:            v.Budget().Fee(),
			IsSubscription: v.Budget().IsSubscription(),
			Description:    v.Description(),
			Status:         v.Status().String(),
		}
	}

	return mrList, err
}
