package userdm

import (
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/vo"
	"golang.org/x/xerrors"
)

type User struct {
	id               UserID
	userName         string
	email            vo.Email
	password         vo.Password
	isMentor         bool
	selfIntroduction string
}

const (
	userNameMaxLength         = 255
	selfIntroductionMaxLength = 2000
)

func NewUser(userID UserID, userName string, email vo.Email, password vo.Password, selfIntroduction string) (*User, error) {
	if userName == "" {
		return nil, xerrors.New("user_name must be set")
	}

	if len(userName) > 255 {
		return nil, xerrors.Errorf("user_name must be less than %d, %s", userNameMaxLength, userName)
	}

	if len(selfIntroduction) > selfIntroductionMaxLength {
		return nil, xerrors.Errorf("self_introduction must be less than %d, %s", selfIntroductionMaxLength, selfIntroduction)
	}

	return &User{
		id:       userID,
		userName: userName,
		email:    email,
		password: password,
	}, nil
}

func (u *User) ID() UserID {
	return u.id
}

func (u *User) UserName() string {
	return u.userName
}

func (u *User) Email() vo.Email {
	return u.email
}

func (u *User) Password() vo.Password {
	return u.password
}

func (u *User) Equals(u2 *User) bool {
	return u.id.Equals(u2.id)
}
