package repoimpl

import "github.com/kazu1029/ddd-menta-sample/src/core/domain/mentor_recruitmentdm"

type MentorRecruitmentRepoImpl struct{}

func NewMentorRecruitmentRepoImpl() *MentorRecruitmentRepoImpl {
	return &MentorRecruitmentRepoImpl{}
}

var (
	mentorRecruitments []*mentor_recruitmentdm.MentorRecruitment = []*mentor_recruitmentdm.MentorRecruitment{}
)

func (repo *MentorRecruitmentRepoImpl) Create(mentorRecruitment *mentor_recruitmentdm.MentorRecruitment) (*mentor_recruitmentdm.MentorRecruitment, error) {
	mentorRecruitments = append(mentorRecruitments, mentorRecruitment)
	lastInsertedMentorRecruitment := mentorRecruitments[len(mentorRecruitments)-1]
	return lastInsertedMentorRecruitment, nil
}
