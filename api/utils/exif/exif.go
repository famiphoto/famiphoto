package exif

import (
	"github.com/dsoprea/go-exif/v3"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
	"github.com/famiphoto/famiphoto/api/errors"
	"github.com/famiphoto/famiphoto/api/utils"
	"time"
)

func ParseDatetime(val string, offset string) (time.Time, error) {
	loc := utils.LocationOrDefaultFromOffset(offset, utils.MustLoadLocation("Asia/Tokyo"))
	return time.ParseInLocation("2006:01:02 15:04:05", val, loc)
}

func ParseExifItemsAll(data []byte) (ExifData, error) {
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

func (i ExifItem) ParseRational() float64 {
	rationals, ok := i.Value.([]exifcommon.Rational)
	if !ok {
		return 0
	}
	if len(rationals) == 0 {
		return 0
	}
	return float64(rationals[0].Numerator) / float64(rationals[0].Denominator)
}

func (i ExifItem) ParseSignedRational() float64 {
	rationals, ok := i.Value.([]exifcommon.SignedRational)
	if !ok {
		return 0
	}
	if len(rationals) == 0 {
		return 0
	}
	return float64(rationals[0].Numerator) / float64(rationals[0].Denominator)
}

type ExifData []*ExifItem

func (m ExifData) findByTagID(tagID uint16) *ExifItem {
	for _, item := range m {
		if item.TagId == tagID {
			return item
		}
	}
	return nil
}

func (m ExifData) Make() string {
	item := m.findByTagID(ExifTagMake)
	if item == nil {
		return ""
	}
	return item.ValueString
}

func (m ExifData) Model() string {
	item := m.findByTagID(ExifTagModel)
	if item == nil {
		return ""
	}
	return item.ValueString
}

func (m ExifData) FNumber() float64 {
	item := m.findByTagID(ExifTagFNumber)
	if item == nil {
		return 0
	}
	return item.ParseRational()
}

func (m ExifData) SerialNumber() string {
	item := m.findByTagID(ExifTagSerialNumber)
	if item == nil {
		return ""
	}
	return item.ValueString
}

func (m ExifData) DateTimeOriginal() string {
	item := m.findByTagID(ExifTagDateTimeOriginal)
	if item == nil {
		return ""
	}
	return item.ValueString
}

func (m ExifData) DateTimeDigitized() string {
	item := m.findByTagID(ExifTagDateTimeDigitized)
	if item == nil {
		return ""
	}
	return item.ValueString
}

func (m ExifData) CreateDate() string {
	item := m.findByTagID(ExifTagCreateDate)
	if item == nil {
		return ""
	}
	return item.ValueString
}

func (m ExifData) SubsecTimeOriginal() string {
	item := m.findByTagID(ExifTagSubsecTimeOriginal)
	if item == nil {
		return ""
	}
	return item.ValueString
}

func (m ExifData) TimezoneOffset() string {
	item := m.findByTagID(ExifTagTimezoneOffset)
	if item == nil {
		return ""
	}
	return item.ValueString
}

func (m ExifData) ExposureTime() float64 {
	item := m.findByTagID(ExifTagExposureTime)
	if item == nil {
		return 0
	}
	return item.ParseRational()
}

func (m ExifData) ISO() int64 {
	item := m.findByTagID(ExifTagISO)
	if item == nil {
		return 0
	}
	if val, ok := item.Value.([]uint16); ok {
		if len(val) == 0 {
			return 0
		}
		return int64(val[0])
	}
	return 0
}

func (m ExifData) FocalLength() float64 {
	item := m.findByTagID(ExifTagFocalLength)
	if item == nil {
		return 0
	}
	return item.ParseRational()
}

func (m ExifData) FocalLengthIn35mm() int64 {
	item := m.findByTagID(ExifTagFocalLengthIn35mm)
	if item == nil {
		return 0
	}
	if val, ok := item.Value.(uint16); ok {
		return int64(val)
	}
	return 0
}

func (m ExifData) ExposureProgram() int64 {
	item := m.findByTagID(ExifTagExposureProgram)
	if item == nil {
		return 0
	}
	if val, ok := item.Value.(uint16); ok {
		return int64(val)
	}
	return 0
}

func (m ExifData) ExposureCompensation() float64 {
	item := m.findByTagID(ExifTagExposureCompensation)
	if item == nil {
		return 0
	}
	return item.ParseSignedRational()
}

func (m ExifData) MeteringMode() int64 {
	item := m.findByTagID(ExifTagMeteringMode)
	if item == nil {
		return 0
	}
	if val, ok := item.Value.(uint16); ok {
		return int64(val)
	}
	return 0
}

func (m ExifData) Flash() int64 {
	item := m.findByTagID(ExifTagFlash)
	if item == nil {
		return 0
	}
	if val, ok := item.Value.(uint16); ok {
		return int64(val)
	}
	return 0
}

func (m ExifData) LensMake() string {
	item := m.findByTagID(ExifTagLensMake)
	if item == nil {
		return ""
	}
	return item.ValueString
}

func (m ExifData) LensModel() string {
	item := m.findByTagID(ExifTagLensModel)
	if item == nil {
		return ""
	}
	return item.ValueString
}

func (m ExifData) LensSerialNumber() string {
	item := m.findByTagID(ExifTagLensSerialNumber)
	if item == nil {
		return ""
	}
	return item.ValueString
}

func (m ExifData) Width() int64 {
	item := m.findByTagID(ExifTagWidth)
	if item == nil {
		return 0
	}
	if val, ok := item.Value.(uint32); ok {
		return int64(val)
	}
	return 0
}

func (m ExifData) Height() int64 {
	item := m.findByTagID(ExifTagHeight)
	if item == nil {
		return 0
	}
	if val, ok := item.Value.(uint32); ok {
		return int64(val)
	}
	return 0
}

func (m ExifData) ColorSpace() int64 {
	item := m.findByTagID(ExifTagColorSpace)
	if item == nil {
		return 0
	}
	if val, ok := item.Value.(uint16); ok {
		return int64(val)
	}
	return 0
}

func (m ExifData) WhiteBalance() int64 {
	item := m.findByTagID(ExifTagWhiteBalance)
	if item == nil {
		return 0
	}
	if val, ok := item.Value.(uint16); ok {
		return int64(val)
	}
	return 0
}

func (m ExifData) Orientation() int64 {
	item := m.findByTagID(ExifTagOrientation)
	if item == nil {
		return 0
	}
	if val, ok := item.Value.(uint16); ok {
		return int64(val)
	}
	return 0
}

func (m ExifData) Software() string {
	item := m.findByTagID(ExifTagSoftware)
	if item == nil {
		return ""
	}
	return item.ValueString
}

func (m ExifData) Firmware() string {
	item := m.findByTagID(ExifTagFirmware)
	if item == nil {
		return ""
	}
	return item.ValueString
}

func (m ExifData) OffsetTimeOriginal() string {
	item := m.findByTagID(ExifTagOffsetTimeOriginal)
	if item == nil {
		return ""
	}
	return item.ValueString
}

const ExifTagWidth = 256                  // 0x0100
const ExifTagHeight = 257                 // 0x0101
const ExifTagMake = 271                   // 0x010f
const ExifTagModel = 272                  // 0x0110
const ExifTagOrientation = 274            // 0x0112
const ExifTagSoftware = 305               // 0x0131
const ExifTagExposureTime = 33434         // 0x829a
const ExifTagFNumber = 33437              // 0x829d
const ExifTagExposureProgram = 34850      // 0x8822
const ExifTagTimezoneOffset = 34858       // 0x882a
const ExifTagGPSInfo = 34853              // 0x8825
const ExifTagISO = 34855                  // 0x8827
const ExifTagDateTimeOriginal = 36867     // 0x9003
const ExifTagDateTimeDigitized = 36868    // 0x9004
const ExifTagCreateDate = 36868           // 0x9004 Same as DateTimeDigitized
const ExifTagOffsetTimeOriginal = 36881   // 0x9011
const ExifTagExposureCompensation = 37380 // 0x9204
const ExifTagMeteringMode = 37383         // 0x9207
const ExifTagFlash = 37385                // 0x9209
const ExifTagFocalLength = 37386          // 0x920a
const ExifTagSubsecTimeOriginal = 37521   // 0x9291
const ExifTagColorSpace = 40961           // 0xa001
const ExifTagWhiteBalance = 41987         // 0xa403
const ExifTagFocalLengthIn35mm = 41989    // 0xa405
const ExifTagFirmware = 42016             // 0xa420 Firmware version
const ExifTagSerialNumber = 42033         // 0xa805
const ExifTagLensMake = 42035             // 0xa433
const ExifTagLensModel = 42036            // 0xa434
const ExifTagLensSerialNumber = 42037     // 0xa435

/*
GPS
*/

const ExifTagGPSLatitude = 2  // 0x0002
const ExifTagGPSLongitude = 4 // 0x0004
const ExifTagGPSAltitude = 6  // 0x0006

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
