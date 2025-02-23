package errors

import (
	native "errors"
	"fmt"
)

type FamiPhotoError struct {
	errorCode FamiPhotoErrorCode
	baseError error
}

func (e *FamiPhotoError) Error() string {
	if e.baseError == nil {
		return e.ErrorCode().ToString()
	}
	return fmt.Sprintf("code: %s, %s", e.errorCode, e.baseError.Error())
}

func (e *FamiPhotoError) ErrorCode() FamiPhotoErrorCode {
	return e.errorCode
}

func New(errCode FamiPhotoErrorCode, baseError error) error {
	return &FamiPhotoError{
		errorCode: errCode,
		baseError: baseError,
	}
}

func UnwrapFPError(err error) *FamiPhotoError {
	var dst *FamiPhotoError
	if ok := native.As(err, &dst); ok {
		return dst
	}
	return nil
}

func GetFPErrorCode(err error) FamiPhotoErrorCode {
	appError := UnwrapFPError(err)
	if appError == nil {
		return Unknown
	}
	return appError.ErrorCode()
}

func Is(err, target error) bool {
	return native.Is(err, target)
}

func IsErrCode(err error, errCode FamiPhotoErrorCode) bool {
	if err == nil {
		return false
	}
	return GetFPErrorCode(err) == errCode
}
