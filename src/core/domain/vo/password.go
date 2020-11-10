package vo

import (
	"reflect"
	"regexp"

	"golang.org/x/xerrors"
)

type Password string

var (
	passwordFormat = `^(?=.*?[a-z])(?=.*?\d)[a-z\d]+$`
	passwordRegExp = regexp.MustCompile(passwordFormat)
)

const (
	passwordMinLength = 12
	passwordMaxLength = 255
)

func NewPassword(password string) (Password, error) {
	if len(password) == 0 {
		return Password(""), xerrors.New("password must set")
	}

	if len(password) < passwordMinLength {
		return Password(""), xerrors.Errorf("password must be more than %d", passwordMinLength)
	}

	if len(password) > passwordMaxLength {
		return Password(""), xerrors.Errorf("password must be less than %d", passwordMaxLength)
	}

	if ok := passwordRegExp.MatchString(password); !ok {
		return Password(""), xerrors.New("invalid password format.")
	}

	return Password(password), nil
}

func (p Password) Value() string {
	return string(p)
}

func (p Password) Equals(p2 Password) bool {
	return reflect.DeepEqual(p, p2)
}
