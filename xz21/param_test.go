package xz21

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPairingParam(t *testing.T) {
	// Generate and save parameters
	param1 := GenPairingParam()
	assert.NotNil(t, param1.Params)
	assert.NotNil(t, param1.Pairing)
	assert.NotNil(t, param1.G)
	assert.NotNil(t, param1.U)
	param1.Save("/tmp/param")

	// Load the saved parameters
	param2 := LoadPairingParamFromFile("/tmp/param")

	// Compare parameters
	assert.Equal(t, param1.Params.String(), param2.Params.String())
	assert.Equal(t, param1.G.Bytes(), param2.G.Bytes())
	assert.Equal(t, param1.U.Bytes(), param2.U.Bytes())
}

func TestXZ21Para(t* testing.T) {
	param1 := GenPairingParam()

	xz21Para := param1.ToXZ21Para()

	param2 := GenParamFromXZ21Para(&xz21Para)

	assert.Equal(t, param1.Params.String(), param2.Params.String())
	assert.Equal(t, param1.G.Bytes(), param2.G.Bytes())
	assert.Equal(t, param1.U.Bytes(), param2.U.Bytes())
}

func TestTwoParams(t *testing.T) {
	param1 := GenPairingParam()

	tmp := param1.ToXZ21Para()
	param2 := GenParamFromXZ21Para(&tmp)

	e1 := param1.SetFromHash([]byte("hello"))
	e2 := param2.SetFromHash([]byte("world"))
	e2Bytes := e2.Bytes()
	e3 := param1.SetBytes(e2Bytes)

	e4 := param1.Mul(e1, e3)

	assert.NotNil(t, e4)
}