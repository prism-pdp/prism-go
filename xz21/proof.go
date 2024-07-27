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

// https://github.com/es3ku/z22m2azuma/blob/main/sp/src/interfaces/crypt/crypt_test.go#L47
func (this *Proof) Gen(_param *PairingParam, _chal *Chal, _meta *Metadata) {
	this.A = make([]uint32, _chal.C)
	this.V = make([]*pbc.Element, _chal.C)
	// this.Mu = _param.Pairing.NewZr().Set0()

	nBig := big.NewInt(int64(len(_meta.Chunk)))

	for i := uint32(0); i < _chal.C; i++ {
		iBig := new(big.Int).SetUint64(uint64(i + 1))

		this.A[i] = hash1(iBig, _chal.K1, nBig)
		this.V[i] = hash2(iBig, _chal.K2, _param)

		m := _meta.Chunk[this.A[i]]
		mu := _param.Pairing.NewZr().MulBig(this.V[i], _param.SetFromHash(m).X())
		if i == 0 {
			this.Mu = mu
		} else {
			this.Mu = _param.Pairing.NewZr().Add(this.Mu, mu)
		}

		// this.Mu = this.Mu.Add(this.Mu, _param.Pairing.NewZr().MulBig(this.V[i], _param.SetFromHash(_meta.Chunk[this.A[i]]).X()))
		// m := _param.SetFromHash(_meta.Chunk[this.A[i]]).X()
		// vm := _param.Pairing.NewZr().MulBig(this.V[i], m)
		// this.Mu = this.Mu.ThenAdd(vm)
	}
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