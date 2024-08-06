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

	// TPA
	var xz21Para XZ21Para
	{
		param := GenPairingParam()
		xz21Para = param.ToXZ21Para()
	}

	// SU
	var keyData PairingKeyData
	{
		// Load param from Ethereum
		param := GenParamFromXZ21Para(&xz21Para)
		// Generate key pair
		key := GenPairingKey(&param)
		keyData = key.Export()
	}

	// SU
	var chunk [][]byte
	var hashChunks [][]byte
	var tagsData TagsData
	var numTags uint32
	{
		// Load param from Ethereum
		param := GenParamFromXZ21Para(&xz21Para)
		// Read key from file
		keySU := keyData.Import(&param)
		// Generate tags
		chunk, _ = SplitData(data, 5)
		tags, hashChunksTmp, numTagsTmp := GenTags(&param, keySU.PrivateKey, chunk)
		// Export data to be sent to SP
		numTags = numTagsTmp
		hashChunks = hashChunksTmp
		tagsData = tags.Export()
		keyData = keySU.Export()
	}

	// SP
	var chalData ChalData
	{
		// Load param from Ethereum
		param := GenParamFromXZ21Para(&xz21Para)
		// Generate challenge for deduplication
		chal := GenChal(&param, numTags)
		// Export data to be sent to SU
		chalData = chal.Export()
	}

	// SU
	var proofData ProofData
	{
		// Load param from Ethereum
		param := GenParamFromXZ21Para(&xz21Para)
		// Receive chalData from SP
		chal := chalData.Import(&param)
		// Generate proof
		proofSU := GenProof(&param, &chal, chunk)
		// Export data to be sent to SP
		proofData = proofSU.Export()
	}

	// SP
	var result bool
	{
		// Load param from Ethereum
		param := GenParamFromXZ21Para(&xz21Para)
		// Read data from file
		keySU := keyData.Import(&param)
		chal := chalData.Import(&param)
		tags := tagsData.Import(&param)
		// Receive data from SU
		proof := proofData.Import(&param)
		// Verify proof
		result = VerifyProof(&param, &tags, hashChunks, &chal, &proof, keySU.PublicKey)
	}

	assert.True(t, result)
}

