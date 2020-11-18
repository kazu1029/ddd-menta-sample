package repoimpl

import (
	"github.com/kazu1029/ddd-menta-sample/src/core/domain/tagdm"
)

type TagRepoImpl struct{}

func NewTagRepoImpl() *TagRepoImpl {
	return &TagRepoImpl{}
}

var (
	tags map[string]*tagdm.Tag
)

func init() {
	tag1, _ := tagdm.NewTag(tagdm.TagID("id1"), "Go")
	tag2, _ := tagdm.NewTag(tagdm.TagID("id2"), "PHP")
	tag3, _ := tagdm.NewTag(tagdm.TagID("id3"), "AWS")
	tag4, _ := tagdm.NewTag(tagdm.TagID("id4"), "JavaScript")
	tag5, _ := tagdm.NewTag(tagdm.TagID("id5"), "GCP")
	tags["id1"] = tag1
	tags["id2"] = tag2
	tags["id3"] = tag3
	tags["id4"] = tag4
	tags["id5"] = tag5
}

func (repo *TagRepoImpl) FindByID(tagID tagdm.TagID) (*tagdm.Tag, error) {
	return tags[tagID.Value()], nil
}

func (repo *TagRepoImpl) FindByIDs(tagIDs []tagdm.TagID) ([]*tagdm.Tag, error) {
	var fetchedTags []*tagdm.Tag
	for _, tagID := range tagIDs {
		if _, ok := tags[tagID.Value()]; ok {
			fetchedTags = append(fetchedTags, tags[tagID.Value()])
		}
	}
	return fetchedTags, nil
}
