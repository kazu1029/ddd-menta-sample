package menteedm

import (
	"unicode/utf8"

	"github.com/kazu1029/ddd-menta-sample/src/core/domain/vo"
	"golang.org/x/xerrors"
)

type Mentee struct {
	id               MenteeID
	menteeName       string
	email            vo.Email
	password         vo.Password
	selfIntroduction string
	skills           []*MenteeSkill
	workExperiences  []*MenteeWorkExperience
}

const (
	userNameMaxLength         = 255
	selfIntroductionMaxLength = 2000
	skillsMinLength           = 0
)

func NewMentee(menteeID MenteeID, menteeName string, email vo.Email, password vo.Password, selfIntroduction string, skills []*MenteeSkill, workExperiences []*MenteeWorkExperience) (*Mentee, error) {
	if err := menteeNameValidation(menteeName); err != nil {
		return nil, err
	}

	if err := selfIntroductionValidation(selfIntroduction); err != nil {
		return nil, err
	}

	if err := skillsValidation(skills); err != nil {
		return nil, err
	}

	return &Mentee{
		id:              menteeID,
		menteeName:      menteeName,
		email:           email,
		password:        password,
		skills:          skills,
		workExperiences: workExperiences,
	}, nil
}

func menteeNameValidation(menteeName string) error {
	if menteeName == "" {
		return xerrors.New("user_name must be set")
	}

	if utf8.RuneCountInString(menteeName) > userNameMaxLength {
		return xerrors.Errorf("mentee name must be less than %d, %s", userNameMaxLength, menteeName)
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

func skillsValidation(skills []*MenteeSkill) error {
	if len(skills) == skillsMinLength {
		return xerrors.New("skillIDs must have more than one")
	}
	return nil
}

func (m *Mentee) ID() MenteeID {
	return m.id
}

func (m *Mentee) MenteeName() string {
	return m.menteeName
}

func (m *Mentee) Email() vo.Email {
	return m.email
}

func (m *Mentee) Password() vo.Password {
	return m.password
}

func (m *Mentee) SelfIntroduction() string {
	return m.selfIntroduction
}

func (m *Mentee) Skills() []*MenteeSkill {
	return m.skills
}

func (m *Mentee) WorkExperiences() []*MenteeWorkExperience {
	return m.workExperiences
}

func (m *Mentee) Equals(m2 *Mentee) bool {
	return m.id.Equals(m2.id)
}

func (m *Mentee) ChangeEmail(email vo.Email) error {
	m.email = email
	return nil
}

func (m *Mentee) ChangePassword(password vo.Password) error {
	m.password = password
	return nil
}

func (m *Mentee) ChangeMenteeName(menteeName string) error {
	if err := menteeNameValidation(menteeName); err != nil {
		return err
	}
	m.menteeName = menteeName
	return nil
}

func (m *Mentee) ChangeSelfIntroduction(selfIntroduction string) error {
	if err := selfIntroductionValidation(selfIntroduction); err != nil {
		return err
	}
	m.selfIntroduction = selfIntroduction
	return nil
}

func (m *Mentee) ChangeSkills(skills []*MenteeSkill) error {
	if err := skillsValidation(skills); err != nil {
		return err
	}
	m.skills = skills
	return nil
}

func (m *Mentee) ChangeWorkExperiences(workExperiences []*MenteeWorkExperience) error {
	m.workExperiences = workExperiences
	return nil
}
