package userdm

import "github.com/kazu1029/ddd-menta-sample/src/core/domain/tagdm"

type UserSkill struct {
	id                tagdm.TagID
	ownerID           UserID
	yearsOfExperience YearsOfExperience
}

func NewUserSkill(skillID tagdm.TagID, ownerID UserID, yearsOfExperience YearsOfExperience) (*UserSkill, error) {
	return &UserSkill{
		id:                skillID,
		ownerID:           ownerID,
		yearsOfExperience: yearsOfExperience,
	}, nil
}

func (s *UserSkill) ID() tagdm.TagID {
	return s.id
}

func (s *UserSkill) OwnerID() UserID {
	return s.ownerID
}

func (s *UserSkill) YearsOfExperience() YearsOfExperience {
	return s.yearsOfExperience
}

func (s *UserSkill) ChangeYearsOfExperience(yearsOfExperience YearsOfExperience) {
	s.yearsOfExperience = yearsOfExperience
}
