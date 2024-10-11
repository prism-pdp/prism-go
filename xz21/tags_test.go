package xz21

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGenTag(t *testing.T) {
	data := []byte{
		 0,  1,  2,  3,  4,  5,  6,  7,  8,  9,
		10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
		20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
		30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
		40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
		50, 51, 52, 53,
	}

	param := GenPairingParam()
	_, sk := GenPairingKey(&param)

	chunk, err := SplitData(data, 5)
	assert.NoError(t, err)
	assert.Equal(t, len(chunk), 5)

	tag, digestSet := GenTag(&param, sk.Key, chunk)
	assert.Equal(t, tag.Size, uint32(5))
	assert.Equal(t, digestSet.Size(), uint32(5))
	assert.Equal(t, len(tag.G), 5)
}