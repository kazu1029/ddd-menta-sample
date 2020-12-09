package categorydm

type CategoryDomainService struct {
	categoryRepo CategoryRepository
}

func NewCategoryDomainService(categoryRepo CategoryRepository) *CategoryDomainService {
	return &CategoryDomainService{
		categoryRepo: categoryRepo,
	}
}

func (service *CategoryDomainService) Exists(categoryID CategoryID) bool {
	category, err := service.categoryRepo.FindByID(categoryID)
	return !(err != nil || category == nil)
}

func (service *CategoryDomainService) ExistsWithIDs(categoryIDs []CategoryID) bool {
	for _, t := range categoryIDs {
		if ok := service.Exists(t); !ok {
			return false
		}
	}
	return true
}
