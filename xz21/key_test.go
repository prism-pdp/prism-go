package xz21

import (
	_ "fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPairingKey(t * testing.T) {
	// Generate Param
	var param PairingParam
	param.Gen()

	// Generate and save key
	var key1 PairingKey
	key1.Gen(&param)
	assert.NotNil(t, key1.PrivateKey)
	assert.NotNil(t, key1.PublicKey)
	key1.Save("/tmp/key")

	// Load the saved key
	key2 := LoadPairingKeyFromFile("/tmp/key", &param)

	// Compare keys
	assert.Equal(t, key1.PrivateKey.Bytes(), key2.PrivateKey.Bytes())
	assert.Equal(t, key1.PublicKey.Bytes(), key2.PublicKey.Bytes())
}