package skilldm

import "golang.org/x/xerrors"

type SkillID uint32

func NewSkillID(skillID uint32) (SkillID, error) {
	if skillID < 0 {
		return SkillID(0), xerrors.New("skill_id must be more than 1")
	}

	return SkillID(skillID), nil
}

func (s SkillID) Value() uint32 {
	return uint32(s)
}

func (sID SkillID) Equals(sID2 SkillID) bool {
	return sID == sID2
}
