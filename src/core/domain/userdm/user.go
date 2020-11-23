package userdm

import (
	"unicode/utf8"

	"github.com/kazu1029/ddd-menta-sample/src/core/domain/vo"
	"golang.org/x/xerrors"
)

type User struct {
	id               UserID
	userName         string
	email            vo.Email
	password         vo.Password
	selfIntroduction string
	skills           []*UserSkill
	workExperiences  []*UserWorkExperience
}

const (
	userNameMaxLength         = 255
	selfIntroductionMaxLength = 2000
	skillsMinLength           = 0
)

func NewUser(userID UserID, userName string, email vo.Email, password vo.Password, selfIntroduction string, skills []*UserSkill, workExperiences []*UserWorkExperience) (*User, error) {
	if err := userNameValidation(userName); err != nil {
		return nil, err
	}

	if err := selfIntroductionValidation(selfIntroduction); err != nil {
		return nil, err
	}

	if err := skillsValidation(skills); err != nil {
		return nil, err
	}

	return &User{
		id:              userID,
		userName:        userName,
		email:           email,
		password:        password,
		skills:          skills,
		workExperiences: workExperiences,
	}, nil
}

func userNameValidation(userName string) error {
	if userName == "" {
		return xerrors.New("user_name must be set")
	}

	if utf8.RuneCountInString(userName) > userNameMaxLength {
		return xerrors.Errorf("user_name must be less than %d, %s", userNameMaxLength, userName)
	}

	return nil
}

func selfIntroductionValidation(selfIntroduction string) error {
	// Check if selfIntroduction is inputted
	if selfIntroduction != "" && len(selfIntroduction) > selfIntroductionMaxLength {
		return xerrors.Errorf("self_introduction must be less than %d, %s", selfIntroductionMaxLength, selfIntroduction)
	}
	return nil
}

func skillsValidation(skills []*UserSkill) error {
	if len(skills) == skillsMinLength {
		return xerrors.New("skillIDs must have more than one")
	}
	return nil
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

func (u *User) Skills() []*UserSkill {
	return u.skills
}

func (u *User) WorkExperiences() []*UserWorkExperience {
	return u.workExperiences
}

func (u *User) Equals(u2 *User) bool {
	return u.id.Equals(u2.id)
}

func (u *User) ChangeEmail(email vo.Email) {
	u.email = email
}

func (u *User) ChangePassword(password vo.Password) {
	u.password = password
}

func (u *User) ChangeUserName(userName string) error {
	if err := userNameValidation(userName); err != nil {
		return err
	}
	u.userName = userName
	return nil
}

func (u *User) ChangeSelfIntroduction(selfIntroduction string) error {
	if err := selfIntroductionValidation(selfIntroduction); err != nil {
		return err
	}
	u.selfIntroduction = selfIntroduction
	return nil
}

func (u *User) ChangeSkills(skills []*UserSkill) error {
	if err := skillsValidation(skills); err != nil {
		return err
	}
	u.skills = skills
	return nil
}

func (u *User) ChangeWorkExperiences(workExperiences []*UserWorkExperience) {
	u.workExperiences = workExperiences
}
