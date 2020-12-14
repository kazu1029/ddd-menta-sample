package mentorplanapp

import (
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/categorydm"
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/tagdm"
)

type ListMentorPlanApp struct {
	mpService    MentorPlanQueryService
	categoryRepo categorydm.CategoryRepository
	tagRepo      tagdm.TagRepository
}

type ListMentorPlanItemRequest struct {
	Page  uint
	Limit uint
}

func NewListMentorPlanApp(mpService MentorPlanQueryService, categoryRepo categorydm.CategoryRepository, tagRepo tagdm.TagRepository) *ListMentorPlanApp {
	return &ListMentorPlanApp{
		mpService:    mpService,
		categoryRepo: categoryRepo,
		tagRepo:      tagRepo,
	}
}

func (app *ListMentorPlanApp) Exec(req *ListMentorPlanItemRequest) ([]*ListMentorPlanItem, error) {
	mentorPlans, err := app.mpService.FindAll(req.Page, req.Limit)
	if err != nil {
		return nil, err
	}
	masterCategories, err := app.categoryRepo.FindAll()
	if err != nil {
		return nil, err
	}
	masterCategoriesMap := makeMasterCategoryMap(masterCategories)
	masterTags, err := app.tagRepo.FindAll()
	if err != nil {
		return nil, err
	}
	masterTagMap := makeMasterTagMap(masterTags)
	mpList := make([]*ListMentorPlanItem, len(mentorPlans))
	for i, v := range mentorPlans {
		categories := make([]ListMentorPlanCategoryItem, len(v.CategoryIDs()))
		for j, category := range v.CategoryIDs() {
			categories[j] = ListMentorPlanCategoryItem{
				ID:   category.Value(),
				Name: masterCategoriesMap[category.Value()],
			}
		}
		skills := make([]ListMentorPlanSkillItem, len(v.SkillIDs()))
		for j, tag := range v.SkillIDs() {
			skills[j] = ListMentorPlanSkillItem{
				ID:   tag.Value(),
				Name: masterTagMap[tag.Value()],
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

func makeMasterCategoryMap(categories []*categorydm.Category) map[string]string {
	cmap := make(map[string]string, len(categories))
	for _, v := range categories {
		cmap[v.ID().Value()] = v.Name()
	}
	return cmap
}

func makeMasterTagMap(tags []*tagdm.Tag) map[string]string {
	tmap := make(map[string]string, len(tags))
	for _, v := range tags {
		tmap[v.ID().Value()] = v.Name()
	}
	return tmap
}
