package repoimpl

import (
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/mentordm"
)

type MentorRepoImpl struct{}

func NewMentorRepoImpl() *MentorRepoImpl {
	return &MentorRepoImpl{}
}

var (
	mentors []*mentordm.Mentor = []*mentordm.Mentor{}
)

// FindByID method return Mentor along with MentorSkills and MentorWorkExperiences
func (repo *MentorRepoImpl) FindByID(mentorID mentordm.MentorID) (*mentordm.Mentor, error) {
	return nil, nil
}

// Create method saves Mentor with MentorSkills and MentorWorkExperiences
func (repo *MentorRepoImpl) Create(mentor mentordm.Mentor) (*mentordm.Mentor, error) {
	// This is sample implementation
	u, err := mentordm.NewMentor(mentor.ID(), mentor.MentorName(), mentor.Email(), mentor.Password(), mentor.SelfIntroduction(), mentor.Skills(), mentor.WorkExperiences())
	if err != nil {
		return nil, err
	}
	mentors = append(mentors, u)

	lastInsertedMentor := mentors[len(mentors)-1]
	return lastInsertedMentor, nil
}

// Update method changes Mentor along with MentorSkills and MentorWorkExperiences
func (repo *MentorRepoImpl) Update(mentor mentordm.Mentor) (*mentordm.Mentor, error) {
	// TODO: Get target mentor from mentors

	// TODO: Update mentor

	// TODO: Update mentor skills

	// TODO: Update mentor work experiences
	return &mentor, nil
}
