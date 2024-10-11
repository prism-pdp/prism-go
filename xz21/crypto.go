package xz21

import (
	"math/big"

	"github.com/Nik-U/pbc" // v0.0.0-20181205041846-3e516ca0c5d6
)

func GenA(_chal *Chal, _chunkSize uint32) []uint32 {
	k1 := _chal.K1
	c := _chal.C
	a := make([]uint32, c)
	n := big.NewInt(int64(_chunkSize))
	for i := uint32(0); i < c; i++ {
		iBig := new(big.Int).SetUint64(uint64(i + 1))
		a[i] = hash1(iBig, k1, n)
	}
	return a
}

func GenV(_chal *Chal, _param *PairingParam) []*pbc.Element {
	k2 := _chal.K2
	c := _chal.C
	v := make([]*pbc.Element, c)
	for i := uint32(0); i < c; i++ {
		iBig := new(big.Int).SetUint64(uint64(i + 1))
		v[i] = hash2(iBig, k2, _param)
	}
	return v
}

func hash1(_i *big.Int, _k1 *pbc.Element, _n *big.Int) uint32 {
	tmp1 := new(big.Int).Mul(_i, _k1.BigInt())
	tmp2 := new(big.Int).Mod(tmp1, _n).Uint64()
	return uint32(tmp2)
}

func hash2(_i *big.Int, _k2 *pbc.Element, _param *PairingParam) *pbc.Element {
	tmp := _param.Pairing.NewZr().SetBig(new(big.Int).Mul(_i, _k2.BigInt()))
	return tmp
}
