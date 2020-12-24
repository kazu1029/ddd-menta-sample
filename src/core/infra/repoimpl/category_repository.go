package repoimpl

import (
	"strconv"

	"github.com/kazu1029/ddd-menta-sample/src/core/domain/categorydm"
)

type CategoryRepoImpl struct{}

func NewCategoryRepoImpl() *CategoryRepoImpl {
	return &CategoryRepoImpl{}
}

var (
	categories map[string]*categorydm.Category
)

func init() {
	category1, _ := categorydm.NewCategory(categorydm.CategoryID("id0"), "Programming")
	category2, _ := categorydm.NewCategory(categorydm.CategoryID("id1"), "Marketing")
	category3, _ := categorydm.NewCategory(categorydm.CategoryID("id2"), "Web Design")
	category4, _ := categorydm.NewCategory(categorydm.CategoryID("id3"), "Writing")
	category5, _ := categorydm.NewCategory(categorydm.CategoryID("id4"), "Language")
	categories["id0"] = category1
	categories["id1"] = category2
	categories["id2"] = category3
	categories["id3"] = category4
	categories["id4"] = category5
}

func (repo *CategoryRepoImpl) FindByID(categoryID categorydm.CategoryID) (*categorydm.Category, error) {
	return categories[categoryID.Value()], nil
}

func (repo *CategoryRepoImpl) FindByIDs(categoryIDs []categorydm.CategoryID) ([]*categorydm.Category, error) {
	var fetchedCategorys []*categorydm.Category
	for _, categoryID := range categoryIDs {
		if _, ok := categories[categoryID.Value()]; ok {
			fetchedCategorys = append(fetchedCategorys, categories[categoryID.Value()])
		}
	}
	return fetchedCategorys, nil
}

func (repo *CategoryRepoImpl) FindAll() ([]*categorydm.Category, error) {
	categorySlice := make([]*categorydm.Category, len(categories))
	for i := 0; i < len(categories); i++ {
		categorySlice[i] = categories["id"+strconv.Itoa(i)]
	}
	return categorySlice, nil
}
