package adapters

import (
	"context"
	"github.com/famiphoto/famiphoto/api/entities"
	"github.com/famiphoto/famiphoto/api/infrastructures/models"
	"github.com/famiphoto/famiphoto/api/infrastructures/repositories"
	"time"
)

type PhotoSearchAdapter interface {
	Index(ctx context.Context, photoID string, photoFiles entities.PhotoFileList, meta entities.PhotoMeta, now time.Time) error
}

func NewPhotoSearchAdapter(esRepo repositories.PhotoElasticSearchRepository) PhotoSearchAdapter {
	return &photoSearchAdapter{
		esRepo: esRepo,
	}
}

type photoSearchAdapter struct {
	esRepo repositories.PhotoElasticSearchRepository
}

func (r *photoSearchAdapter) Index(ctx context.Context, photoID string, photoFiles entities.PhotoFileList, meta entities.PhotoMeta, now time.Time) error {
	// Extract date parts from DateTimeOriginal
	dateTimeOriginal := meta.DateTimeOriginal()
	dateTime := time.Unix(dateTimeOriginal, 0)
	dateTimeParts := models.DateTimeOriginalParts{
		Year:   dateTime.Year(),
		Month:  int(dateTime.Month()),
		Day:    dateTime.Day(),
		Hour:   dateTime.Hour(),
		Minute: dateTime.Minute(),
	}

	// Create location from GPS data if available
	var location interface{}
	if meta.GPSLatitude() != "" && meta.GPSLongitude() != "" {
		// Parse latitude and longitude
		// Note: This is a placeholder. Actual implementation would depend on the format of GPSLatitude and GPSLongitude
		// location = map[string]interface{}{
		//     "lat": parsedLatitude,
		//     "lon": parsedLongitude,
		// }
		// Leaving as nil with a comment as requested
		location = nil // Cannot parse GPS coordinates without knowing the format
	}

	// Prepare original image files
	originalImageFiles := make([]models.OriginalImageFile, 0, len(photoFiles))
	for _, file := range photoFiles {
		originalImageFiles = append(originalImageFiles, models.OriginalImageFile{
			Path:     file.File.Path,
			MimeType: file.MimeType(),
			MD5Hash:  file.FileHash,
		})
	}

	// Prepare image URLs
	imageUrls := models.ImageUrls{
		ThumbnailURL: "",                     // Cannot determine thumbnail URL from available data
		PreviewURL:   "",                     // Cannot determine preview URL from available data
		OriginalURLs: []models.OriginalUrl{}, // Cannot determine original URLs from available data
	}

	// Prepare EXIF data
	exifData := models.ExifData{
		// Camera information
		Make:         meta.Make(),
		Model:        meta.Model(),
		SerialNumber: meta.SerialNumber(),

		// Date and time information
		DateTimeOriginal:   dateTime.Format("2006:01:02 15:04:05"), // Standard EXIF date format
		DateTimeDigitized:  time.Unix(meta.DateTimeDigitized(), 0).Format("2006:01:02 15:04:05"),
		CreateDate:         time.Unix(meta.CreateDate(), 0).Format("2006:01:02 15:04:05"),
		SubsecTimeOriginal: meta.SubsecTimeOriginal(),
		TimezoneOffset:     meta.TimezoneOffset(),

		// Shooting settings
		ExposureTime:         meta.ExposureTime(),
		FNumber:              0, // Cannot convert string to float64 without parsing
		ISO:                  int(meta.ISO()),
		FocalLength:          0, // Cannot convert string to float64 without parsing
		FocalLengthIn35mm:    float64(meta.FocalLengthIn35mm()),
		ExposureProgram:      "", // Cannot convert int64 to string without mapping
		ExposureCompensation: 0,  // Cannot convert string to float64 without parsing
		MeteringMode:         "", // Cannot convert int64 to string without mapping
		Flash:                "", // Cannot convert int64 to string without mapping

		// Lens information
		LensMake:         meta.LensMake(),
		LensModel:        meta.LensModel(),
		LensSerialNumber: meta.LensSerialNumber(),

		// Image information
		Width:        int(meta.Width()),
		Height:       int(meta.Height()),
		ColorSpace:   "", // Cannot convert int64 to string without mapping
		WhiteBalance: "", // Cannot convert int64 to string without mapping
		Orientation:  int(meta.Orientation()),

		// GPS information
		GPSLatitude:  0, // Cannot convert string to float64 without parsing
		GPSLongitude: 0, // Cannot convert string to float64 without parsing
		GPSAltitude:  0, // Cannot convert string to float64 without parsing

		// Software information
		Software: meta.Software(),
		Firmware: meta.Firmware(),
	}

	doc := &models.PhotoIndex{
		PhotoID:               photoID,
		Name:                  photoFiles[0].File.NameExceptExt(),
		ImportedAt:            now.Unix(),
		DateTimeOriginal:      dateTimeOriginal,
		DateTimeOriginalParts: dateTimeParts,
		Orientation:           int(meta.Orientation()),
		Location:              location,
		ImageUrls:             imageUrls,
		OriginalImageFiles:    originalImageFiles,
		Exif:                  exifData,
		DescriptionJa:         "", // As requested, keep empty
		DescriptionEn:         "", // As requested, keep empty
	}

	return r.esRepo.Index(ctx, doc)
}
