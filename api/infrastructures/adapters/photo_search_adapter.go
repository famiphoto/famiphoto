package adapters

import (
	"context"
	"github.com/famiphoto/famiphoto/api/entities"
	"github.com/famiphoto/famiphoto/api/infrastructures/models"
	"github.com/famiphoto/famiphoto/api/infrastructures/repositories"
	"github.com/famiphoto/famiphoto/api/utils/exif"
	"time"
)

type PhotoSearchAdapter interface {
	CreateIndexIfNotExist(ctx context.Context) error
	Index(ctx context.Context, photoID string, photoFiles entities.PhotoFileList, meta exif.ExifData, now time.Time) error
}

func NewPhotoSearchAdapter(esRepo repositories.PhotoElasticSearchRepository) PhotoSearchAdapter {
	return &photoSearchAdapter{
		esRepo: esRepo,
	}
}

type photoSearchAdapter struct {
	esRepo repositories.PhotoElasticSearchRepository
}

func (r *photoSearchAdapter) CreateIndexIfNotExist(ctx context.Context) error {
	if exist, err := r.esRepo.ExistsIndex(ctx); err != nil {
		return err
	} else if exist {
		return nil
	}
	return r.esRepo.CreateIndex(ctx)
}

func (r *photoSearchAdapter) Index(ctx context.Context, photoID string, photoFiles entities.PhotoFileList, meta exif.ExifData, now time.Time) error {
	// Extract date parts from DateTimeOriginal
	dateTimeOriginal, err := exif.ParseDatetime(meta.DateTimeOriginal(), meta.OffsetTimeOriginal())
	if err != nil {
		dateTimeOriginal = time.Unix(0, 0)
	}
	
	dateTimeParts := models.DateTimeOriginalParts{
		Year:   dateTimeOriginal.Year(),
		Month:  int(dateTimeOriginal.Month()),
		Day:    dateTimeOriginal.Day(),
		Hour:   dateTimeOriginal.Hour(),
		Minute: dateTimeOriginal.Minute(),
	}

	createDate, err := exif.ParseDatetime(meta.CreateDate(), meta.OffsetTimeOriginal())
	if err != nil {
		createDate = time.Unix(0, 0)
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

	// Prepare EXIF data
	exifData := models.ExifData{
		// Camera information
		Make:         meta.Make(),
		Model:        meta.Model(),
		SerialNumber: meta.SerialNumber(),

		// Date and time information
		DateTimeOriginal:      dateTimeOriginal.Unix(),
		DateTimeOriginalParts: dateTimeParts,
		CreateDate:            createDate.Unix(),
		SubsecTimeOriginal:    meta.SubsecTimeOriginal(),
		TimezoneOffset:        meta.TimezoneOffset(),

		// Shooting settings
		ExposureTime:         meta.ExposureTime(),
		FNumber:              meta.FNumber(),
		ISO:                  int(meta.ISO()),
		FocalLength:          meta.FocalLength(),
		FocalLengthIn35mm:    float64(meta.FocalLengthIn35mm()),
		ExposureProgram:      meta.ExposureProgram(),
		ExposureCompensation: meta.ExposureCompensation(),
		MeteringMode:         meta.MeteringMode(),
		Flash:                meta.Flash(),

		// Lens information
		LensMake:         meta.LensMake(),
		LensModel:        meta.LensModel(),
		LensSerialNumber: meta.LensSerialNumber(),

		// Image information
		Width:        int(meta.Width()),
		Height:       int(meta.Height()),
		ColorSpace:   meta.ColorSpace(),
		WhiteBalance: meta.WhiteBalance(),
		Orientation:  int(meta.Orientation()),

		// Software information
		Software: meta.Software(),
		Firmware: meta.Firmware(),
	}

	doc := &models.PhotoIndex{
		PhotoID:               photoID,
		Name:                  photoFiles[0].File.NameExceptExt(),
		ImportedAt:            now.Unix(),
		DateTimeOriginal:      dateTimeOriginal.Unix(),
		DateTimeOriginalParts: dateTimeParts,
		Orientation:           int(meta.Orientation()),
		OriginalImageFiles:    originalImageFiles,
		Exif:                  exifData,
		DescriptionJa:         "", // 他関数で遅延処理セット
		DescriptionEn:         "", // 他関数で遅延処理セット
	}

	return r.esRepo.Index(ctx, doc)
}
