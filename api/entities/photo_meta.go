package entities

import (
	"fmt"
	"github.com/famiphoto/famiphoto/api/config"
	"github.com/famiphoto/famiphoto/api/utils"
	"github.com/labstack/gommon/log"
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

func (m PhotoMeta) DateTimeOriginal() int64 {
	item := m.findByTagID(utils.ExifTagIDDateTimeOriginal)
	if item == nil {
		return 0
	}
	at, err := utils.ParseDatetime(item.ValueString, utils.MustLoadLocation(config.Env.ExifTimezone))
	if err != nil {
		log.Error("failed to parse date time original exif data" + err.Error())
		return 0
	}
	return at.Unix()
}

func (m PhotoMeta) findByTagID(tagID int64) *PhotoMetaItem {
	for _, item := range m {
		if item.TagID == tagID {
			return item
		}
	}
	return nil
}

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
		fmt.Println("tagid", list[i].TagID, item.TagId)
	}
	return list
}
