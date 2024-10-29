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

	chunkSet := GenChunkSet(data, 5)
	assert.NoError(t, err)
	assert.Equal(t, uint32(chunkSet.Size()), uint32(5))

	tag, digestSet := GenTags(param, sk.Key, chunkSet)
	assert.Equal(t, tag.Size, uint32(5))
	assert.Equal(t, digestSet.Size(), uint32(5))
	assert.Equal(t, len(tag.Tag), 5)
}