package xz21

import (
	"bytes"
	"encoding/gob"

	"github.com/Nik-U/pbc" // v0.0.0-20181205041846-3e516ca0c5d6
)

type Proof struct {
	Mu *pbc.Element
}

type ProofData struct {
	Mu []byte `json:'mu'`
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

func (this *ProofData) Encode() ([]byte, error) {
	var buf bytes.Buffer
    err := gob.NewEncoder(&buf).Encode(this)
	if err != nil { return []byte{}, err }

    return buf.Bytes(), nil
}

func DecodeToProofData(_b []byte) (ProofData, error) {
	var proofData ProofData
	dec := gob.NewDecoder(bytes.NewReader(_b))
	err := dec.Decode(&proofData)
	if err != nil { return ProofData{}, err }

	return proofData, nil
}
// https://github.com/es3ku/z22m2azuma/blob/main/sp/src/interfaces/crypt/crypt_test.go#L47
func GenProof(_param *PairingParam, _chal *Chal, _chunkNum uint32, _data []byte) (*DigestSet, Proof) {
	setA := _chal.GenA(_chunkNum)
	setV := _chal.GenV(_param)

	chunkSubset := GenChunkSubsetByIndex(_data, _chunkNum, setA)
	digestSet := chunkSubset.Hash()

	var proof Proof
	proof.Mu = _param.Pairing.NewZr().Set0()
	for i := uint32(0); i < _chal.C; i++ {
		m := digestSet.Get(setA[i])
		mu := _param.Pairing.NewZr().MulBig(setV[i], _param.SetFromHash(m).X())
		proof.Mu = _param.Pairing.NewZr().Add(proof.Mu, mu)
	}

	return digestSet, proof
}

