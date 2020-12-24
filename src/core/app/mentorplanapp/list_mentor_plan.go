package mentorplanapp

import (
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/categorydm"
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/tagdm"
)

var (
	pageLimit uint = 20
)

type ListMentorPlanApp struct {
	mpService    MentorPlanQueryService
	categoryRepo categorydm.CategoryRepository
	tagRepo      tagdm.TagRepository
}

type ListMentorPlanItemRequest struct {
	Status int
	Page   uint
	// Limit  uint
}

func NewListMentorPlanApp(mpService MentorPlanQueryService, categoryRepo categorydm.CategoryRepository, tagRepo tagdm.TagRepository) *ListMentorPlanApp {
	return &ListMentorPlanApp{
		mpService:    mpService,
		categoryRepo: categoryRepo,
		tagRepo:      tagRepo,
	}
}

func (app *ListMentorPlanApp) Exec(req *ListMentorPlanItemRequest) ([]*ListMentorPlanItem, error) {
	mentorPlans, err := app.mpService.FindAllByStatus(req.Status, req.Page, pageLimit)
	if err != nil {
		return nil, err
	}
	mpList := make([]*ListMentorPlanItem, len(mentorPlans))
	for i, v := range mentorPlans {
		categories := make([]ListMentorPlanCategoryItem, len(v.CategoryIDs()))
		for j, category := range v.Categories {
			categories[j] = ListMentorPlanCategoryItem{
				ID:   category.ID().Value(),
				Name: category.Name(),
			}
		}
		skills := make([]ListMentorPlanSkillItem, len(v.SkillIDs()))
		for j, tag := range v.Skills {
			skills[j] = ListMentorPlanSkillItem{
				ID:   tag.ID().Value(),
				Name: tag.Name(),
			}
		}

		mpList[i] = &ListMentorPlanItem{
			ID:             v.ID().Value(),
			MentorID:       v.MentorID().Value(),
			Title:          v.Title(),
			Description:    v.Description(),
			IsSubscription: v.IsSubScription(),
			Status:         v.Status().String(),
			Price:          v.Price().Value(),
			Categories:     categories,
			Skills:         skills,
		}
	}

	return mpList, nil
}
