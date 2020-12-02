package mentor_recruitmentdm

import (
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type BudgetID string

func NewBudgetID() BudgetID {
	return BudgetID(uuid.New().String())
}

func NewBudgetIDWithStr(id string) (BudgetID, error) {
	if id == "" {
		return "", xerrors.New("user id must be not empty")
	}
	return BudgetID(id), nil
}

func (b BudgetID) Value() string {
	return string(b)
}

func (bID BudgetID) Equals(bID2 BudgetID) bool {
	return bID == bID2
}
