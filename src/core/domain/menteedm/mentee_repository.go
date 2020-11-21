package menteedm

type MenteeRepository interface {
	FindByID(MenteeID) (*Mentee, error)
	Create(*Mentee) (*Mentee, error)
	Update(*Mentee) (*Mentee, error)
}
