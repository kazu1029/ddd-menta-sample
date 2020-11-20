package tagdm

import (
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type TagID string

func NewTagID() TagID {
	return TagID(uuid.New().String())
}

func NewTagIDWithStr(id string) (TagID, error) {
	if id == "" {
		return "", xerrors.New("tag id must be not empty")
	}

	return TagID(id), nil
}

func NewTagIDs(tagIDs []string) ([]TagID, error) {
	var ids []TagID
	for _, s := range tagIDs {
		tagID, err := NewTagIDWithStr(s)
		if err != nil {
			return nil, err
		}
		ids = append(ids, tagID)
	}
	return ids, nil
}

func (s TagID) Value() string {
	return string(s)
}

func (sID TagID) Equals(sID2 TagID) bool {
	return sID == sID2
}
