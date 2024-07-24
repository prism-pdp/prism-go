package xz21

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestChal(t *testing.T) {
	var param PairingParam
	param.Gen()

	var chal Chal
	chal.Gen(&param, 10)
	assert.LessOrEqual(t, chal.C, 10)
}