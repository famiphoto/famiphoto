package entities

import (
	"path/filepath"
	"strings"
)

type PhotoFile struct {
	PhotoFileID string
	PhotoID     string
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

func (list PhotoFileList) FindFileByFileType(photoID string, fileType PhotoFileType) *PhotoFile {
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
