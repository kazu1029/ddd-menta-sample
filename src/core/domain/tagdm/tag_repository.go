package tagdm

import "github.com/kazu1029/ddd-menta-sample/src/core/domain/vo"

type TagRepository interface {
	FindByID(vo.TagID) (*Tag, error)
}
