package adapters

import (
	"context"
	"github.com/famiphoto/famiphoto/api/entities"
	"github.com/famiphoto/famiphoto/api/errors"
	"github.com/famiphoto/famiphoto/api/infrastructures/dbmodels"
	"github.com/famiphoto/famiphoto/api/infrastructures/repositories"
)

type PhotoMetaAdapter interface {
	Upsert(ctx context.Context, photoID int64, meta entities.PhotoMeta) error
}

func NewPhotoMetaAdapter(photoExifRepo repositories.PhotoExifRepository) PhotoMetaAdapter {
	return &photoMetaAdapter{photoExifRepo: photoExifRepo}
}

type photoMetaAdapter struct {
	photoExifRepo repositories.PhotoExifRepository
}

func (a *photoMetaAdapter) Upsert(ctx context.Context, photoID int64, meta entities.PhotoMeta) error {
	for _, item := range meta {
		dbModel := &dbmodels.PhotoExif{
			PhotoExifID: 0,
			PhotoID:     photoID,
			TagID:       int(item.TagID),
			TagName:     item.TagName,
			TagType:     item.TagType,
			ValueString: item.ValueString,
			SortOrder:   0,
		}
		row, err := a.photoExifRepo.GetPhotoExifByPhotoIDTagID(ctx, photoID, item.TagID)
		if err != nil {
			if !errors.IsErrCode(err, errors.DBNotFoundError) {
				return err
			}
			if _, err := a.photoExifRepo.Insert(ctx, dbModel); err != nil {
				return err
			}
			continue
		}

		dbModel.PhotoExifID = row.PhotoExifID
		if _, err := a.photoExifRepo.Update(ctx, dbModel); err != nil {
			return err
		}
	}
	return nil
}
