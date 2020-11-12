package skilldm

type SkillRepository interface {
	FindByID(SkillID) (*Skill, error)
}
