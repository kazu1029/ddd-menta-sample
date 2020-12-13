package categorydm

import (
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type CategoryID string

func NewCategoryID() CategoryID {
	return CategoryID(uuid.New().String())
}

func NewCategoryIDWithStr(id string) (CategoryID, error) {
	if id == "" {
		return "", xerrors.New("category id must be not empty")
	}

	return CategoryID(id), nil
}

func NewCategoryIDs(categoryIDs []string) ([]CategoryID, error) {
	var ids []CategoryID
	for _, s := range categoryIDs {
		categoryID, err := NewCategoryIDWithStr(s)
		if err != nil {
			return nil, err
		}
		ids = append(ids, categoryID)
	}
	return ids, nil
}

func (c CategoryID) Value() string {
	return string(c)
}

func (cID CategoryID) Equals(cID2 CategoryID) bool {
	return cID == cID2
}
