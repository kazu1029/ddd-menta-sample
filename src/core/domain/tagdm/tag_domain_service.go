package tagdm

type TagDomainService struct {
	tagRepo TagRepository
}

func NewTagDomainService(tagRepo TagRepository) *TagDomainService {
	return &TagDomainService{
		tagRepo: tagRepo,
	}
}

func (service *TagDomainService) Exists(tagID TagID) bool {
	tag, err := service.tagRepo.FindByID(tagID)
	return !(err != nil || tag == nil)
}

func (service *TagDomainService) ExistsWithIDs(tagIDs []TagID) bool {
	for _, t := range tagIDs {
		if ok := service.Exists(t); !ok {
			return false
		}
	}
	return true
}
