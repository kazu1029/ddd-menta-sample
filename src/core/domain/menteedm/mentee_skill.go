package menteedm

import "github.com/kazu1029/ddd-menta-sample/src/core/domain/tagdm"

type MenteeSkill struct {
	id                tagdm.TagID
	ownerID           MenteeID
	yearsOfExperience YearsOfExperience
}

func NewMenteeSkill(skillID tagdm.TagID, yearsOfExperience YearsOfExperience) (*MenteeSkill, error) {
	return &MenteeSkill{
		id:                skillID,
		yearsOfExperience: yearsOfExperience,
	}, nil
}

func (s *MenteeSkill) ID() tagdm.TagID {
	return s.id
}

func (s *MenteeSkill) YearsOfExperience() YearsOfExperience {
	return s.yearsOfExperience
}
