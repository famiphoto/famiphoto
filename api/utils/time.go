package utils

import (
	"fmt"
	"github.com/famiphoto/famiphoto/api/errors"
	"strconv"
	"strings"
	"time"
)

func LocalTime(tm time.Time, tz string) (time.Time, error) {
	location, err := time.LoadLocation(tz)
	if err != nil {
		return time.Time{}, errors.New(errors.InvalidTimezoneFatal, err)
	}
	return tm.In(location), nil
}

func MustLocalTime(tm time.Time, tz string) time.Time {
	tm, err := LocalTime(tm, tz)
	if err != nil {
		panic(err)
	}
	return tm
}

func MustLoadLocation(loc string) *time.Location {
	l, err := time.LoadLocation(loc)
	if err != nil {
		panic(err)
	}
	return l
}

// LocationFromOffset +09:00といった時差を示す文字列からタイムゾーンを取得します。
func LocationFromOffset(offsetStr string) (*time.Location, error) {
	sign := 1
	if strings.HasPrefix(offsetStr, "-") {
		sign = -1
		offsetStr = offsetStr[1:]
	} else if strings.HasPrefix(offsetStr, "+") {
		offsetStr = offsetStr[1:]
	}

	parts := strings.Split(offsetStr, ":")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid offset format")
	}

	h, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}
	m, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}

	totalSeconds := sign * (h*3600 + m*60)
	return time.FixedZone(fmt.Sprintf("%+03d:%02d", sign*h, m), totalSeconds), nil
}

// LocationOrDefaultFromOffset +09:00といった時差を示す文字列からタイムゾーンを取得します。取得に失敗したらデフォルト値を返します。
func LocationOrDefaultFromOffset(offsetStr string, defaultLocation *time.Location) *time.Location {
	loc, err := LocationFromOffset(offsetStr)
	if err != nil {
		return defaultLocation
	}
	return loc
}
