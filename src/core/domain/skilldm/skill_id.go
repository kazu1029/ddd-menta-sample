package skilldm

import "golang.org/x/xerrors"

type SkillID uint32

func NewSkillID(skillID uint32) (SkillID, error) {
	if skillID < 0 {
		return SkillID(0), xerrors.New("skill_id must be more than 1")
	}

	return SkillID(skillID), nil
}

func NewSkillIDs(skillIDs []uint32) ([]SkillID, error) {
	var ids []SkillID
	for _, s := range skillIDs {
		skillID, err := NewSkillID(s)
		if err != nil {
			return nil, err
		}
		ids = append(ids, skillID)
	}
	return ids, nil
}

func (s SkillID) Value() uint32 {
	return uint32(s)
}

func (sID SkillID) Equals(sID2 SkillID) bool {
	return sID == sID2
}
