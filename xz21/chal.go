package xz21

import (
	"math/rand"

	"github.com/Nik-U/pbc" // v0.0.0-20181205041846-3e516ca0c5d6
)

type Chal struct {
	C  uint32
	K1 *pbc.Element
	K2 *pbc.Element
}

func (this *Chal) Gen(_param *PairingParam, _chunkNum uint32) {
	this.C = (rand.Uint32() % _chunkNum)
	this.K1 = _param.Pairing.NewZr().Rand()
	this.K2 = _param.Pairing.NewZr().Rand()
}