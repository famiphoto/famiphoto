package routers

import (
	"github.com/labstack/echo/v4"
	"path/filepath"
)

func (h *handler) AssetsGetPreview(ctx echo.Context, photoId string) error {
	filePath, err := h.assetUseCase.GetPreview(ctx.Request().Context(), photoId)
	if err != nil {
		return err
	}

	return ctx.File(filePath)
}

func (h *handler) AssetsGetThumbnail(ctx echo.Context, photoId string) error {
	filePath, err := h.assetUseCase.GetThumbnail(ctx.Request().Context(), photoId)
	if err != nil {
		return err
	}

	return ctx.File(filePath)
}

func (h *handler) AssetsGetOriginalFile(ctx echo.Context, photoFileId string) error {
	filePath, err := h.assetUseCase.GetOriginalFile(ctx.Request().Context(), photoFileId)
	if err != nil {
		return err
	}

	return ctx.Attachment(filePath, filepath.Base(filePath))
}
