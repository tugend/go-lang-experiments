package lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindSingle(t *testing.T) {
	elements := []int{2, 4, 6, 8, 10}

	isNumber := func(needle int) func(element int) bool {
		return func(element int) bool { return element == needle }
	}

	match, err := Find(elements, isNumber(2))

	assert.Nil(t, err)
	assert.Equal(t, 2, match)
}
