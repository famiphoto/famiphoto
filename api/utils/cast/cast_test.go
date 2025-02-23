package cast

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArray(t *testing.T) {
	type typeA struct {
		value int
	}
	type typeB struct {
		value int64
	}

	src := []*typeA{
		{
			value: 3,
		},
		{
			value: 4,
		},
	}

	castFunc := func(a *typeA) *typeB {
		return &typeB{
			value: int64(a.value),
		}
	}

	actual := Array(src, castFunc)
	expected := []*typeB{
		{
			value: 3,
		},
		{
			value: 4,
		},
	}
	assert.Equal(t, expected, actual)
}
