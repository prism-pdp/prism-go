package xz21

import (
	_ "fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPairingKey(t * testing.T) {
	// Generate Param
	param1 := GenPairingParam()

	// Generate and save key
	key1 := GenPairingKey(&param1)
	assert.NotNil(t, key1.PrivateKey)
	assert.NotNil(t, key1.PublicKey)

	param1.Save("/tmp/paramX")
	key1.Save("/tmp/key")

	// Load the saved key
	param2 := LoadPairingParamFromFile("/tmp/paramX")
	key2 := LoadPairingKeyFromFile("/tmp/key", &param2)

	// Compare keys
	assert.Equal(t, key1.PrivateKey.Bytes(), key2.PrivateKey.Bytes())
	assert.Equal(t, key1.PublicKey.Bytes(), key2.PublicKey.Bytes())
}