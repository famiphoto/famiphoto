package schema

import (
	"fmt"
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
