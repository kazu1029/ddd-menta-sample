package repoimpl

import (
	"strconv"

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
	tag1, _ := tagdm.NewTag(tagdm.TagID("id0"), "Go")
	tag2, _ := tagdm.NewTag(tagdm.TagID("id1"), "PHP")
	tag3, _ := tagdm.NewTag(tagdm.TagID("id2"), "AWS")
	tag4, _ := tagdm.NewTag(tagdm.TagID("id3"), "JavaScript")
	tag5, _ := tagdm.NewTag(tagdm.TagID("id4"), "GCP")
	tags["id0"] = tag1
	tags["id1"] = tag2
	tags["id2"] = tag3
	tags["id3"] = tag4
	tags["idd4"] = tag5
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

func (repo *TagRepoImpl) FindAll() ([]*tagdm.Tag, error) {
	tagSlice := make([]*tagdm.Tag, len(tags))
	for i := 0; i < len(categories); i++ {
		tagSlice[i] = tags["id"+strconv.Itoa(i)]
	}
	return tagSlice, nil
}
