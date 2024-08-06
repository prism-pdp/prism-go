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
	pk1, sk1 := GenPairingKey(&param1)
	assert.NotNil(t, pk1.Key)
	assert.NotNil(t, sk1.Key)
}