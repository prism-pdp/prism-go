package xz21

import (
	"io/ioutil"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGenTags(t *testing.T) {
	// Test data
	data, err := ioutil.ReadFile("testdata/dummy.data")
	assert.NoError(t, err)

	param := GenPairingParam()
	_, sk := GenPairingKey(param)

	setChunk := GenChunkSet(data, 5)
	assert.NoError(t, err)
	assert.Equal(t, uint32(setChunk.Size()), uint32(5))

	setTag, digestSet := GenTags(param, sk, setChunk)
	assert.Equal(t, digestSet.Size(), uint32(5))
	assert.Equal(t, setTag.Size(), uint32(5))
}