package routers

import (
	"github.com/famiphoto/famiphoto/api/entities"
	"github.com/famiphoto/famiphoto/api/errors"
	"github.com/famiphoto/famiphoto/api/interfaces/http/schema"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *handler) PhotosGetPhotoList(ctx echo.Context, params schema.PhotosGetPhotoListParams) error {
	q := &entities.PhotoSearchQuery{
		Limit:  params.GetLimitOrDefault(30),
		Offset: params.GetOffsetOrDefault(0),
	}
	result, err := h.photoSearchUseCase.Search(ctx.Request().Context(), q)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, schema.NewPhotosGetPhotoListResponse(result))
}

func (h *handler) PhotosGetPhoto(ctx echo.Context, photoId string) error {
	result, err := h.photoSearchUseCase.GetByPhotoID(ctx.Request().Context(), photoId)
	if err != nil {
		return err
	}
	if len(result.Items) == 0 {
		return errors.New(errors.PhotoNotFoundError, nil)
	}

	return ctx.JSON(http.StatusOK, schema.NewPhotosGetPhotoResponse(result.Items[0]))
}
