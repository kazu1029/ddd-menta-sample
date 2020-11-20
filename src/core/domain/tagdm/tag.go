package tagdm

import (
	"unicode/utf8"

	"golang.org/x/xerrors"
)

type Tag struct {
	id   TagID
	name string
}

const (
	nameMaxLength = 20
)

func NewTag(tagID TagID, name string) (*Tag, error) {
	if name == "" {
		return nil, xerrors.New("name must be set")
	}

	if utf8.RuneCountInString(name) > nameMaxLength {
		return nil, xerrors.Errorf("name must be less than %d, %s", nameMaxLength, name)
	}

	return &Tag{
		id:   tagID,
		name: name,
	}, nil
}

func (s *Tag) ID() TagID {
	return s.id
}

func (s *Tag) Name() string {
	return s.name
}
