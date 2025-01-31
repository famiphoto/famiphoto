package entities

import (
	"path/filepath"
	"strings"
	"time"
)

type Photo struct {
	PhotoID       int64
	Name          string
	DescriptionJa string
	DescriptionEn string
	ImportedAt    time.Time
	Files         PhotoFileList
}

func (e *Photo) HasJpeg() bool {
	return e.Files.FindFileByFileType(e.PhotoID, PhotoFileTypeJPEG) != nil
}

type PhotoFile struct {
	PhotoFileID int64
	PhotoID     int64
	FileHash    string
	File        StorageFileInfo
}

func (f PhotoFile) FileType() PhotoFileType {
	ext := filepath.Ext(f.File.Path)
	switch strings.ToLower(ext) {
	case ".jpeg", ".jpg":
		return PhotoFileTypeJPEG
	case ".arw":
		return PhotoFileTypeRAW
	}
	return PhotoFileTypeUnknown
}

func (f PhotoFile) MimeType() string {
	switch f.FileType() {
	case PhotoFileTypeJPEG:
		return "image/jpeg"
	case PhotoFileTypeRAW:
		return "image/x-dcraw"
	default:
		return "application/octet-stream"
	}
}

type PhotoFileList []*PhotoFile

func (list PhotoFileList) FindFileByFileType(photoID int64, fileType PhotoFileType) *PhotoFile {
	for _, item := range list {
		if item.PhotoID != photoID {
			continue
		}
		if item.FileType() != fileType {
			continue
		}
		return item
	}
	return nil
}

type PhotoFileType string

func (t PhotoFileType) ToString() string {
	return string(t)
}

const (
	PhotoFileTypeJPEG    PhotoFileType = "jpeg"
	PhotoFileTypeRAW     PhotoFileType = "raw"
	PhotoFileTypeUnknown PhotoFileType = "unknown"
)
