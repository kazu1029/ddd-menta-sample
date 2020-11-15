package vo

import (
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type UserID string

func NewUserID() UserID {
	return UserID(uuid.New().String())
}

func NewUserIDWithStr(id string) (UserID, error) {
	if id == "" {
		return "", xerrors.New("user id must be not empty")
	}
	return UserID(id), nil
}

func (u UserID) Value() string {
	return string(u)
}

func (uID UserID) Equals(uID2 UserID) bool {
	return uID == uID2
}
