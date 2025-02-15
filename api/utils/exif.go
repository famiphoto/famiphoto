package utils

import (
	"fmt"
	"github.com/dsoprea/go-exif/v3"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
	"github.com/famiphoto/famiphoto/api/errors"
	"time"
)

func ParseDatetime(val string, loc *time.Location) (time.Time, error) {
	return time.ParseInLocation("2006:01:02 15:04:05", val, loc)
}

func ParseExifItemsAll(data []byte) (ExifItemList, error) {
	rawExif, err := exif.SearchAndExtractExif(data)
	if err != nil {
		if errors.Is(err, exif.ErrNoExif) {
			return make([]*ExifItem, 0), nil
		}
		return nil, err
	}

	entries, _, err := exif.GetFlatExifDataUniversalSearch(rawExif, nil, true)
	if err != nil {
		return nil, err
	}

	list := make([]*ExifItem, len(entries))
	for i, entry := range entries {
		list[i] = &ExifItem{
			IfdPath:     entry.IfdPath,
			TagId:       entry.TagId,
			TagName:     entry.TagName,
			TagTypeId:   entry.TagTypeId,
			TagTypeName: entry.TagTypeName,
			UnitCount:   entry.UnitCount,
			Value:       entry.Value,
			ValueString: entry.Formatted,
		}
	}

	return list, nil
}

func ParseExifItem(data []byte, exifTagID int) (*ExifItem, error) {
	list, err := ParseExifItemsAll(data)
	if err != nil {
		return nil, err
	}

	for _, item := range list {
		if item.TagId == uint16(exifTagID) {
			return item, nil
		}
	}

	return nil, errors.New(errors.NoExifError, nil)
}

func ExtractThumbnail(data []byte) ([]byte, error) {
	rawExif, err := exif.SearchAndExtractExif(data)
	if err != nil {
		// Exifデータ無し、取得失敗
		return nil, err
	}

	im, err := exifcommon.NewIfdMappingWithStandard()
	if err != nil {
		return nil, err
	}

	ti := exif.NewTagIndex()
	_, index, err := exif.Collect(im, ti, rawExif)
	if err != nil {
		return nil, err
	}

	dt, err := index.RootIfd.NextIfd().Thumbnail()
	if err != nil {
		fmt.Println("extract fail")
		return nil, err
	}

	return dt, nil
}

type ExifItem struct {
	IfdPath     string
	TagId       uint16
	TagName     string
	TagTypeId   exifcommon.TagTypePrimitive
	TagTypeName string
	UnitCount   uint32
	Value       interface{}
	ValueString string
}

type ExifItemList []*ExifItem

const ExifTagOrientation = 274
const ExifTagIDDateTimeOriginal = 36867

const (
	ExifOrientationNone                = 1 // 不要
	ExifOrientationHorizontal          = 2 // 水平方向に反転
	ExifOrientationRotate180           = 3 // 時計回りに180度回転
	ExifOrientationVertical            = 4 // 垂直方向に反転
	ExifOrientationHorizontalRotate270 = 5 // 水平方向に反転 + 時計回りに270度回転
	ExifOrientationRotate90            = 6 // 時計回りに90度回転
	ExifOrientationHorizontalRotate90  = 7 // 水平方向に反転 + 時計回りに90度回転
	ExifOrientationRotate270           = 8 // 時計回りに270度回転
)
