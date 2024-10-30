package xz21

import (
	"crypto/sha256"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestXZ21Param(t* testing.T) {
	param1 := GenPairingParam()
	xz21Param := param1.ToXZ21Param()
	param2 := GenParamFromXZ21Param(xz21Param)

	assert.Equal(t, param1.P.String(), param2.P.String())
	assert.Equal(t, param1.G.Bytes(), param2.G.Bytes())
	assert.Equal(t, param1.U.Bytes(), param2.U.Bytes())
}

func TestSign(t *testing.T) {
	param := GenPairingParam()
	pk, sk := GenPairingKey(param)

	data := []byte("Hello World")
	hash := sha256.Sum256(data)

	h := param.Pairing.NewG1().SetFromHash(hash[:])
	sig := param.Pairing.NewG1().PowZn(h, sk.Base())

	lhs := param.Pairing.NewGT().Pair(sig, param.G)
	rhs := param.Pairing.NewGT().Pair(h, pk.Base())
	assert.True(t, lhs.Equals(rhs))

	pkDummy, _ := GenPairingKey(param)
	rhsDummy := param.Pairing.NewGT().Pair(h, pkDummy.Base())
	assert.False(t, lhs.Equals(rhsDummy))
}