package runtime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	result := add(1, 1)
	assert.Equal(t, int64(2), result)

	result = add(1, 1.2)
	assert.Equal(t, float64(2.2), result)
}
