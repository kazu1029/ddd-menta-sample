package planapplicationapp

import (
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/menteedm"
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/planapplicationdm"
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/plandm"
)

type CreatePlanApplicationApp struct {
	pApplicationRepo planapplicationdm.PlanApplicationRepository
}

func NewCreawtePlanApplicationApp(pApplicationRepo planapplicationdm.PlanApplicationRepository) *CreatePlanApplicationApp {
	return &CreatePlanApplicationApp{
		pApplicationRepo: pApplicationRepo,
	}
}

type CreatePlanApplicationRequest struct {
	PlanID   string
	MenteeID string
}

type CreatePlanApplicationResponse struct {
	ID string
}

func (app *CreatePlanApplicationApp) Exec(req *CreatePlanApplicationRequest) (*CreatePlanApplicationResponse, error) {
	planID, err := plandm.NewPlanIDWithStr(req.PlanID)
	if err != nil {
		return nil, err
	}

	menteeID, err := menteedm.NewMenteeIDWithStr(req.MenteeID)
	if err != nil {
		return nil, err
	}

	status := planapplicationdm.NewStatus()
	planApplicationID := planapplicationdm.NewPlanApplicationID()

	planApplication, err := planapplicationdm.NewPlanApplication(
		planApplicationID,
		planID,
		menteeID,
		status,
	)
	if err != nil {
		return nil, err
	}

	insertedPlanApplicaton, err := app.pApplicationRepo.Create(planApplication)
	if err != nil {
		return nil, err
	}

	return &CreatePlanApplicationResponse{ID: insertedPlanApplicaton.ID().Value()}, nil
}
