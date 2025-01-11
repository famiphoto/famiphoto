package utils

import "strings"

func RemoveTrailingSlash(url string) string {
	return strings.TrimSuffix(url, "/")
}
