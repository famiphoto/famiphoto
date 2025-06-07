package entities

import (
	"github.com/stretchr/testify/assert"
	"strings"
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

func TestStorageFileInfoList_GroupByBaseFileName(t *testing.T) {
	t.Run("Normal case with files having same base names", func(t *testing.T) {
		fileList := StorageFileInfoList{
			// Group 1: sample files
			&StorageFileInfo{
				Name:  "sample.jpg",
				Path:  "/path/to/sample.jpg",
				Ext:   ".jpg",
				IsDir: false,
			},
			&StorageFileInfo{
				Name:  "sample.png",
				Path:  "/path/to/sample.png",
				Ext:   ".png",
				IsDir: false,
			},
			// Group 2: other files
			&StorageFileInfo{
				Name:  "other.jpg",
				Path:  "/path/to/other.jpg",
				Ext:   ".jpg",
				IsDir: false,
			},
			&StorageFileInfo{
				Name:  "other.arw",
				Path:  "/path/to/other.arw",
				Ext:   ".arw",
				IsDir: false,
			},
			// Different file (no group)
			&StorageFileInfo{
				Name:  "unique.jpg",
				Path:  "/path/to/unique.jpg",
				Ext:   ".jpg",
				IsDir: false,
			},
			// Directory (should be ignored)
			&StorageFileInfo{
				Name:  "sample",
				Path:  "/path/to/sample",
				Ext:   "",
				IsDir: true,
			},
		}

		// Execute the method
		result := fileList.GroupByBaseFileName()

		// Verify the results
		assert.Equal(t, 3, len(result), "Should return 3 groups")

		// Check each group
		for _, group := range result {
			// All files in a group should have the same base name
			if len(group) > 0 {
				baseHash := group[0].FilePathExceptExtHash()
				for _, file := range group {
					assert.Equal(t, baseHash, file.FilePathExceptExtHash(), 
						"Files in the same group should have the same base name hash")
					assert.False(t, file.IsDir, "Groups should not contain directories")
				}
			}
		}

		// Verify specific groups
		// Find the group with "sample" files
		var sampleGroup StorageFileInfoList
		var otherGroup StorageFileInfoList
		var uniqueGroup StorageFileInfoList

		for _, group := range result {
			if len(group) > 0 {
				if strings.Contains(group[0].Path, "sample") {
					sampleGroup = group
				} else if strings.Contains(group[0].Path, "other") {
					otherGroup = group
				} else if strings.Contains(group[0].Path, "unique") {
					uniqueGroup = group
				}
			}
		}

		assert.Equal(t, 2, len(sampleGroup), "Sample group should have 2 files")
		assert.Equal(t, 2, len(otherGroup), "Other group should have 2 files")
		assert.Equal(t, 1, len(uniqueGroup), "Unique group should have 1 file")
	})

	t.Run("Empty list", func(t *testing.T) {
		fileList := StorageFileInfoList{}

		// Execute the method
		result := fileList.GroupByBaseFileName()

		// Verify the results
		assert.Equal(t, 0, len(result), "Should return empty list when input list is empty")
	})

	t.Run("List with only directories", func(t *testing.T) {
		fileList := StorageFileInfoList{
			&StorageFileInfo{
				Name:  "dir1",
				Path:  "/path/to/dir1",
				Ext:   "",
				IsDir: true,
			},
			&StorageFileInfo{
				Name:  "dir2",
				Path:  "/path/to/dir2",
				Ext:   "",
				IsDir: true,
			},
		}

		// Execute the method
		result := fileList.GroupByBaseFileName()

		// Verify the results
		assert.Equal(t, 0, len(result), "Should return empty list when input contains only directories")
	})

	t.Run("List with files having different base names", func(t *testing.T) {
		fileList := StorageFileInfoList{
			&StorageFileInfo{
				Name:  "file1.jpg",
				Path:  "/path/to/file1.jpg",
				Ext:   ".jpg",
				IsDir: false,
			},
			&StorageFileInfo{
				Name:  "file2.jpg",
				Path:  "/path/to/file2.jpg",
				Ext:   ".jpg",
				IsDir: false,
			},
			&StorageFileInfo{
				Name:  "file3.jpg",
				Path:  "/path/to/file3.jpg",
				Ext:   ".jpg",
				IsDir: false,
			},
		}

		// Execute the method
		result := fileList.GroupByBaseFileName()

		// Verify the results
		assert.Equal(t, 3, len(result), "Should return 3 groups for 3 different files")
		for _, group := range result {
			assert.Equal(t, 1, len(group), "Each group should contain exactly 1 file")
		}
	})

	t.Run("List with mix of directories and files", func(t *testing.T) {
		fileList := StorageFileInfoList{
			// Files with same base name
			&StorageFileInfo{
				Name:  "sample.jpg",
				Path:  "/path/to/sample.jpg",
				Ext:   ".jpg",
				IsDir: false,
			},
			&StorageFileInfo{
				Name:  "sample.png",
				Path:  "/path/to/sample.png",
				Ext:   ".png",
				IsDir: false,
			},
			// Directories (should be ignored)
			&StorageFileInfo{
				Name:  "dir1",
				Path:  "/path/to/dir1",
				Ext:   "",
				IsDir: true,
			},
			&StorageFileInfo{
				Name:  "dir2",
				Path:  "/path/to/dir2",
				Ext:   "",
				IsDir: true,
			},
		}

		// Execute the method
		result := fileList.GroupByBaseFileName()

		// Verify the results
		assert.Equal(t, 1, len(result), "Should return 1 group for files with same base name")
		assert.Equal(t, 2, len(result[0]), "The group should contain 2 files")

		// Check that all files in the group have the same base name
		baseHash := result[0][0].FilePathExceptExtHash()
		for _, file := range result[0] {
			assert.Equal(t, baseHash, file.FilePathExceptExtHash(), 
				"Files in the same group should have the same base name hash")
			assert.False(t, file.IsDir, "Groups should not contain directories")
		}
	})
}
