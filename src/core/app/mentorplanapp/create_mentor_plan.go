package mentorplanapp

import (
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/categorydm"
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/mentordm"
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/plandm"
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/tagdm"
	"golang.org/x/xerrors"
)

type CreateMentorPlanApp struct {
	planRepo     plandm.PlanRepository
	categoryRepo categorydm.CategoryRepository
	tagRepo      tagdm.TagRepository
}

func NewCreateMentorPlanApp(planRepo plandm.PlanRepository, categoryRepo categorydm.CategoryRepository, tagRepo tagdm.TagRepository) *CreateMentorPlanApp {
	return &CreateMentorPlanApp{
		planRepo:     planRepo,
		categoryRepo: categoryRepo,
		tagRepo:      tagRepo,
	}
}

type CreateMentorPlanRequest struct {
	MentorID       string
	Title          string
	Status         int
	Description    string
	IsSubscription bool
	Price          uint
	CategoryIDs    []string
	SkillIDs       []string
}

type CreateMentorPlanResponse struct {
	ID string
}

func (app *CreateMentorPlanApp) Exec(req *CreateMentorPlanRequest) (*CreateMentorPlanResponse, error) {
	mentorID, err := mentordm.NewMentorIDWithStr(req.MentorID)
	if err != nil {
		return nil, err
	}

	status, err := plandm.NewStatus(req.Status)
	if err != nil {
		return nil, err
	}

	price, err := plandm.NewPrice(req.Price)
	if err != nil {
		return nil, err
	}

	categoryIDs, err := categorydm.NewCategoryIDs(req.CategoryIDs)
	if err != nil {
		return nil, err
	}
	categoryDomainService := categorydm.NewCategoryDomainService(app.categoryRepo)
	if ok := categoryDomainService.ExistsWithIDs(categoryIDs); !ok {
		return nil, xerrors.Errorf("invalid category IDs, %v", categoryIDs)
	}

	tagIDs, err := tagdm.NewTagIDs(req.SkillIDs)
	if err != nil {
		return nil, err
	}
	tagDomainService := tagdm.NewTagDomainService(app.tagRepo)
	if ok := tagDomainService.ExistsWithIDs(tagIDs); !ok {
		return nil, xerrors.Errorf("invalid tag IDs, %v", tagIDs)
	}

	planID := plandm.NewPlanID()

	mentorPlan, err := plandm.NewPlan(planID, mentorID, req.Title, req.Description, status, req.IsSubscription, price, categoryIDs, tagIDs)
	if err != nil {
		return nil, err
	}

	insertedMentorPlan, err := app.planRepo.Create(mentorPlan)
	if err != nil {
		return nil, err
	}

	return &CreateMentorPlanResponse{ID: insertedMentorPlan.ID().Value()}, nil
}
