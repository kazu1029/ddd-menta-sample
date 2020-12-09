package categorydm

import (
	"unicode/utf8"

	"golang.org/x/xerrors"
)

type Category struct {
	id   CategoryID
	name string
}

const (
	nameMaxLength = 20
)

func NewCategory(categoryID CategoryID, name string) (*Category, error) {
	if name == "" {
		return nil, xerrors.New("name must be set")
	}

	if utf8.RuneCountInString(name) > nameMaxLength {
		return nil, xerrors.Errorf("name must be less than %d, %s", nameMaxLength, name)
	}

	return &Category{
		id:   categoryID,
		name: name,
	}, nil
}

func (s *Category) ID() CategoryID {
	return s.id
}

func (s *Category) Name() string {
	return s.name
}
