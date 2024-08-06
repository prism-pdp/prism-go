package xz21

import (
	"crypto/sha256"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestXZ21Para(t* testing.T) {
	param1 := GenPairingParam()
	xz21Para := param1.ToXZ21Para()
	param2 := GenParamFromXZ21Para(&xz21Para)

	assert.Equal(t, param1.Params.String(), param2.Params.String())
	assert.Equal(t, param1.G.Bytes(), param2.G.Bytes())
	assert.Equal(t, param1.U.Bytes(), param2.U.Bytes())
}

func TestSign(t *testing.T) {
	param := GenPairingParam()
	pk, sk := GenPairingKey(&param)

	data := []byte("Hello World")
	hash := sha256.Sum256(data)

	h := param.Pairing.NewG1().SetFromHash(hash[:])
	sig := param.Pairing.NewG1().PowZn(h, sk.Key)

	lhs := param.Pairing.NewGT().Pair(sig, param.G)
	rhs := param.Pairing.NewGT().Pair(h, pk.Key)
	assert.True(t, lhs.Equals(rhs))

	pkDummy, _ := GenPairingKey(&param)
	rhsDummy := param.Pairing.NewGT().Pair(h, pkDummy.Key)
	assert.False(t, lhs.Equals(rhsDummy))
}