package menteedm

import (
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type MenteeID string

func NewMenteeID() MenteeID {
	return MenteeID(uuid.New().String())
}

func NewMenteeIDWithStr(id string) (MenteeID, error) {
	if id == "" {
		return "", xerrors.New("user id must be not empty")
	}
	return MenteeID(id), nil
}

func (m MenteeID) Value() string {
	return string(m)
}

func (mID MenteeID) Equals(mID2 MenteeID) bool {
	return mID == mID2
}
