package mentordm

import (
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type MentorID string

func NewMentorID() MentorID {
	return MentorID(uuid.New().String())
}

func NewMentorIDWithStr(id string) (MentorID, error) {
	if id == "" {
		return "", xerrors.New("mentor id must be not empty")
	}
	return MentorID(id), nil
}

func (mID MentorID) Value() string {
	return string(mID)
}

func (mID MentorID) Equals(mID2 MentorID) bool {
	return mID == mID2
}
