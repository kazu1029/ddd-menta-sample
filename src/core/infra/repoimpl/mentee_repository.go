package repoimpl

import (
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/menteedm"
)

type MenteeRepoImpl struct{}

func NewMenteeRepoImpl() *MenteeRepoImpl {
	return &MenteeRepoImpl{}
}

var (
	mentees []*menteedm.Mentee = []*menteedm.Mentee{}
)

func (repo *MenteeRepoImpl) FindByID(menteeID menteedm.MenteeID) (*menteedm.Mentee, error) {
	return nil, nil
}

func (repo *MenteeRepoImpl) Create(mentee menteedm.Mentee) (*menteedm.Mentee, error) {
	// This is sample implementation
	u, err := menteedm.NewMentee(mentee.ID(), mentee.MenteeName(), mentee.Email(), mentee.Password(), mentee.SelfIntroduction(), mentee.Skills(), mentee.WorkExperiences())
	if err != nil {
		return nil, err
	}
	mentees = append(mentees, u)

	lastInsertedMentee := mentees[len(mentees)-1]
	return lastInsertedMentee, nil
}

func (repo *MenteeRepoImpl) Update(mentee menteedm.Mentee) (*menteedm.Mentee, error) {
	// TODO: Get target mentee from mentees

	// TODO: Update mentee

	// TODO: Update mentee skills

	// TODO: Update mentee work experiences
	return &mentee, nil
}
