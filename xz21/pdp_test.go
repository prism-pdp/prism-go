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

func TestX1(t *testing.T) {
	data := []byte{1,2,3,4,5,6,7,8,9,10,11,12,13}
	splitedData, err := SplitData(data, 5)
	assert.NoError(t, err)

	var param PairingParam
	param.Gen()
	var key PairingKey
	key.Gen(&param)

	metadata := GenMetadata(&param, &key, splitedData)
	assert.Equal(t, len(metadata.Hash), 3)
	assert.Equal(t, len(metadata.Chunk), 3)
	assert.Equal(t, len(metadata.Tag), 3)
}