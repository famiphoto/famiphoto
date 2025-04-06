package routers

import (
	"github.com/famiphoto/famiphoto/api/interfaces/http/schema"
	"github.com/labstack/echo/v4"
)

func (h *handler) GetPhotos(ctx echo.Context, params schema.GetPhotosParams) error {
	panic("")
}

func (h *handler) GetPhoto(ctx echo.Context, photoId int) error {
	panic("")
}
