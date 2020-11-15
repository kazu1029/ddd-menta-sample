package repoimpl

import (
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/tagdm"
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/vo"
)

type TagRepoImpl struct{}

func NewTagRepoImpl() *TagRepoImpl {
	return &TagRepoImpl{}
}

var (
	tags = []tagdm.Tag{}
)

func (repo *TagRepoImpl) FindByID(tagID vo.TagID) (*tagdm.Tag, error) {
	return nil, nil
}
