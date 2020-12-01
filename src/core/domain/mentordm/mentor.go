package mentordm

import (
	"unicode/utf8"

	"github.com/kazu1029/ddd-menta-sample/src/core/domain/vo"
	"golang.org/x/xerrors"
)

type Mentor struct {
	id               MentorID
	mentorName       string
	email            vo.Email
	password         vo.Password
	selfIntroduction string
	skills           []*MentorSkill
	workExperiences  []*MentorWorkExperience
}

const (
	mentorNameMaxLength       = 255
	selfIntroductionMaxLength = 2000
	skillsMinLength           = 0
)

func NewMentor(mentorID MentorID, mentorName string, email vo.Email, password vo.Password, selfIntroduction string, skills []*MentorSkill, workExperiences []*MentorWorkExperience) (*Mentor, error) {
	if err := mentorNameValidation(mentorName); err != nil {
		return nil, err
	}

	if err := selfIntroductionValidation(selfIntroduction); err != nil {
		return nil, err
	}

	if err := skillsValidation(skills); err != nil {
		return nil, err
	}

	return &Mentor{
		id:              mentorID,
		mentorName:      mentorName,
		email:           email,
		password:        password,
		skills:          skills,
		workExperiences: workExperiences,
	}, nil
}

func mentorNameValidation(mentorName string) error {
	if mentorName == "" {
		return xerrors.New("mentor_name must be set")
	}

	if utf8.RuneCountInString(mentorName) > mentorNameMaxLength {
		return xerrors.Errorf("mentor_name must be less than %d, %s", mentorNameMaxLength, mentorName)
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

func skillsValidation(skills []*MentorSkill) error {
	if len(skills) == skillsMinLength {
		return xerrors.New("skillIDs must have more than one")
	}
	return nil
}

func (m *Mentor) ID() MentorID {
	return m.id
}

func (m *Mentor) MentorName() string {
	return m.mentorName
}

func (m *Mentor) Email() vo.Email {
	return m.email
}

func (m *Mentor) Password() vo.Password {
	return m.password
}

func (m *Mentor) SelfIntroduction() string {
	return m.selfIntroduction
}

func (m *Mentor) Skills() []*MentorSkill {
	return m.skills
}

func (m *Mentor) WorkExperiences() []*MentorWorkExperience {
	return m.workExperiences
}

func (m *Mentor) Equals(m2 *Mentor) bool {
	return m.id.Equals(m2.id)
}

func (m *Mentor) ChangeEmail(email vo.Email) {
	m.email = email
}

func (m *Mentor) ChangePassword(password vo.Password) {
	m.password = password
}

func (m *Mentor) ChangeMentorName(mentorName string) error {
	if err := mentorNameValidation(mentorName); err != nil {
		return err
	}
	m.mentorName = mentorName
	return nil
}

func (u *Mentor) ChangeSelfIntroduction(selfIntroduction string) error {
	if err := selfIntroductionValidation(selfIntroduction); err != nil {
		return err
	}
	u.selfIntroduction = selfIntroduction
	return nil
}

func (u *Mentor) ChangeSkills(skills []*MentorSkill) error {
	if err := skillsValidation(skills); err != nil {
		return err
	}
	u.skills = skills
	return nil
}

func (u *Mentor) ChangeWorkExperiences(workExperiences []*MentorWorkExperience) {
	u.workExperiences = workExperiences
}
