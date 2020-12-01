package mentordm

type MentorRepository interface {
	FindByID(MentorID) (*Mentor, error)
	Create(*Mentor) (*Mentor, error)
	Update(*Mentor) (*Mentor, error)
}
