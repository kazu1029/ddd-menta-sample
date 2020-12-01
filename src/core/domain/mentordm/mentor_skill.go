package mentordm

import "github.com/kazu1029/ddd-menta-sample/src/core/domain/tagdm"

type MentorSkill struct {
	id                tagdm.TagID
	yearsOfExperience YearsOfExperience
}

func NewMentorSkill(skillID tagdm.TagID, yearsOfExperience YearsOfExperience) (*MentorSkill, error) {
	return &MentorSkill{
		id:                skillID,
		yearsOfExperience: yearsOfExperience,
	}, nil
}

func (s *MentorSkill) ID() tagdm.TagID {
	return s.id
}

func (s *MentorSkill) YearsOfExperience() YearsOfExperience {
	return s.yearsOfExperience
}

func (s *MentorSkill) ChangeYearsOfExperience(yearsOfExperience YearsOfExperience) {
	s.yearsOfExperience = yearsOfExperience
}
