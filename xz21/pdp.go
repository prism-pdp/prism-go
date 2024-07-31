package xz21

import (
	"math/big"
	"math/rand"

	"github.com/Nik-U/pbc" // v0.0.0-20181205041846-3e516ca0c5d6
)

type Chal struct {
	C  uint32
	K1 *pbc.Element
	K2 *pbc.Element
}

func GenChal(_param *PairingParam, _chunkNum uint32) *Chal {
	chal := new(Chal)
	chal.C = (rand.Uint32() % _chunkNum)
	chal.K1 = _param.Pairing.NewZr().Rand()
	chal.K2 = _param.Pairing.NewZr().Rand()
	return chal
}

// https://github.com/es3ku/z22m2azuma/blob/main/sp/src/interfaces/crypt/crypt_test.go#L47
func GenProof(_param *PairingParam, _chal *Chal, _chunk [][]byte) *pbc.Element {

	a := make([]uint32, _chal.C)
	v := make([]*pbc.Element, _chal.C)
	n := big.NewInt(int64(len(_chunk)))

	proof := _param.Pairing.NewZr().Set0()
	for i := uint32(0); i < _chal.C; i++ {
		iBig := new(big.Int).SetUint64(uint64(i + 1))

		a[i] = hash1(iBig, _chal.K1, n)
		v[i] = hash2(iBig, _chal.K2, _param)

		m := _chunk[a[i]]
		mu := _param.Pairing.NewZr().MulBig(v[i], _param.SetFromHash(m).X())
		proof = _param.Pairing.NewZr().Add(proof, mu)
	}

	return proof
}

// https://github.com/es3ku/z22m2azuma/blob/main/sp/src/interfaces/crypt/crypt.go#L98
func VerifyProof(_param *PairingParam, _meta *Metadata, _chal *Chal, _proof *pbc.Element, _pubKey *pbc.Element) bool {

	tag := _param.Pairing.NewG1().Set1()
	m   := _param.Pairing.NewG1().Set1()

	setA := GenA(_chal.K1, _chal.C, _meta.Size)
	setV := GenV(_chal.K2, _chal.C, _param)

	for i := uint32(0); i < _chal.C; i++ {
		a := setA[i]
		v := setV[i]

		// Left
		t1 := _meta.Tags[a]
		t2 := _param.PowZn(t1, v)
		tag = _param.Mul(tag, t2)

		// Right
		m1 := _param.SetFromHash(_meta.Hash[a])
		m2 := _param.PowZn(m1, v)
		m   = _param.Mul(m, m2)
	}

	u1  := _param.PowZn(_param.U, _proof)
	right_hand := _param.Mul(m, u1)

	lhs := _param.Pairing.NewGT().Pair(tag, _param.G)
	rhs := _param.Pairing.NewGT().Pair(right_hand, _pubKey)

	return lhs.Equals(rhs)
}

func GenA(_k1 *pbc.Element, _c uint32, _chunkSize uint32) []uint32 {
	a := make([]uint32, _c)
	n := big.NewInt(int64(_chunkSize))
	for i := uint32(0); i < _c; i++ {
		iBig := new(big.Int).SetUint64(uint64(i + 1))
		a[i] = hash1(iBig, _k1, n)
	}
	return a
}

func GenV(_k2 *pbc.Element, _c uint32, _param *PairingParam) []*pbc.Element {
	v := make([]*pbc.Element, _c)
	for i := uint32(0); i < _c; i++ {
		iBig := new(big.Int).SetUint64(uint64(i + 1))
		v[i] = hash2(iBig, _k2, _param)
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
