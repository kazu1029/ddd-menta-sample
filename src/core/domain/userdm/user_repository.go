package userdm

import "github.com/kazu1029/ddd-menta-sample/src/core/domain/vo"

type UserRepository interface {
	FindByID(vo.UserID) (*User, error)
	Create(*User) (*User, error)
}
