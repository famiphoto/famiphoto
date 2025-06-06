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

// Camera information
const ExifTagMake = 271           // 0x010f
const ExifTagModel = 272          // 0x0110
const ExifTagSerialNumber = 42033 // 0xa805

// Date and time information
const ExifTagDateTimeOriginal = 36867 // 0x9003
const ExifTagDateTimeDigitized = 36868 // 0x9004
const ExifTagCreateDate = 36868 // 0x9004 Same as DateTimeDigitized
const ExifTagSubsecTimeOriginal = 37521 // 0x9291
const ExifTagTimezoneOffset = 34858 // 0x882a OffsetTime

// Shooting settings
const ExifTagExposureTime = 33434 // 0x829a
const ExifTagFNumber = 33437 // 0x829d
const ExifTagISO = 34855 // 0x8827
const ExifTagFocalLength = 37386 // 0x920a
const ExifTagFocalLengthIn35mm = 41989 // 0xa405
const ExifTagExposureProgram = 34850 // 0x8822
const ExifTagExposureCompensation = 37380 // 0x9204
const ExifTagMeteringMode = 37383 // 0x9207
const ExifTagFlash = 37385 // 0x9209

// Lens information
const ExifTagLensMake = 42035 // 0xa433
const ExifTagLensModel = 42036 // 0xa434
const ExifTagLensSerialNumber = 42037 // 0xa435

// Image information
const ExifTagWidth = 256 // 0x0100
const ExifTagHeight = 257 // 0x0101
const ExifTagColorSpace = 40961 // 0xa001
const ExifTagWhiteBalance = 41987 // 0xa403
const ExifTagOrientation = 274 // 0x0112

// GPS information
const ExifTagGPSLatitude = 2 // 0x0002
const ExifTagGPSLongitude = 4 // 0x0004
const ExifTagGPSAltitude = 6 // 0x0006

// Software information
const ExifTagSoftware = 305 // 0x0131
const ExifTagFirmware = 42016 // 0xa420 Firmware version

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
