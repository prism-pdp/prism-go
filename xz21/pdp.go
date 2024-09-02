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

type Proof struct {
	Mu *pbc.Element
}

type ProofData struct {
	Mu []byte `json:'mu'`
}

type AuditingReq struct {
	ChalData ChalData `json:'chal'`
	ProofData ProofData `json:'proof'`
}

type AuditingLog struct {
	ChalData ChalData `json:'chal'`
	ProofData ProofData `json:'proof'`
	Result bool `json:'result'`
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

func (this *AuditingReq) Import(_src *XZ21AuditingReq) error {
	var err error
	this.ChalData, err = DecodeToChalData(_src.Chal)
	if err != nil { return err }
	this.ProofData, err = DecodeToProofData(_src.Proof)
	if err != nil { return err }

	return nil
}

func (this *AuditingLog) Import(_src *XZ21AuditingLog) error {
	var err error
	this.ChalData, err = DecodeToChalData(_src.Chal)
	if err != nil { return err }
	this.ProofData, err = DecodeToProofData(_src.Proof)
	if err != nil { return err }
	this.Result = _src.Result

	return nil
}

func DecodeToProofData(_b []byte) (ProofData, error) {
	var proofData ProofData
	dec := gob.NewDecoder(bytes.NewReader(_b))
	err := dec.Decode(&proofData)
	if err != nil { return ProofData{}, err }

	return proofData, nil
}
// https://github.com/es3ku/z22m2azuma/blob/main/sp/src/interfaces/crypt/crypt_test.go#L47
func GenProof(_param *PairingParam, _chal *Chal, _chunk [][]byte) Proof {

	n := uint32(len(_chunk))

	setA := GenA(_chal.K1, _chal.C, n)
	setV := GenV(_chal.K2, _chal.C, _param)

	var proof Proof
	proof.Mu = _param.Pairing.NewZr().Set0()
	for i := uint32(0); i < _chal.C; i++ {
		m := _chunk[setA[i]]
		mu := _param.Pairing.NewZr().MulBig(setV[i], _param.SetFromHash(m).X())
		proof.Mu = _param.Pairing.NewZr().Add(proof.Mu, mu)
	}

	return proof
}

// https://github.com/es3ku/z22m2azuma/blob/main/sp/src/interfaces/crypt/crypt.go#L98
func VerifyProof(_param *PairingParam, _tag *Tag, _hashChunks [][]byte, _chal *Chal, _proof *Proof, _pubKey *pbc.Element) bool {

	left  := _param.Pairing.NewG1().Set1()
	right := _param.Pairing.NewG1().Set1()
	n := uint32(len(_tag.Mu))

	setA := GenA(_chal.K1, _chal.C, n)
	setV := GenV(_chal.K2, _chal.C, _param)

	for i := uint32(0); i < _chal.C; i++ {
		a := setA[i]
		v := setV[i]

		// Left
		t1 := _tag.Mu[a]
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
