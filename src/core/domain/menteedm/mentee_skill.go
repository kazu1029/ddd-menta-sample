package menteedm

import "github.com/kazu1029/ddd-menta-sample/src/core/domain/tagdm"

type MenteeSkill struct {
	id                tagdm.TagID
	ownerID           MenteeID
	yearsOfExperience YearsOfExperience
}

func NewMenteeSkill(skillID tagdm.TagID, ownerID MenteeID, yearsOfExperience YearsOfExperience) (*MenteeSkill, error) {
	return &MenteeSkill{
		id:                skillID,
		ownerID:           ownerID,
		yearsOfExperience: yearsOfExperience,
	}, nil
}

func (s *MenteeSkill) ID() tagdm.TagID {
	return s.id
}

func (s *MenteeSkill) OwnerID() MenteeID {
	return s.ownerID
}

func (s *MenteeSkill) YearsOfExperience() YearsOfExperience {
	return s.yearsOfExperience
}
