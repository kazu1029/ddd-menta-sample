package userdm

import (
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/vo"
	"golang.org/x/xerrors"
)

type User struct {
	id       UserID
	userName string
	email    vo.Email
	password vo.Password
}

const (
	userNameMaxLength = 255
)

func NewUser(userID UserID, userName string, email vo.Email, password vo.Password) (*User, error) {
	if userName == "" {
		return nil, xerrors.New("user_name must be set")
	}

	if len(userName) > 255 {
		return nil, xerrors.Errorf("user_name must be less than %d, %s", userNameMaxLength, userName)
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
