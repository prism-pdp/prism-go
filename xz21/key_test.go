package xz21

import (
	_ "fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPairingKey(t * testing.T) {
	var param PairingParam
	var key1 PairingKey
	var key2 PairingKey

	param.Gen()
	key1.Gen(&param)
	assert.NotNil(t, key1.PrivateKey)
	assert.NotNil(t, key1.PublicKey)
	key1.Save("/tmp/key")

	key2.Load("/tmp/key", &param)
	assert.Equal(t, key1.PrivateKey.Bytes(), key2.PrivateKey.Bytes())
	assert.Equal(t, key1.PublicKey.Bytes(), key2.PublicKey.Bytes())
}