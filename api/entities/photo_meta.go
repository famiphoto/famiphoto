package entities

import (
	"github.com/famiphoto/famiphoto/api/utils"
	"sort"
	"strconv"
)

type PhotoMetaItem struct {
	TagID       int64
	TagName     string
	TagType     string
	ValueString string
}

func (i PhotoMetaItem) ValueInt() int64 {
	val, err := strconv.Atoi(i.ValueString)
	if err != nil {
		return 0
	}
	return int64(val)
}

func (i PhotoMetaItem) sortOrder() int64 {
	return i.TagID
}

type PhotoMeta []*PhotoMetaItem

func (m PhotoMeta) Sort() {
	sort.Slice(m, func(i, j int) bool {
		return m[i].sortOrder() < m[j].sortOrder()
	})
}

func NewPhotoMeta(exif utils.ExifItemList) PhotoMeta {
	list := make(PhotoMeta, len(exif))
	for i, item := range exif {
		list[i] = &PhotoMetaItem{
			TagID:       int64(item.TagId),
			TagName:     item.TagName,
			TagType:     item.TagTypeName,
			ValueString: item.ValueString,
		}
	}
	return list
}
