package userdm

import "github.com/kazu1029/ddd-menta-sample/src/core/domain/vo"

type UserSkill struct {
	id                vo.TagID
	ownerID           vo.UserID
	yearsOfExperience YearsOfExperience
}

func NewUserSkill(skillID vo.TagID, ownerID vo.UserID, yearsOfExperience YearsOfExperience) (*UserSkill, error) {
	return &UserSkill{
		id:                skillID,
		ownerID:           ownerID,
		yearsOfExperience: yearsOfExperience,
	}, nil
}

func (s *UserSkill) ID() vo.TagID {
	return s.id
}

func (s *UserSkill) OwnerID() vo.UserID {
	return s.ownerID
}

func (s *UserSkill) YearsOfExperience() YearsOfExperience {
	return s.yearsOfExperience
}
