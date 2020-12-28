package planapplicationdm

import (
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/menteedm"
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/plandm"
	"golang.org/x/xerrors"
)

type PlanApplication struct {
	id       PlanApplicationID
	planID   plandm.PlanID
	menteeID menteedm.MenteeID
	status   Status
}

func NewPlanApplication(id PlanApplicationID, planID plandm.PlanID, menteeID menteedm.MenteeID, status Status) (*PlanApplication, error) {
	return &PlanApplication{
		id:       id,
		planID:   planID,
		menteeID: menteeID,
		status:   status,
	}, nil
}

func (pa *PlanApplication) ID() PlanApplicationID {
	return pa.id
}

func (pa *PlanApplication) PlanID() plandm.PlanID {
	return pa.planID
}

func (pa *PlanApplication) MenteeID() menteedm.MenteeID {
	return pa.menteeID
}

func (pa *PlanApplication) Status() Status {
	return pa.status
}

func (pa *PlanApplication) Contract() error {
	if pa.status == Contracted {
		return xerrors.New("plan application status is already contracted")
	}

	if pa.status == Rejected {
		return xerrors.New("plan application status is already rejected")
	}

	contractedStatus, err := NewStatusForUpdate(Contracted.Value())
	if err != nil {
		return err
	}

	pa.status = contractedStatus

	return nil
}

func (pa *PlanApplication) Reject() error {
	if pa.status == Rejected {
		return xerrors.New("plan application is already rejected")
	}

	if pa.status == Contracted {
		return xerrors.New("plan application is already contracted")
	}

	rejectedStatus, err := NewStatusForUpdate(Rejected.Value())
	if err != nil {
		return err
	}

	pa.status = rejectedStatus

	return nil
}
