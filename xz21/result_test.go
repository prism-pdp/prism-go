package xz21

import (
	"fmt"
	"math/big"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	data := []byte{
		 0,  1,  2,  3,  4,  5,  6,  7,  8,  9,
		10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
		20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
		30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
		40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
		50, 51, 52, 53,
	}
	splitedData, err := SplitData(data, 10)
	assert.NoError(t, err)

	var param PairingParam
	param.Gen()
	var key PairingKey
	key.Gen(&param)

	metadata := GenMetadata(&param, &key, splitedData)
	chunkNum := uint32(len(metadata.Chunk))

	var chal Chal
	chal.Gen(&param, chunkNum)

	var proof Proof
	proof.Gen(&param, &chal, metadata)
	assert.Equal(t, len(proof.A), int(chal.C))
	assert.Equal(t, len(proof.V), int(chal.C))

	var result Result
	isEqual := result.Verify(&param, metadata, &chal, &proof, &key)

	assert.True(t, isEqual)
}

func Test2(t *testing.T) {
	var param PairingParam
	param.Gen()

	sk := param.Pairing.NewZr().Rand()

	a := param.Pairing.NewG1().SetFromHash([]byte{1,1,1,1})
	b := param.Pairing.NewG1().Rand()
	c := param.Pairing.NewG1().SetFromHash([]byte{1,1,1,2})
	g1 := a.ThenMul(b).ThenPowBig(c.X())

	g2 := param.Pairing.NewG1().SetFromHash([]byte{2,2,2,2})

	g1_sk := param.Pairing.NewG1().PowZn(g1, sk)
	g2_sk := param.Pairing.NewG1().PowZn(g2, sk)

	l := param.Pairing.NewGT().Pair(g1_sk, g2)
	r := param.Pairing.NewGT().Pair(g1, g2_sk)
	fmt.Println(l)
	fmt.Println(r)

	assert.Equal(t, l.Bytes(), r.Bytes())
}

func Test3(t *testing.T) {
	tmp1 := new(big.Int).Mul(big.NewInt(6), big.NewInt(8))
	tmp2 := new(big.Int).Mod(tmp1, big.NewInt(5))
	fmt.Println(tmp2)
	assert.NotNil(t, nil)
}