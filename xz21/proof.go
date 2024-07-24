package xz21

import (
	"math/big"

	"github.com/Nik-U/pbc" // v0.0.0-20181205041846-3e516ca0c5d6
)

type Proof struct {
	A []uint32
	V []*pbc.Element
	Mu *pbc.Element
}

func (this *Proof) Gen(_param *PairingParam, _chal *Chal, _meta *Metadata) {
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
}