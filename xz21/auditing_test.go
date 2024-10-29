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
	var chunkNum uint32 = 8
	{
		// Load param
		param := GenParamFromXZ21Param(&xz21Param)
		// Read key from file
		skSU := skDataSU.Import(&param)
		// Generate tag
		chunkSet := GenChunkSet(data, chunkNum)
		tag, _ := GenTags(&param, skSU.Key, chunkSet)
		// Export data to be sent to SP
		tagDataSet = tag.Export()
	}

	// =================================
	// Auditing Phase
	// =================================
	var auditingReqData AuditingReqData
	// SU generates challenge for deduplication.
	{
		// Load param
		param := GenParamFromXZ21Param(&xz21Param)
		// Generate challenge for deduplication
		chal := NewChal(&param, tagDataSet.Size)
		// Export data to be sent to TPA
		auditingReqData.ChalData = chal.Export()
	}
	// SP generates proof with data owned by itself.
	var digestSubset *DigestSet
	{
		var proof Proof

		// Load param
		param := GenParamFromXZ21Param(&xz21Param)
		// Receive chalData from SU
		chal := auditingReqData.ImportChal(&param)
		// Generate proof
		digestSubset, proof = GenProof(&param, &chal, chunkNum, data)
		// Export data to be sent to TPA
		auditingReqData.ProofData = proof.Export()
	}
	// TPA verifies proof with tag and public key of SU.
	var result bool
	{
		// Load param
		param := GenParamFromXZ21Param(&xz21Param)

		// Read data from public
		pkSU := pkDataSU.Import(&param)
		auditingReq := auditingReqData.Import(&param)

		// Receive data from SU
		tagDataSubset := tagDataSet.DuplicateSubset(chunkNum, &auditingReq.Chal)
		assert.NoError(t, err)

		// Verify proof
		tagSubset := tagDataSubset.ImportAll(&param) // TODO: Rename ImportAll
		result, err = auditingReq.VerifyProof(&param, chunkNum, &tagSubset, digestSubset, pkSU.Key)
		assert.NoError(t, err)
	}

	assert.True(t, result)
}
