package mentorplanapp

type ListMentorPlanApp struct {
	mpService MentorPlanQueryService
}

type ListMentorPlanItemRequest struct {
	Page  uint
	Limit uint
}

func NewListMentorPlanApp(mpService MentorPlanQueryService) *ListMentorPlanApp {
	return &ListMentorPlanApp{
		mpService: mpService,
	}
}

func (app *ListMentorPlanApp) Exec(req *ListMentorPlanItemRequest) ([]*ListMentorPlanItem, error) {
	mentorPlans, err := app.mpService.FindAll(req.Page, req.Limit)
	if err != nil {
		return nil, err
	}
	mpList := make([]*ListMentorPlanItem, len(mentorPlans))
	for i, v := range mentorPlans {
		// Categories and Skills are needed to get from domain service.
		mpList[i] = &ListMentorPlanItem{
			ID:             v.ID().Value(),
			MentorID:       v.MentorID().Value(),
			Title:          v.Title(),
			Description:    v.Description(),
			IsSubscription: v.IsSubScription(),
			Status:         v.Status().String(),
		}
	}
}
