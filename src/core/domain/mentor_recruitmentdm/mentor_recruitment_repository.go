package mentor_recruitmentdm

type MentorRecruitmentRepository interface {
	Create(*MentorRecruitment) (*MentorRecruitment, error)
}
