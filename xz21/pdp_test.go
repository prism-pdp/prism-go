package xz21

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestChal(t *testing.T) {
	chunkNum := uint32(100)

	param := GenPairingParam()
	chal := GenChal(&param, chunkNum)
	assert.LessOrEqual(t, chal.C, chunkNum)
	assert.NotEqual(t, chal.C, uint32(0))

	chalData1 := chal.Export()
	chalBytes, err := chalData1.Encode()
	assert.NoError(t, err)

	chalData2, err := DecodeToChalData(chalBytes)
	assert.NoError(t, err)

	assert.Equal(t, chalData1.C, chalData2.C)
	assert.Equal(t, chalData1.K1, chalData2.K1)
	assert.Equal(t, chalData1.K2, chalData2.K2)
}

func TestProof(t *testing.T) {
	// Test data
	data := [][]byte{
		{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09},
		{0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19},
		{0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29},
		{0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39},
		{0x40, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49},
		{0x50, 0x51, 0x52, 0x53},
	}
	chunkSet := NewChunkSet()
	chunkSet.Fill(data)
	chunkNum := uint32(6)

	param := GenPairingParam()
	chal := GenChal(&param, chunkNum)

	proof := GenProof(&param, &chal, chunkSet)

	proofData1 := proof.Export()
	proofBytes, err := proofData1.Encode()
	assert.NoError(t, err)

	proofData2, err := DecodeToProofData(proofBytes)
	assert.NoError(t, err)

	assert.Equal(t, proofData1.Mu, proofData2.Mu)
}

func TestPdp(t *testing.T) {
	var err error

	// Test data
	data := []byte{
		 0,  1,  2,  3,  4,  5,  6,  7,  8,  9,
		10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
		20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
		30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
		40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
		50, 51, 52, 53, 54, 55, 56, 57, 58, 59,
		60, 61, 62, 63, 64, 65, 66, 67, 68, 69,
		70, 71, 72, 73, 74, 75, 76, 77, 78, 79,
		80, 81, 82, 83, 84, 85, 86, 87, 88, 89,
		90, 91, 92, 93, 94, 95, 96, 97, 98, 99,
	}

	// =================================
	// Setup Phase
	// =================================
	// TPA generates param for pairing.
	var xz21Param XZ21Param
	{
		param := GenPairingParam()
		xz21Param = param.ToXZ21Param()
	}

	// SU1 generates its own key.
	var pkDataSU1 PublicKeyData
	var skDataSU1 PrivateKeyData
	{
		// Load param
		param := GenParamFromXZ21Param(&xz21Param)
		// Generate key pair
		pk, sk := GenPairingKey(&param)
		pkDataSU1 = pk.Export()
		skDataSU1 = sk.Export()
	}

	// =================================
	// Upload Phase (New File)
	// =================================
	// SU1 generates tag of data to be uploaded
	var tagData TagData
	{
		// Load param
		param := GenParamFromXZ21Param(&xz21Param)
		// Read key from file
		skSU1 := skDataSU1.Import(&param)
		// Generate tag
		chunkSet, _ := SplitData(data, 8)
		tag, _ := GenTag(&param, skSU1.Key, chunkSet)
		// Export data to be sent to SP
		tagData = tag.Export()
	}

	// =================================
	// Upload Phase (Dedup File)
	// =================================
	// SP generates challenge for deduplication.
	var chalData ChalData
	var tagSize uint32
	{
		// Load param
		param := GenParamFromXZ21Param(&xz21Param)
		// Generate challenge for deduplication
		chal := GenChal(&param, tagData.Size)
		// Export data to be sent to SU
		chalData = chal.Export()
		tagSize = tagData.Size
	}
	// SU2 generates proof with data owned by itself.
	var proofData ProofData
	{
		var chal Chal
		var chunkSet *ChunkSet
		var param PairingParam
		var proof Proof

		// Load param
		param = GenParamFromXZ21Param(&xz21Param)
		// Receive chalData from SP
		chal = chalData.Import(&param)
		// Generate proof
		chunkSet, _ = SplitData(data, tagSize)
		proof = GenProof(&param, &chal, chunkSet)
		// Export data to be sent to SP
		proofData = proof.Export()
	}
	// SP verifies proof with tag and public key of SU1.
	var result bool
	{
		var digestSubset *DigestSet

		// Load param
		param := GenParamFromXZ21Param(&xz21Param)

		// Read data from file
		pkSU1 := pkDataSU1.Import(&param)
		chal := chalData.Import(&param)

		// Receive data from SU
		proof := proofData.Import(&param)

		// Verify proof
		tagDataSubset := NewTagData()
		digestSubset, tagDataSubset, err = MakeSubset(data, &tagData, &chal)
		assert.NoError(t, err)

		tagSubset := tagDataSubset.ImportAll(&param)
		result, err = VerifyProof(&param, &tagSubset, digestSubset, &chal, &proof, pkSU1.Key)
		assert.NoError(t, err)
	}

	assert.True(t, result)
}

