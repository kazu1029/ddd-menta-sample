package repoimpl

import "github.com/kazu1029/ddd-menta-sample/src/core/domain/mentorrecruitmentdm"

type MentorRecruitmentRepoImpl struct{}

func NewMentorRecruitmentRepoImpl() *MentorRecruitmentRepoImpl {
	return &MentorRecruitmentRepoImpl{}
}

var (
	mentorRecruitments []*mentorrecruitmentdm.MentorRecruitment = []*mentorrecruitmentdm.MentorRecruitment{}
)

func (repo *MentorRecruitmentRepoImpl) Create(mentorRecruitment *mentorrecruitmentdm.MentorRecruitment) (*mentorrecruitmentdm.MentorRecruitment, error) {
	mentorRecruitments = append(mentorRecruitments, mentorRecruitment)
	lastInsertedMentorRecruitment := mentorRecruitments[len(mentorRecruitments)-1]
	return lastInsertedMentorRecruitment, nil
}

func (repo *MentorRecruitmentRepoImpl) FindAll() ([]*mentorrecruitmentdm.MentorRecruitment, error) {
	return mentorRecruitments, nil
}
