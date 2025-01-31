package errors

type FamiPhotoErrorCode string

func (c FamiPhotoErrorCode) ToString() string {
	return string(c)
}

const (
	Unknown              FamiPhotoErrorCode = "Unknown"
	InvalidRequestError  FamiPhotoErrorCode = "InvalidRequestError"
	InvalidTimezoneFatal FamiPhotoErrorCode = "InvalidTimezoneFatal"
	NoExifError          FamiPhotoErrorCode = "NoExifError"
	FileNotFoundError    FamiPhotoErrorCode = "FileNotFoundError"
	DBNotFoundError      FamiPhotoErrorCode = "DBNotFoundError"
)
