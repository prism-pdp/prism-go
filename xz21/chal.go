package xz21

import (
	"bytes"
	"encoding/gob"
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
	C uint32  `json:'c'`
	K1 []byte `json:'k1'`
	K2 []byte `json:'k2'`
}

type ChalSet struct {
	SetA []uint32
	SetV []*pbc.Element
}

func NewChal(_param *PairingParam, _chunkNum uint32) Chal {
	var chal Chal
	r := rand.Uint32()
	chal.C = (r % _chunkNum) + 1
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

func (this *ChalData) Encode() ([]byte, error) {
	var buf bytes.Buffer
    err := gob.NewEncoder(&buf).Encode(this)
	if err != nil { return []byte{}, err }

    return buf.Bytes(), nil
}

func DecodeToChalData(_b []byte) (ChalData, error) {
	var chalData ChalData
	dec := gob.NewDecoder(bytes.NewReader(_b))
	err := dec.Decode(&chalData)
	if err != nil { return ChalData{}, err }

	return chalData, nil
}

func (this *Chal) GenChalSet(_param *PairingParam, _chunkNum uint32) *ChalSet {
	chalSet := ChalSet{}
	chalSet.SetA = this.GenA(_chunkNum)
	chalSet.SetV = this.GenV(_param)
	return &chalSet
}

func (this *Chal) GenA(_chunkNum uint32) []uint32 {
	setA := make([]uint32, this.C)
	n := big.NewInt(int64(_chunkNum))
	for i := uint32(0); i < this.C; i++ {
		iBig := new(big.Int).SetUint64(uint64(i + 1))
		setA[i] = hash1(iBig, this.K1, n)
	}
	return setA
}

func (this *Chal) GenV(_param *PairingParam) []*pbc.Element {
	setV := make([]*pbc.Element, this.C)
	for i := uint32(0); i < this.C; i++ {
		iBig := new(big.Int).SetUint64(uint64(i + 1))
		setV[i] = hash2(iBig, this.K2, _param)
	}
	return setV
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