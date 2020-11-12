package skilldm

import "golang.org/x/xerrors"

type Skill struct {
	id                SkillID
	name              string
	yearsOfExperience YearsOfExperience
}

const (
	nameMaxLength = 20
)

func NewSkill(skillID SkillID, name string, yearsOfExperience YearsOfExperience) (*Skill, error) {
	if name == "" {
		return nil, xerrors.New("name must be set")
	}

	if len(name) > nameMaxLength {
		return nil, xerrors.Errorf("name must be less than %d, %s", nameMaxLength, name)
	}

	return &Skill{
		id:                skillID,
		name:              name,
		yearsOfExperience: yearsOfExperience,
	}, nil
}

func (s *Skill) ID() SkillID {
	return s.id
}

func (s *Skill) Name() string {
	return s.name
}

func (s *Skill) YearsOfExperience() YearsOfExperience {
	return s.yearsOfExperience
}
