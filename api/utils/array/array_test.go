package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsContain_含んでいる(t *testing.T) {
	list := []int64{1, 2, 3, 4, 5}
	actual := IsContain(3, list)
	assert.True(t, actual)
}

func TestIsContain_含んでいない(t *testing.T) {
	list := []int64{1, 2, 3, 4, 5}
	actual := IsContain(6, list)
	assert.False(t, actual)
}

func TestIsContain_リストが空(t *testing.T) {
	list := []int64{}
	actual := IsContain(1, list)
	assert.False(t, actual)
}

func TestRemoveDuplicates(t *testing.T) {
	list := []string{"a", "b", "c", "c", "b", "d"}
	actual := RemoveDuplicates(list)
	expected := []string{"a", "b", "c", "d"}
	assert.Equal(t, expected, actual)
}
