package categorydm

type CategoryRepository interface {
	FindByID(CategoryID) (*Category, error)
	FindByIDs([]CategoryID) ([]*Category, error)
}
