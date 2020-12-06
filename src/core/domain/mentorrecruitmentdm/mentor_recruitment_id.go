package mentorrecruitmentdm

import (
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type MentorRecruitmentID string

func NewMentorRecruitmentID() MentorRecruitmentID {
	return MentorRecruitmentID(uuid.New().String())
}

func NewMentorRecruitmentIDWithStr(id string) (MentorRecruitmentID, error) {
	if id == "" {
		return "", xerrors.New("mentor recruitment id must be not empty")
	}
	return MentorRecruitmentID(id), nil
}

func (mrID MentorRecruitmentID) Value() string {
	return string(mrID)
}

func (mrID MentorRecruitmentID) Equals(mrID2 MentorRecruitmentID) bool {
	return mrID == mrID2
}
