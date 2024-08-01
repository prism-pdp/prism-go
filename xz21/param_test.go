package xz21

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPairingParam(t *testing.T) {
	// Generate and save parameters
	var param1 PairingParam
	param1.Gen()
	assert.NotNil(t, param1.Params)
	assert.NotNil(t, param1.Pairing)
	assert.NotNil(t, param1.G)
	param1.Save("/tmp/param")

	// Load the saved parameters
	var param2 PairingParam
	param2.Load("/tmp/param")

	// Compare parameters
	assert.Equal(t, param1.Params.String(), param2.Params.String())
	assert.Equal(t, param1.G.Bytes(), param2.G.Bytes())
}

func TestXZ21Para(t* testing.T) {
	var param1 PairingParam
	param1.Gen()

	xz21Para := param1.ToXZ21Para()

	param2 := GenParamFromXZ21Para(xz21Para)

	assert.Equal(t, param1.Params.String(), param2.Params.String())
	assert.Equal(t, param1.G.Bytes(), param2.G.Bytes())
	assert.Equal(t, param1.U.Bytes(), param2.U.Bytes())
}