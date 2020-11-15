package tagdm

import "github.com/kazu1029/ddd-menta-sample/src/core/domain/vo"

type TagDomainService struct {
	tagRepo TagRepository
}

func NewTagDomainService(tagRepo TagRepository) *TagDomainService {
	return &TagDomainService{
		tagRepo: tagRepo,
	}
}

func (service *TagDomainService) Exists(tagID vo.TagID) bool {
	tag, err := service.tagRepo.FindByID(tagID)
	return !(err != nil || tag == nil)
}

func (service *TagDomainService) ExistsWithIDs(tagIDs []vo.TagID) bool {
	for _, t := range tagIDs {
		if ok := service.Exists(t); !ok {
			return false
		}
	}
	return true
}
