package xz21

import (
	"fmt"
	"errors"

	"github.com/Nik-U/pbc" // v0.0.0-20181205041846-3e516ca0c5d6
)

type AuditingReq struct {
	ChalData ChalData `json:'chal'`
	ProofData ProofData `json:'proof'`
}

type AuditingLog struct {
	ChalData ChalData `json:'chal'`
	ProofData ProofData `json:'proof'`
	Result bool `json:'result'`
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

func VerifyProof(_param *PairingParam, _tag *Tag, _digestSet *DigestSet, _chal *Chal, _proof *Proof, _pubKey *pbc.Element) (bool, error) {

	left  := _param.Pairing.NewG1().Set1()
	right := _param.Pairing.NewG1().Set1()
	n := _tag.Size

	setA := GenA(_chal.K1, _chal.C, n)
	setV := GenV(_chal.K2, _chal.C, _param)

	for i := uint32(0); i < _chal.C; i++ {
		a := setA[i]
		v := setV[i]

		// Left
		if t1, ok := _tag.G[a]; ok {
			t2 := _param.PowZn(t1, v)
			left = _param.Mul(left, t2)
		} else {
			return false, errors.New(fmt.Sprintf("Invalid tags (index:%d)\n", a))
		}

		// Right
		if digest := _digestSet.Get(a); digest != nil {
			m1 := _param.SetFromHash(digest)
			m2 := _param.PowZn(m1, v)
			right  = _param.Mul(right, m2)
		} else {
			return false, errors.New(fmt.Sprintf("Invalid digests (index:%d)\n", a))
		}
	}

	u := _param.PowZn(_param.U, _proof.Mu)
	right = _param.Mul(right, u)

	lhs := _param.Pairing.NewGT().Pair(left, _param.G)
	rhs := _param.Pairing.NewGT().Pair(right, _pubKey)

	return lhs.Equals(rhs), nil
}
