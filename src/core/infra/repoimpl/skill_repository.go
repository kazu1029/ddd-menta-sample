package repoimpl

import "github.com/kazu1029/ddd-menta-sample/src/core/domain/skilldm"

type SkillRepoImpl struct{}

func NewSkillRepoImpl() *SkillRepoImpl {
	return &SkillRepoImpl{}
}

func (repo *SkillRepoImpl) FindByID(skillID skilldm.SkillID) (*skilldm.Skill, error) {
	return nil, nil
}
