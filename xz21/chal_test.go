package xz21

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestChal(t *testing.T) {
	chunkNum := uint32(50)

	param := GenPairingParam()
	chal, err := NewChal(param, chunkNum, 0.85)
	assert.NoError(t, err)
	assert.Equal(t, chal.GetTargetBlockCount(), uint32(43))

	chalData1 := chal.Export()
	chalBytes, err := chalData1.Encode()
	assert.NoError(t, err)

	chalData2, err := DecodeToChalData(chalBytes)
	assert.NoError(t, err)

	assert.Equal(t, chalData1.C, chalData2.C)
	assert.Equal(t, chalData1.K1, chalData2.K1)
	assert.Equal(t, chalData1.K2, chalData2.K2)
}
