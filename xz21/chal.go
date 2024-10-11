package xz21

import (
	"bytes"
	"encoding/gob"
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