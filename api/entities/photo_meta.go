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
	item := m.findByTagID(utils.ExifTagDateTimeOriginal)
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

// Camera information
func (m PhotoMeta) Make() string {
	item := m.findByTagID(utils.ExifTagMake)
	if item == nil {
		return ""
	}
	return item.ValueString
}

func (m PhotoMeta) Model() string {
	item := m.findByTagID(utils.ExifTagModel)
	if item == nil {
		return ""
	}
	return item.ValueString
}

func (m PhotoMeta) SerialNumber() string {
	item := m.findByTagID(utils.ExifTagSerialNumber)
	if item == nil {
		return ""
	}
	return item.ValueString
}

// Date and time information
func (m PhotoMeta) DateTimeDigitized() int64 {
	item := m.findByTagID(utils.ExifTagDateTimeDigitized)
	if item == nil {
		return 0
	}
	at, err := utils.ParseDatetime(item.ValueString, utils.MustLoadLocation(config.Env.ExifTimezone))
	if err != nil {
		log.Error("failed to parse date time digitized exif data" + err.Error())
		return 0
	}
	return at.Unix()
}

func (m PhotoMeta) CreateDate() int64 {
	return m.DateTimeDigitized() // Same as DateTimeDigitized
}

func (m PhotoMeta) SubsecTimeOriginal() string {
	item := m.findByTagID(utils.ExifTagSubsecTimeOriginal)
	if item == nil {
		return ""
	}
	return item.ValueString
}

func (m PhotoMeta) TimezoneOffset() string {
	item := m.findByTagID(utils.ExifTagTimezoneOffset)
	if item == nil {
		return ""
	}
	return item.ValueString
}

// Shooting settings
func (m PhotoMeta) ExposureTime() string {
	item := m.findByTagID(utils.ExifTagExposureTime)
	if item == nil {
		return ""
	}
	return item.ValueString
}

func (m PhotoMeta) FNumber() string {
	item := m.findByTagID(utils.ExifTagFNumber)
	if item == nil {
		return ""
	}
	return item.ValueString
}

func (m PhotoMeta) ISO() int64 {
	item := m.findByTagID(utils.ExifTagISO)
	if item == nil {
		return 0
	}
	return item.ValueInt()
}

func (m PhotoMeta) FocalLength() string {
	item := m.findByTagID(utils.ExifTagFocalLength)
	if item == nil {
		return ""
	}
	return item.ValueString
}

func (m PhotoMeta) FocalLengthIn35mm() int64 {
	item := m.findByTagID(utils.ExifTagFocalLengthIn35mm)
	if item == nil {
		return 0
	}
	return item.ValueInt()
}

func (m PhotoMeta) ExposureProgram() int64 {
	item := m.findByTagID(utils.ExifTagExposureProgram)
	if item == nil {
		return 0
	}
	return item.ValueInt()
}

func (m PhotoMeta) ExposureCompensation() string {
	item := m.findByTagID(utils.ExifTagExposureCompensation)
	if item == nil {
		return ""
	}
	return item.ValueString
}

func (m PhotoMeta) MeteringMode() int64 {
	item := m.findByTagID(utils.ExifTagMeteringMode)
	if item == nil {
		return 0
	}
	return item.ValueInt()
}

func (m PhotoMeta) Flash() int64 {
	item := m.findByTagID(utils.ExifTagFlash)
	if item == nil {
		return 0
	}
	return item.ValueInt()
}

// Lens information
func (m PhotoMeta) LensMake() string {
	item := m.findByTagID(utils.ExifTagLensMake)
	if item == nil {
		return ""
	}
	return item.ValueString
}

func (m PhotoMeta) LensModel() string {
	item := m.findByTagID(utils.ExifTagLensModel)
	if item == nil {
		return ""
	}
	return item.ValueString
}

func (m PhotoMeta) LensSerialNumber() string {
	item := m.findByTagID(utils.ExifTagLensSerialNumber)
	if item == nil {
		return ""
	}
	return item.ValueString
}

// Image information
func (m PhotoMeta) Width() int64 {
	item := m.findByTagID(utils.ExifTagWidth)
	if item == nil {
		return 0
	}
	return item.ValueInt()
}

func (m PhotoMeta) Height() int64 {
	item := m.findByTagID(utils.ExifTagHeight)
	if item == nil {
		return 0
	}
	return item.ValueInt()
}

func (m PhotoMeta) ColorSpace() int64 {
	item := m.findByTagID(utils.ExifTagColorSpace)
	if item == nil {
		return 0
	}
	return item.ValueInt()
}

func (m PhotoMeta) WhiteBalance() int64 {
	item := m.findByTagID(utils.ExifTagWhiteBalance)
	if item == nil {
		return 0
	}
	return item.ValueInt()
}

func (m PhotoMeta) Orientation() int64 {
	item := m.findByTagID(utils.ExifTagOrientation)
	if item == nil {
		return 0
	}
	return item.ValueInt()
}

// GPS information
func (m PhotoMeta) GPSLatitude() string {
	item := m.findByTagID(utils.ExifTagGPSLatitude)
	if item == nil {
		return ""
	}
	return item.ValueString
}

func (m PhotoMeta) GPSLongitude() string {
	item := m.findByTagID(utils.ExifTagGPSLongitude)
	if item == nil {
		return ""
	}
	return item.ValueString
}

func (m PhotoMeta) GPSAltitude() string {
	item := m.findByTagID(utils.ExifTagGPSAltitude)
	if item == nil {
		return ""
	}
	return item.ValueString
}

// Software information
func (m PhotoMeta) Software() string {
	item := m.findByTagID(utils.ExifTagSoftware)
	if item == nil {
		return ""
	}
	return item.ValueString
}

func (m PhotoMeta) Firmware() string {
	item := m.findByTagID(utils.ExifTagFirmware)
	if item == nil {
		return ""
	}
	return item.ValueString
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
