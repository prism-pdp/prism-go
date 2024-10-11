package xz21

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAuditing(t *testing.T) {
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

	// SU generates its own key.
	var pkDataSU PublicKeyData
	var skDataSU PrivateKeyData
	{
		// Load param
		param := GenParamFromXZ21Param(&xz21Param)
		// Generate key pair
		pk, sk := GenPairingKey(&param)
		pkDataSU = pk.Export()
		skDataSU = sk.Export()
	}

	// =================================
	// Upload Phase (New File)
	// =================================
	// SU generates tag of data to be uploaded
	var tagData TagData
	{
		// Load param
		param := GenParamFromXZ21Param(&xz21Param)
		// Read key from file
		skSU := skDataSU.Import(&param)
		// Generate tag
		chunkSet, _ := SplitData(data, 8)
		tag, _ := GenTag(&param, skSU.Key, chunkSet)
		// Export data to be sent to SP
		tagData = tag.Export()
	}

	// =================================
	// Auditing Phase
	// =================================
	var auditingReqData AuditingReqData
	// SU generates challenge for deduplication.
	var tagSize uint32
	{
		// Load param
		param := GenParamFromXZ21Param(&xz21Param)
		// Generate challenge for deduplication
		chal := NewChal(&param, tagData.Size)
		// Export data to be sent to TPA
		auditingReqData.ChalData = chal.Export()
		tagSize = tagData.Size
	}
	// SP generates proof with data owned by itself.
	{
		var chal Chal
		var chunkSet *ChunkSet
		var param PairingParam
		var proof Proof

		// Load param
		param = GenParamFromXZ21Param(&xz21Param)
		// Receive chalData from SU
		chal = auditingReqData.ImportChal(&param)
		// Generate proof
		chunkSet, _ = SplitData(data, tagSize)
		proof = GenProof(&param, &chal, chunkSet)
		// Export data to be sent to TPA
		auditingReqData.ProofData = proof.Export()
	}
	// TPA verifies proof with tag and public key of SU.
	var result bool
	{
		var digestSubset *DigestSet

		// Load param
		param := GenParamFromXZ21Param(&xz21Param)

		// Read data from public
		pkSU := pkDataSU.Import(&param)
		auditingReq := auditingReqData.Import(&param)

		// Receive data from SU
		tagDataSubset := NewTagData()
		digestSubset, tagDataSubset, err = MakeSubset(data, &tagData, &auditingReq.Chal)
		assert.NoError(t, err)

		// Verify proof
		tagSubset := tagDataSubset.ImportAll(&param)
		result, err = VerifyProof(&param, &tagSubset, digestSubset, &auditingReq, pkSU.Key)
		assert.NoError(t, err)
	}

	assert.True(t, result)
}
