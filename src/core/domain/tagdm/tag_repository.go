package tagdm

type TagRepository interface {
	FindByID(TagID) (*Tag, error)
	FindByIDs([]TagID) ([]*Tag, error)
	FindAll() ([]*Tag, error)
}
