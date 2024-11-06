package xz21

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/maps"
)

func TestProof(t *testing.T) {
	// Test data
	data := []byte{
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09,
		0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19,
		0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29,
		0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39,
		0x40, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49,
		0x50, 0x51, 0x52, 0x53,
	}
	chunkNum := uint32(6)

	param := GenPairingParam()
	chal, err := NewChal(param, chunkNum, 0.6)
	assert.NoError(t, err)

	proof, subsetDigest, subsetChunk := GenProof(param, chal, chunkNum, data)
	assert.Equal(t, len(subsetDigest), 4)
	assert.Equal(t, len(subsetChunk), 4)

	proofData1 := proof.Export()
	proofData2 := LoadProofData(proofData1.Base())

	assert.Equal(t, proofData1.Base(), proofData2.Base())

	fmt.Printf("chal.C: ")
	fmt.Println(chal.C)
	fmt.Printf("chal.K1: ")
	fmt.Println(chal.K1.BigInt())
	fmt.Printf("Block list: ")
	fmt.Println(maps.Keys(subsetChunk))
}