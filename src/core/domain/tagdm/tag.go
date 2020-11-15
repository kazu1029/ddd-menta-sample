package tagdm

import (
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/vo"
	"golang.org/x/xerrors"
)

type Tag struct {
	id   vo.TagID
	name string
}

const (
	nameMaxLength = 20
)

func NewTag(tagID vo.TagID, name string) (*Tag, error) {
	if name == "" {
		return nil, xerrors.New("name must be set")
	}

	if len(name) > nameMaxLength {
		return nil, xerrors.Errorf("name must be less than %d, %s", nameMaxLength, name)
	}

	return &Tag{
		id:   tagID,
		name: name,
	}, nil
}

func (s *Tag) ID() vo.TagID {
	return s.id
}

func (s *Tag) Name() string {
	return s.name
}
