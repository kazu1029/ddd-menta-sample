package queryserviceimpl

import "github.com/kazu1029/ddd-menta-sample/src/core/domain/mentorrecruitmentdm"

type MentorRecruitmentQueryServiceImpl struct{}

func NewMentorRecruitmentQueryService() *MentorRecruitmentQueryServiceImpl {
	return &MentorRecruitmentQueryServiceImpl{}
}

var (
	mentorRecruitments []*mentorrecruitmentdm.MentorRecruitment = []*mentorrecruitmentdm.MentorRecruitment{}
)

func (repo *MentorRecruitmentQueryServiceImpl) FindAll() ([]*mentorrecruitmentdm.MentorRecruitment, error) {
	return mentorRecruitments, nil
}
