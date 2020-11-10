package userdm

import "golang.org/x/xerrors"

type UserID uint32

func NewUserID(userID uint32) (UserID, error) {
	if userID < 0 {
		return UserID(0), xerrors.New("user_id must be more than 1")
	}

	return UserID(0), nil
}

func (u UserID) Value() uint32 {
	return uint32(u)
}

func (uID UserID) Equals(uID2 UserID) bool {
	return uID == uID2
}
