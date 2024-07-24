package xz21

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSplitFile(t *testing.T) {
	data := []byte{1,2,3,4,5,6,7,8,9,10,11,12,13}
	splitedData, err := SplitData(data, 5)
	assert.NoError(t, err)
	assert.Equal(t, len(splitedData), 3)
	assert.Equal(t, len(splitedData[0]), 5)
	assert.Equal(t, len(splitedData[1]), 5)
	assert.Equal(t, len(splitedData[2]), 3)
}

func TestGenMetadata(t *testing.T) {
	data := []byte{
		 0,  1,  2,  3,  4,  5,  6,  7,  8,  9,
		10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
		20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
		30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
		40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
		50, 51, 52, 53,
	}
	splitedData, err := SplitData(data, 10)
	assert.NoError(t, err)

	var param PairingParam
	param.Gen()
	var key PairingKey
	key.Gen(&param)

	metadata := GenMetadata(&param, &key, splitedData)
	assert.Equal(t, len(metadata.Hash), 6)
	assert.Equal(t, len(metadata.Chunk), 6)
	assert.Equal(t, len(metadata.Tag), 6)
}