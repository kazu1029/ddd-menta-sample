package plandm

import (
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type PlanID string

func NewPlanID() PlanID {
	return PlanID(uuid.New().String())
}

func NewPlanIDWithStr(id string) (PlanID, error) {
	if id == "" {
		return "", xerrors.New("plan id must be not empty")
	}
	return PlanID(id), nil
}

func (pID PlanID) Value() string {
	return string(pID)
}

func (pID PlanID) Equals(pID2 PlanID) bool {
	return pID == pID2
}
