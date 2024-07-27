package xz21

import (
	// "math/big"
	"fmt"

	"github.com/Nik-U/pbc" // v0.0.0-20181205041846-3e516ca0c5d6
)

type Result struct {
	A []uint32
	V []*pbc.Element
	Mu *pbc.Element
}

// https://github.com/es3ku/z22m2azuma/blob/main/sp/src/interfaces/crypt/crypt.go#L98
func (this *Result) Verify(_param *PairingParam, _meta *Metadata, _chal *Chal, _proof *Proof, _key *PairingKey) bool {

	tag := _param.Pairing.NewG1().Set1()
	m   := _param.Pairing.NewG1().Set1()

	for i := uint32(0); i < _chal.C; i++ {
		a := _proof.A[i]
		v := _proof.V[i]

		// Left
		t1 := _meta.Tag[a]
		t2 := _param.PowZn(t1, v)
		tag = _param.Mul(tag, t2)

		// Right
		m1 := _param.SetFromHash(_meta.Hash[a])
		m2 := _param.PowZn(m1, v)
		m   = _param.Mul(m, m2)
	}

	u1  := _param.PowZn(_param.U, _proof.Mu)
	right_hand := _param.Mul(m, u1)

	lhs := _param.Pairing.NewGT().Pair(tag, _param.G)
	rhs := _param.Pairing.NewGT().Pair(right_hand, _key.PublicKey)
	fmt.Println(lhs.Bytes())
	fmt.Println(rhs.Bytes())

	// return lhs.Bytes() rhs.Bytes()
	return lhs.Equals(rhs)

/*
	this.A = make([]uint32, _chal.C)
	this.V = make([]*pbc.Element, _chal.C)
	this.Mu = _param.Pairing.NewZr().SetInt32(0)

	nBig := big.NewInt(int64(len(_meta.Chunk)))

	for i := uint32(0); i < _chal.C; i++ {
		iBig := new(big.Int).SetUint64(uint64(i))
		tmp := new(big.Int)

		this.A[i] = uint32(tmp.Mul(iBig, _chal.K1.BigInt()).Mod(tmp, nBig).Uint64())
		this.V[i] = _param.Pairing.NewZr().SetBig(new(big.Int).Mul(iBig, _chal.K2.BigInt()))

		this.Mu = this.Mu.Add(this.Mu, _param.Pairing.NewZr().MulBig(this.V[i], new(big.Int).SetBytes(_meta.Chunk[this.A[i]])))
	}
*/
}