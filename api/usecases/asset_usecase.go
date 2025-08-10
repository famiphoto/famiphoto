package usecases

import (
	"context"
	"github.com/famiphoto/famiphoto/api/config"
	"github.com/famiphoto/famiphoto/api/infrastructures/adapters"
	"path"
)

type AssetUseCase interface {
	GetPreview(ctx context.Context, photoID string) (string, error)
	GetThumbnail(ctx context.Context, photoID string) (string, error)
	GetOriginalFile(ctx context.Context, photoFileID string) (string, error)
}

func NewAssetUseCase(
	photoAdapter adapters.PhotoAdapter,
	photoFileAdapter adapters.PhotoFileAdapter,
	photoStorageAdapter adapters.PhotoStorageAdapter,
) AssetUseCase {
	return &assetUseCase{
		photoAdapter:        photoAdapter,
		photoFileAdapter:    photoFileAdapter,
		photoStorageAdapter: photoStorageAdapter,
	}
}

type assetUseCase struct {
	photoAdapter        adapters.PhotoAdapter
	photoFileAdapter    adapters.PhotoFileAdapter
	photoStorageAdapter adapters.PhotoStorageAdapter
}

func (u *assetUseCase) GetPreview(ctx context.Context, photoID string) (string, error) {
	// DBから写真情報の存在を確認（将来的に拡張のために取得）
	if _, err := u.photoAdapter.FindByID(ctx, photoID); err != nil {
		return "", err
	}
	filePath := path.Join(config.Env.AssetRootPath, "previews", photoID)
	return filePath, nil
}

func (u *assetUseCase) GetThumbnail(ctx context.Context, photoID string) (string, error) {
	// DBから写真情報の存在を確認（将来的に拡張のために取得）
	if _, err := u.photoAdapter.FindByID(ctx, photoID); err != nil {
		return "", err
	}
	filePath := path.Join(config.Env.AssetRootPath, "thumbnail", photoID)
	return filePath, nil
}

func (u *assetUseCase) GetOriginalFile(ctx context.Context, photoFileID string) (string, error) {
	photoFile, err := u.photoFileAdapter.FindByPhotoFileID(ctx, photoFileID)
	if err != nil {
		return "", err
	}

	storageFile, err := u.photoStorageAdapter.GetFileInfo(photoFile)
	if err != nil {
		return "", err
	}
	
	return storageFile.Path, nil
}
