package xz21

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNow(t *testing.T) {
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
	chunkNum := uint32(len(metadata.Chunk))

	var chal Chal
	chal.Gen(&param, chunkNum)

	var proof Proof
	proof.Gen(&param, &chal, metadata)
	assert.Equal(t, len(proof.A), int(chal.C))
	assert.Equal(t, len(proof.V), int(chal.C))

	assert.NotNil(t, nil)
	fmt.Println(chal.K1.BigInt())
	fmt.Println(proof)
}