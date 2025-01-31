package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParseDatetime(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	actual, err := ParseDatetime("2022:07:19 21:53:00", loc)
	assert.NoError(t, err)
	expected := time.Date(2022, 7, 19, 21, 53, 0, 0, loc)
	assert.Equal(t, expected, actual)
}
