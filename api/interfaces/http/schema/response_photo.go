package schema

import (
	"fmt"
	"github.com/famiphoto/famiphoto/api/infrastructures/models"
	"path"
	"time"

	"github.com/famiphoto/famiphoto/api/config"
	"github.com/famiphoto/famiphoto/api/entities"
	"github.com/famiphoto/famiphoto/api/utils"
)

func NewPhotosGetPhotoListResponse(result *entities.PhotoSearchResult) *PhotosPhotoListResponse {
	if result == nil {
		return &PhotosPhotoListResponse{
			Items:  []PhotosPhotoItem{},
			Offset: 0,
			Total:  0,
		}
	}

	items := make([]PhotosPhotoItem, 0, len(result.Items))
	for _, p := range result.Items {
		if p == nil {
			continue
		}
		// Determine timezone location from EXIF offset or default config
		loc := utils.LocationOrDefaultFromOffset(p.Exif.TimezoneOffset, utils.MustLoadLocation(config.Env.ExifTimezone))
		// Build times
		dateTimeOriginal := time.Unix(p.DateTimeOriginal, 0).In(loc)
		importedAt := time.Unix(p.ImportedAt, 0).In(loc).Format(time.RFC3339)

		// Build asset URLs
		previewURL := fmt.Sprintf("%s/previews/%s", config.Env.AssetBaseURL, p.PhotoID)
		thumbnailURL := fmt.Sprintf("%s/thumbnails/%s", config.Env.AssetBaseURL, p.PhotoID)

		items = append(items, PhotosPhotoItem{
			PhotoId:          p.PhotoID,
			Name:             p.Name,
			ImportedAt:       importedAt,
			DateTimeOriginal: dateTimeOriginal,
			PreviewUrl:       previewURL,
			ThumbnailUrl:     thumbnailURL,
		})
	}

	return &PhotosPhotoListResponse{
		Items:  items,
		Offset: result.Offset,
		Total:  result.Total,
	}
}

func NewPhotosGetPhotoResponse(item *models.PhotoIndex) *PhotosPhoto {
	// Prepare an empty default response to avoid nil slices
	if item == nil {
		return &PhotosPhoto{
			ExifData:  PhotosExifData{},
			FileTypes: []string{},
			Files:     []PhotosFile{},
		}
	}

	// Determine timezone location from EXIF offset or default config
	loc := utils.LocationOrDefaultFromOffset(item.Exif.TimezoneOffset, utils.MustLoadLocation(config.Env.ExifTimezone))

	// Build times
	dateTimeOriginal := time.Unix(item.DateTimeOriginal, 0).In(loc)
	importedAt := time.Unix(item.ImportedAt, 0).In(loc).Format(time.RFC3339)

	// Build asset URLs
	previewURL := fmt.Sprintf("%s/previews/%s", config.Env.AssetBaseURL, item.PhotoID)
	thumbnailURL := fmt.Sprintf("%s/thumbnails/%s", config.Env.AssetBaseURL, item.PhotoID)

	// Map original files to schema files and collect file types
	files := make([]PhotosFile, 0, len(item.OriginalImageFiles))
	fileTypes := make([]string, 0, len(item.OriginalImageFiles))
	for _, f := range item.OriginalImageFiles {
		files = append(files, PhotosFile{
			FileHash: f.MD5Hash,
			FileId:   f.PhotoFileID,
			FileName: path.Base(f.Path),
			FileType: f.MimeType,
			PhotoId:  item.PhotoID,
		})
		fileTypes = append(fileTypes, f.MimeType)
	}

	// Map EXIF data
	exif := PhotosExifData{
		Make:                 item.Exif.Make,
		Model:                item.Exif.Model,
		SerialNumber:         item.Exif.SerialNumber,
		DateTimeOriginal:     item.Exif.DateTimeOriginal,
		CreateDate:           item.Exif.CreateDate,
		SubsecTimeOriginal:   item.Exif.SubsecTimeOriginal,
		TimezoneOffset:       item.Exif.TimezoneOffset,
		ExposureTime:         item.Exif.ExposureTime,
		FNumber:              item.Exif.FNumber,
		Iso:                  int64(item.Exif.ISO),
		FocalLength:          item.Exif.FocalLength,
		FocalLengthIn35mm:    item.Exif.FocalLengthIn35mm,
		ExposureProgram:      item.Exif.ExposureProgram,
		ExposureCompensation: item.Exif.ExposureCompensation,
		MeteringMode:         item.Exif.MeteringMode,
		Flash:                item.Exif.Flash,
		LensMake:             item.Exif.LensMake,
		LensModel:            item.Exif.LensModel,
		LensSerialNumber:     item.Exif.LensSerialNumber,
		Width:                int64(item.Exif.Width),
		Height:               int64(item.Exif.Height),
		ColorSpace:           item.Exif.ColorSpace,
		WhiteBalance:         item.Exif.WhiteBalance,
		Orientation:          PhotosPhotoOrientation(item.Exif.Orientation),
		Software:             item.Exif.Software,
		Firmware:             item.Exif.Firmware,
	}

	return &PhotosPhoto{
		PhotoId:          item.PhotoID,
		Name:             item.Name,
		ImportedAt:       importedAt,
		DateTimeOriginal: dateTimeOriginal,
		PreviewUrl:       previewURL,
		ThumbnailUrl:     thumbnailURL,
		ExifData:         exif,
		Files:            files,
		FileTypes:        fileTypes,
	}
}
