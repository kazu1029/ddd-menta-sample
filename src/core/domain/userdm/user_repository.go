package userdm

type UserRepository interface {
	FindByID(UserID) (*User, error)
	Create(*User) error
}
