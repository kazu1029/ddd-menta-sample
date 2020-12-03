package mentorrecruitmentdm

type MentorRecruitmentRepository interface {
	Create(*MentorRecruitment) (*MentorRecruitment, error)
	FindAll() ([]*MentorRecruitment, error)
}
