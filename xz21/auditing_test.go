package xz21

import (
	"io/ioutil"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAuditing(t *testing.T) {
	var err error

	// Test data
	data, err := ioutil.ReadFile("testdata/dummy.data")
	assert.NoError(t, err)

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
	var tagDataSet TagDataSet
	{
		// Load param
		param := GenParamFromXZ21Param(&xz21Param)
		// Read key from file
		skSU := skDataSU.Import(&param)
		// Generate tag
		chunkSet, _ := SplitData(data, 8)
		tag, _ := GenTags(&param, skSU.Key, chunkSet)
		// Export data to be sent to SP
		tagDataSet = tag.Export()
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
		chal := NewChal(&param, tagDataSet.Size)
		// Export data to be sent to TPA
		auditingReqData.ChalData = chal.Export()
		tagSize = tagDataSet.Size
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
		tagDataSubset := NewTagDataSet()
		digestSubset, tagDataSubset, err = MakeSubset(data, &tagDataSet, &auditingReq.Chal)
		assert.NoError(t, err)

		// Verify proof
		tagSubset := tagDataSubset.ImportAll(&param)
		result, err = VerifyProof(&param, &tagSubset, digestSubset, &auditingReq.Chal, &auditingReq.Proof, pkSU.Key)
		assert.NoError(t, err)
	}

	assert.True(t, result)
}
