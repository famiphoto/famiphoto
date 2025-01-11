package utils

import (
	"fmt"
	"github.com/famiphoto/famiphoto/api/utils/array"

	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func SplitPath(path string) []string {
	return array.Filter(strings.Split(path, "/"), func(t string) bool {
		if len(t) == 0 {
			return false
		}
		if t == "." {
			return false
		}
		return true
	})
}

func FileNameExceptExt(p string) string {
	name := filepath.Base(p)
	return strings.TrimSuffix(name, filepath.Ext(name))
}

func IncrementFileNameSuffix(p string) string {
	ext := path.Ext(p)
	name := strings.TrimSuffix(path.Base(p), ext)

	regex := regexp.MustCompile(`(.*)+\-(\d+)`)
	matches := regex.FindStringSubmatch(name)
	if len(matches) < 3 {
		return fmt.Sprintf("%s-%d%s", name, 1, ext)
	}
	i, err := strconv.Atoi(matches[2])
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s-%d%s", matches[1], i+1, ext)
}
