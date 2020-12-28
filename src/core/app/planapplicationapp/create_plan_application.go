package planapplicationapp

import (
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/menteedm"
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/planapplicationdm"
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/plandm"
	"golang.org/x/xerrors"
)

type CreatePlanApplicationApp struct {
	pApplicationRepo planapplicationdm.PlanApplicationRepository
	planRepo         plandm.PlanRepository
	menteeRepo       menteedm.MenteeRepository
}

func NewCreawtePlanApplicationApp(pApplicationRepo planapplicationdm.PlanApplicationRepository, planRepo plandm.PlanRepository, menteeRepo menteedm.MenteeRepository) *CreatePlanApplicationApp {
	return &CreatePlanApplicationApp{
		pApplicationRepo: pApplicationRepo,
		planRepo:         planRepo,
		menteeRepo:       menteeRepo,
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
	planDomainService := plandm.NewPlanDomainService(app.planRepo)
	if !planDomainService.Exists(planID) {
		return nil, xerrors.Errorf("plan id is invalid: %s", planID)
	}

	menteeID, err := menteedm.NewMenteeIDWithStr(req.MenteeID)
	if err != nil {
		return nil, err
	}

	menteeDomainService := menteedm.NewMenteeDomainService(app.menteeRepo)
	if !menteeDomainService.Exists(menteeID) {
		return nil, xerrors.Errorf("mentee ID is invalid: %s", menteeID)
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
