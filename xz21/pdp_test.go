package xz21

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPdp(t *testing.T) {
	// Test data
	data := []byte{
		 0,  1,  2,  3,  4,  5,  6,  7,  8,  9,
		10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
		20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
		30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
		40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
		50, 51, 52, 53,
	}

	// =================================
	// Setup Phase
	// =================================
	// TPA generates param for pairing.
	var xz21Para XZ21Para
	{
		param := GenPairingParam()
		xz21Para = param.ToXZ21Para()
	}

	// SU1 generates its own key.
	var pkDataSU1 PublicKeyData
	var skDataSU1 PrivateKeyData
	{
		// Load param from Ethereum
		param := GenParamFromXZ21Para(&xz21Para)
		// Generate key pair
		pk, sk := GenPairingKey(&param)
		pkDataSU1 = pk.Export()
		skDataSU1 = sk.Export()
	}

	// =================================
	// Upload Phase (New File)
	// =================================
	// SU1 generates tags of data to be uploaded
	var tagsData TagsData
	{
		// Load param from Ethereum
		param := GenParamFromXZ21Para(&xz21Para)
		// Read key from file
		skSU1 := skDataSU1.Import(&param)
		// Generate tags
		chunk, _ := SplitData(data, 5)
		tags, _, _ := GenTags(&param, skSU1.Key, chunk)
		// Export data to be sent to SP
		tagsData = tags.Export()
	}

	// =================================
	// Upload Phase (Dedup File)
	// =================================
	// SU2 requests SP to upload data.

	// SP generates challenge for deduplication.
	var chalData ChalData
	var numTags uint32
	{
		// Load param from Ethereum
		param := GenParamFromXZ21Para(&xz21Para)
		// Generate challenge for deduplication
		numTagsTmp := uint32(len(tagsData.Tags))
		chal := GenChal(&param, numTagsTmp)
		// Export data to be sent to SU
		chalData = chal.Export()
		numTags = numTagsTmp
	}

	// SU2 generates proof with data owned by itself.
	var proofData ProofData
	{
		// Load param from Ethereum
		param := GenParamFromXZ21Para(&xz21Para)
		// Receive chalData from SP
		chal := chalData.Import(&param)
		// Generate proof
		chunk, _ := SplitData(data, numTags)
		proof := GenProof(&param, &chal, chunk)
		// Export data to be sent to SP
		proofData = proof.Export()
	}

	// SP verifies proof with tags and public key of SU1.
	var result bool
	{
		// Load param from Ethereum
		param := GenParamFromXZ21Para(&xz21Para)
		// Read data from file
		pkSU1 := pkDataSU1.Import(&param)
		chal := chalData.Import(&param)
		tags := tagsData.Import(&param)
		// Receive data from SU
		proof := proofData.Import(&param)
		// Verify proof
		chunk, _ := SplitData(data, numTags)
		hashChunks := HashChunks(chunk)
		result = VerifyProof(&param, &tags, hashChunks, &chal, &proof, pkSU1.Key)
	}

	assert.True(t, result)
}

