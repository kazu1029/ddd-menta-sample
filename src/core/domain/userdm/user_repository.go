package userdm

import "github.com/kazu1029/ddd-menta-sample/src/core/domain/tagdm"

type UserRepository interface {
	FindByID(UserID) (*User, error)
	FindWorkExperienceByWorkExperienceID(UserID, WorkExperienceID) (*UserWorkExperience, error)
	FindSkillBySkillID(UserID, tagdm.TagID) (*UserSkill, error)
	Create(*User) (*User, error)
	Update(*User) (*User, error)
}
