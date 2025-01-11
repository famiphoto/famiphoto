package utils

import (
	"github.com/famiphoto/famiphoto/api/errors"
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
