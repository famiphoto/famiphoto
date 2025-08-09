package routers

import "github.com/labstack/echo/v4"

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
