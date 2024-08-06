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

type ChalData struct {
	C uint32
	K1 []byte
	K2 []byte
}

type Proof struct {
	Mu *pbc.Element
}

type ProofData struct {
	Mu []byte
}

func GenChal(_param *PairingParam, _chunkNum uint32) Chal {
	var chal Chal
	chal.C = (rand.Uint32() % _chunkNum) + 1
	chal.K1 = _param.Pairing.NewZr().Rand()
	chal.K2 = _param.Pairing.NewZr().Rand()
	return chal
}

func (this *Chal) Export() ChalData {
	var data ChalData
	data.C = this.C
	data.K1 = this.K1.Bytes()
	data.K2 = this.K2.Bytes()
	return data
}

func (this *ChalData) Import(_param *PairingParam) Chal {
	var obj Chal
	obj.C = this.C
	obj.K1 = _param.Pairing.NewZr().SetBytes(this.K1)
	obj.K2 = _param.Pairing.NewZr().SetBytes(this.K2)
	return obj
}

func (this *Proof) Export() ProofData {
	var data ProofData
	data.Mu = this.Mu.Bytes()
	return data
}

func (this *ProofData) Import(_param *PairingParam) Proof {
	var obj Proof
	obj.Mu = _param.Pairing.NewZr().SetBytes(this.Mu)
	return obj
}

// https://github.com/es3ku/z22m2azuma/blob/main/sp/src/interfaces/crypt/crypt_test.go#L47
func GenProof(_param *PairingParam, _chal *Chal, _chunk [][]byte) Proof {

	a := make([]uint32, _chal.C)
	v := make([]*pbc.Element, _chal.C)
	n := big.NewInt(int64(len(_chunk)))

	var proof Proof
	proof.Mu = _param.Pairing.NewZr().Set0()
	for i := uint32(0); i < _chal.C; i++ {
		iBig := new(big.Int).SetUint64(uint64(i + 1))

		a[i] = hash1(iBig, _chal.K1, n)
		v[i] = hash2(iBig, _chal.K2, _param)

		m := _chunk[a[i]]
		mu := _param.Pairing.NewZr().MulBig(v[i], _param.SetFromHash(m).X())
		proof.Mu = _param.Pairing.NewZr().Add(proof.Mu, mu)
	}

	return proof
}

// https://github.com/es3ku/z22m2azuma/blob/main/sp/src/interfaces/crypt/crypt.go#L98
func VerifyProof(_param *PairingParam, _tags *Tags, _hashChunks [][]byte, _chal *Chal, _proof *Proof, _pubKey *pbc.Element) bool {

	left  := _param.Pairing.NewG1().Set1()
	right := _param.Pairing.NewG1().Set1()
	n := uint32(len(_tags.Tags))

	setA := GenA(_chal.K1, _chal.C, n)
	setV := GenV(_chal.K2, _chal.C, _param)

	for i := uint32(0); i < _chal.C; i++ {
		a := setA[i]
		v := setV[i]

		// Left
		t1 := _tags.Tags[a]
		t2 := _param.PowZn(t1, v)
		left = _param.Mul(left, t2)

		// Right
		m1 := _param.SetFromHash(_hashChunks[a])
		m2 := _param.PowZn(m1, v)
		right  = _param.Mul(right, m2)
	}

	u := _param.PowZn(_param.U, _proof.Mu)
	right = _param.Mul(right, u)

	lhs := _param.Pairing.NewGT().Pair(left, _param.G)
	rhs := _param.Pairing.NewGT().Pair(right, _pubKey)

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
