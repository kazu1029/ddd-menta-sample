package userdm

import (
	"unicode/utf8"

	"github.com/kazu1029/ddd-menta-sample/src/core/domain/tagdm"
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/vo"
	"golang.org/x/xerrors"
)

type User struct {
	id                UserID
	userName          string
	email             vo.Email
	password          vo.Password
	selfIntroduction  string
	skillIDs          []tagdm.TagID
	workExperienceIDs []WorkExperienceID
}

const (
	userNameMaxLength         = 255
	selfIntroductionMaxLength = 2000
	skillIDsMinLength         = 0
)

func NewUser(userID UserID, userName string, email vo.Email, password vo.Password, selfIntroduction string, skillIDs []tagdm.TagID, workExperienceIDs []WorkExperienceID) (*User, error) {
	if userName == "" {
		return nil, xerrors.New("user_name must be set")
	}

	if utf8.RuneCountInString(userName) > userNameMaxLength {
		return nil, xerrors.Errorf("user_name must be less than %d, %s", userNameMaxLength, userName)
	}

	// Check if selfIntroduction is inputted
	if selfIntroduction != "" && len(selfIntroduction) > selfIntroductionMaxLength {
		return nil, xerrors.Errorf("self_introduction must be less than %d, %s", selfIntroductionMaxLength, selfIntroduction)
	}

	if len(skillIDs) == skillIDsMinLength {
		return nil, xerrors.New("skillIDs must have more than one")
	}

	return &User{
		id:                userID,
		userName:          userName,
		email:             email,
		password:          password,
		skillIDs:          skillIDs,
		workExperienceIDs: workExperienceIDs,
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

func (u *User) SelfIntroduction() string {
	return u.selfIntroduction
}

func (u *User) SkillIDs() []tagdm.TagID {
	return u.skillIDs
}

func (u *User) WorkExperienceIDs() []WorkExperienceID {
	return u.workExperienceIDs
}

func (u *User) Equals(u2 *User) bool {
	return u.id.Equals(u2.id)
}
