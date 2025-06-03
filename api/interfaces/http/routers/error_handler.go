package routers

import (
	"github.com/famiphoto/famiphoto/api/interfaces/http/schema"
	"net/http"

	"github.com/famiphoto/famiphoto/api/config"
	"github.com/famiphoto/famiphoto/api/errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func ErrorHandle(err error, ctx echo.Context) {
	res := newErrorHandle(err, ctx)
	if res.StatusCode >= http.StatusInternalServerError {
		log.Error(ctx.Path(), ctx.QueryString(), err)
	}

	_ = ctx.JSON(res.StatusCode, res)
}

func newErrorHandle(err error, ctx echo.Context) *schema.ErrorResponse {
	if err == nil {
		return nil
	}

	var errorMessage string
	if config.Env.AppEnv == "prod" {
		errorMessage = err.Error()
	}

	fpError := errors.UnwrapFPError(err)
	if fpError == nil {
		return &schema.ErrorResponse{
			StatusCode:   http.StatusInternalServerError,
			ErrorCode:    errors.Unknown.ToString(),
			ErrorMessage: &errorMessage,
		}
	}
	statusCode := getHTTPStatusCode(fpError)

	return &schema.ErrorResponse{
		StatusCode:   statusCode,
		ErrorCode:    fpError.ErrorCode().ToString(),
		ErrorMessage: &errorMessage,
	}
}

func getHTTPStatusCode(fpError *errors.FamiPhotoError) int {
	switch fpError.ErrorCode() {
	case errors.InvalidRequestError:
		return http.StatusBadRequest
	case errors.UserAuthorizeError:
		return http.StatusUnauthorized
	case
		errors.FileNotFoundError,
		errors.DBNotFoundError,
		errors.NoExifError:
		return http.StatusNotFound
	case errors.UserIDAlreadyUsedError:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
