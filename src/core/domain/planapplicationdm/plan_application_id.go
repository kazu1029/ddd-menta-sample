package planapplicationdm

import (
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type PlanApplicationID string

func NewPlanApplicationID() PlanApplicationID {
	return PlanApplicationID(uuid.New().String())
}

func NewPlanApplicationIDWithStr(id string) (PlanApplicationID, error) {
	if id == "" {
		return "", xerrors.New("mentor recruitment id must be not empty")
	}
	return PlanApplicationID(id), nil
}

func (paID PlanApplicationID) Value() string {
	return string(paID)
}

func (paID PlanApplicationID) Equals(paID2 PlanApplicationID) bool {
	return paID == paID2
}
