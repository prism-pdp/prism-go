package xz21

import (
	_ "fmt"
	"reflect"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPairingKey(t * testing.T) {
	// Generate Param
	param1 := GenPairingParam()

	// Generate and save key
	pk1, sk1 := GenPairingKey(param1)
	assert.NotNil(t, pk1)
	assert.NotNil(t, sk1)

	pkData1 := pk1.Export()
	pkByte1 := pkData1.Base()

	var pkData2 PublicKeyData
	pkData2.Load(pkByte1)
	pkByte2 := pkData2.Base()

	assert.True(t, reflect.DeepEqual(pkByte1, pkByte2))
}