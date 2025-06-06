package entities

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStorageFileInfoList_FilterSameNameFiles(t *testing.T) {
	// Common test data
	targetFile := &StorageFileInfo{
		Name:  "sample.jpg",
		Path:  "/path/to/sample.jpg",
		Ext:   ".jpg",
		IsDir: false,
	}

	t.Run("Normal case with all extensions", func(t *testing.T) {
		fileList := StorageFileInfoList{
			// Same name, different extension
			&StorageFileInfo{
				Name:  "sample.png",
				Path:  "/path/to/sample.png",
				Ext:   ".png",
				IsDir: false,
			},
			// Same name, same extension
			&StorageFileInfo{
				Name:  "sample.jpg",
				Path:  "/path/to/sample.jpg",
				Ext:   ".jpg",
				IsDir: false,
			},
			// Different name
			&StorageFileInfo{
				Name:  "other.jpg",
				Path:  "/path/to/other.jpg",
				Ext:   ".jpg",
				IsDir: false,
			},
			// Directory with same name
			&StorageFileInfo{
				Name:  "sample",
				Path:  "/path/to/sample",
				Ext:   "",
				IsDir: true,
			},
		}

		// Execute the method with all extensions
		allExtensions := []string{".jpg", ".png", ".gif"}
		result := fileList.FilterSameNameFiles(targetFile, allExtensions)

		// Verify the results
		assert.Equal(t, 2, len(result), "Should return 2 files with the same name")

		// Check that the result contains only files with the same name (excluding extension)
		for _, file := range result {
			assert.Equal(t, targetFile.FilePathExceptExtHash(), file.FilePathExceptExtHash(), 
				"Files should have the same path hash (excluding extension)")
			assert.False(t, file.IsDir, "Result should not contain directories")
		}
	})

	t.Run("Filter by specific extension", func(t *testing.T) {
		fileList := StorageFileInfoList{
			// Same name, different extension
			&StorageFileInfo{
				Name:  "sample.png",
				Path:  "/path/to/sample.png",
				Ext:   ".png",
				IsDir: false,
			},
			// Same name, same extension
			&StorageFileInfo{
				Name:  "sample.jpg",
				Path:  "/path/to/sample.jpg",
				Ext:   ".jpg",
				IsDir: false,
			},
			// Different name
			&StorageFileInfo{
				Name:  "other.jpg",
				Path:  "/path/to/other.jpg",
				Ext:   ".jpg",
				IsDir: false,
			},
		}

		// Execute the method with only jpg extension
		jpgOnly := []string{".jpg"}
		result := fileList.FilterSameNameFiles(targetFile, jpgOnly)

		// Verify the results
		assert.Equal(t, 1, len(result), "Should return 1 file with the same name and jpg extension")
		assert.Equal(t, ".jpg", result[0].Ext, "Should only return jpg files")
	})

	t.Run("Empty list", func(t *testing.T) {
		fileList := StorageFileInfoList{}

		// Execute the method
		allExtensions := []string{".jpg", ".png"}
		result := fileList.FilterSameNameFiles(targetFile, allExtensions)

		// Verify the results
		assert.Equal(t, 0, len(result), "Should return empty list when input list is empty")
	})

	t.Run("No matches found", func(t *testing.T) {
		fileList := StorageFileInfoList{
			// Different name
			&StorageFileInfo{
				Name:  "other1.jpg",
				Path:  "/path/to/other1.jpg",
				Ext:   ".jpg",
				IsDir: false,
			},
			// Different name
			&StorageFileInfo{
				Name:  "other2.png",
				Path:  "/path/to/other2.png",
				Ext:   ".png",
				IsDir: false,
			},
			// Directory
			&StorageFileInfo{
				Name:  "dir",
				Path:  "/path/to/dir",
				Ext:   "",
				IsDir: true,
			},
		}

		// Execute the method
		allExtensions := []string{".jpg", ".png"}
		result := fileList.FilterSameNameFiles(targetFile, allExtensions)

		// Verify the results
		assert.Equal(t, 0, len(result), "Should return empty list when no matches found")
	})

	t.Run("No matching extensions", func(t *testing.T) {
		fileList := StorageFileInfoList{
			// Same name, different extension
			&StorageFileInfo{
				Name:  "sample.png",
				Path:  "/path/to/sample.png",
				Ext:   ".png",
				IsDir: false,
			},
			// Same name, same extension
			&StorageFileInfo{
				Name:  "sample.jpg",
				Path:  "/path/to/sample.jpg",
				Ext:   ".jpg",
				IsDir: false,
			},
		}

		// Execute the method with non-matching extensions
		gifOnly := []string{".gif"}
		result := fileList.FilterSameNameFiles(targetFile, gifOnly)

		// Verify the results
		assert.Equal(t, 0, len(result), "Should return empty list when no files match the specified extensions")
	})
}
