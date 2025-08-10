package adapters

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/famiphoto/famiphoto/api/entities"
	"github.com/famiphoto/famiphoto/api/infrastructures/models"
	"github.com/famiphoto/famiphoto/api/infrastructures/repositories"
	"github.com/famiphoto/famiphoto/api/utils/cast"
	"github.com/famiphoto/famiphoto/api/utils/exif"
	"time"
)

type PhotoSearchAdapter interface {
	CreateIndexIfNotExist(ctx context.Context) error
	Index(ctx context.Context, photoID string, photoFiles entities.PhotoFileList, meta exif.ExifData, now time.Time) error
	Search(ctx context.Context, query *entities.PhotoSearchQuery) (*entities.PhotoSearchResult, error)
}

func NewPhotoSearchAdapter(esRepo repositories.PhotoElasticSearchRepository) PhotoSearchAdapter {
	return &photoSearchAdapter{
		esRepo: esRepo,
	}
}

type photoSearchAdapter struct {
	esRepo repositories.PhotoElasticSearchRepository
}

func (a *photoSearchAdapter) CreateIndexIfNotExist(ctx context.Context) error {
	if exist, err := a.esRepo.ExistsIndex(ctx); err != nil {
		return err
	} else if exist {
		return nil
	}
	return a.esRepo.CreateIndex(ctx)
}

func (a *photoSearchAdapter) Index(ctx context.Context, photoID string, photoFiles entities.PhotoFileList, meta exif.ExifData, now time.Time) error {
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
			PhotoFileID: file.PhotoFileID,
			Path:        file.File.Path,
			MimeType:    file.MimeType(),
			MD5Hash:     file.FileHash,
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

	return a.esRepo.Index(ctx, doc)
}

func (a *photoSearchAdapter) Search(ctx context.Context, query *entities.PhotoSearchQuery) (*entities.PhotoSearchResult, error) {
	req := a.createSearchRequest(query)
	hits, total, err := a.esRepo.Search(ctx, req)
	if err != nil {
		return nil, err
	}

	return &entities.PhotoSearchResult{
		Limit:  query.Limit,
		Offset: query.Offset,
		Total:  total,
		Items:  hits,
	}, nil
}

func (a *photoSearchAdapter) createSearchRequest(query *entities.PhotoSearchQuery) *search.Request {
	sortDesc := "desc"
	req := &search.Request{
		Size: cast.Ptr(int(query.Limit)),
		From: cast.Ptr(int(query.Offset)),
		Sort: []types.SortCombinations{
			map[string]interface{}{
				"date_time_original": map[string]string{
					"order": sortDesc,
				},
			},
		},
	}

	if query.PhotoID != "" {
		req.Query = &types.Query{
			Bool: &types.BoolQuery{
				Filter: []types.Query{
					{
						Term: map[string]types.TermQuery{
							"photo_id": {Value: query.PhotoID},
						},
					},
				},
			},
		}
	}

	return req
}
