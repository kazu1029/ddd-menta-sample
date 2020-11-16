package userdm

import (
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type WorkExperienceID string

func NewWorkExperienceID() WorkExperienceID {
	return WorkExperienceID(uuid.New().String())
}

func NewWorkExperienceIDWithStr(id string) (WorkExperienceID, error) {
	if id == "" {
		return "", xerrors.New("work experience id must be not empty")
	}
	return WorkExperienceID(id), nil
}

func (we WorkExperienceID) Value() string {
	return string(we)
}

func (weID WorkExperienceID) Equals(we2ID WorkExperienceID) bool {
	return weID == we2ID
}
