package errors

type FamiPhotoErrorCode string

func (c FamiPhotoErrorCode) ToString() string {
	return string(c)
}

const (
	Unknown             FamiPhotoErrorCode = "Unknown"
	InvalidRequestError FamiPhotoErrorCode = "InvalidRequestError"
)
