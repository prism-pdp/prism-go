package xz21

import (
	"fmt"
	"github.com/Nik-U/pbc" // v0.0.0-20181205041846-3e516ca0c5d6
)

type Proof pbc.Element
type ProofData []byte

func (this *Proof) Base() *pbc.Element {
	return (*pbc.Element)(this)
}

func (this ProofData) Base() []byte {
	return ([]byte)(this)
}

func (this *Proof) Export() ProofData {
	data := (ProofData)(this.Base().Bytes())
	return data
}

func (this ProofData) Import(_param *PairingParam) *Proof {
	proof := (*Proof)(_param.Pairing.NewZr().SetBytes(this))
	return proof
}

func LoadProofData(_src []byte) ProofData {
	proofData := make([]byte, len(_src))
	copy(proofData, _src)
	return proofData
}

func GenProof(_param *PairingParam, _chal *Chal, _chunkNum uint32, _data []byte) (*Proof, DigestSet, ChunkSet) {
	setA := _chal.GenA(_chunkNum)
	setV := _chal.GenV(_param)

	subsetChunk := GenChunkSubsetByIndex(_data, _chunkNum, setA)
	subsetDigest := subsetChunk.Hash()

	proof := (*Proof)(_param.Pairing.NewZr().Set0())
	for i := uint32(0); i < _chal.C; i++ {
		m := subsetDigest[setA[i]]
		mu := _param.Pairing.NewZr().MulBig(setV[i], _param.SetFromHash(m).X())
		proof = (*Proof)(_param.Pairing.NewZr().Add(proof.Base(), mu))
	}

	return proof, subsetDigest, subsetChunk
}
