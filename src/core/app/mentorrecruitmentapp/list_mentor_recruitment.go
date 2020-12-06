package mentorrecruitmentapp

type ListMentorRecruitmentApp struct {
	mrService MentorRecruitmentQueryService
}

func NewListMentorRecruitmentApp(mrService MentorRecruitmentQueryService) *ListMentorRecruitmentApp {
	return &ListMentorRecruitmentApp{
		mrService: mrService,
	}
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
