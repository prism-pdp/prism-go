package xz21

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPdp(t *testing.T) {
	data := []byte{
		 0,  1,  2,  3,  4,  5,  6,  7,  8,  9,
		10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
		20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
		30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
		40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
		50, 51, 52, 53,
	}

	var param PairingParam
	param.Gen()
	var key PairingKey
	key.Gen(&param)

	chunk, _ := SplitData(data, 10)
	tags, hashChunks, numTags := GenTags(&param, key.PrivateKey, chunk)

	chal := GenChal(&param, numTags)
	proof := GenProof(&param, chal, chunk)
	result := VerifyProof(&param, tags, hashChunks, chal, proof, key.PublicKey)

	assert.True(t, result)
}