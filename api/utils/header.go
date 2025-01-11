package utils

import "strings"

func ParseAuthHeader(header, authType string) (string, bool) {
	list := strings.Split(header, " ")
	if len(list) != 2 {
		return "", false
	}
	if strings.ToLower(list[0]) != strings.ToLower(authType) {
		return "", false
	}

	return list[1], true
}
