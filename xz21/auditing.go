package xz21

import (
	"fmt"
	"errors"

	"github.com/Nik-U/pbc" // v0.0.0-20181205041846-3e516ca0c5d6
)

type AuditingReq struct {
	Chal Chal `json:'chal'`
	Proof Proof `json:'proof'`
}

type AuditingReqData struct {
	ChalData ChalData `json:'chal'`
	ProofData ProofData `json:'proof'`
}

type AuditingLog struct {
	ChalData ChalData `json:'chal'`
	ProofData ProofData `json:'proof'`
	Result bool `json:'result'`
}

func (this *AuditingReq) Export() AuditingReqData {
	var data AuditingReqData
	data.ChalData = this.Chal.Export()
	data.ProofData = this.Proof.Export()
	return data
}

func (this *AuditingReqData) Import(_param *PairingParam) AuditingReq {
	var obj AuditingReq
	obj.Chal = this.ChalData.Import(_param)
	obj.Proof = this.ProofData.Import(_param)
	return obj
}

func (this *AuditingReqData) ToXZ21(_src *XZ21AuditingReq) error {
	var err error
	this.ChalData, err = DecodeToChalData(_src.Chal)
	if err != nil { return err }
	this.ProofData, err = DecodeToProofData(_src.Proof)
	if err != nil { return err }

	return nil
}

func (this *AuditingReqData) ImportChal(_param *PairingParam) Chal {
	return this.ChalData.Import(_param)
}

func (this *AuditingReqData) ImportProof(_param *PairingParam) Proof {
	return this.ProofData.Import(_param)
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

func VerifyProof(_param *PairingParam, _tag *Tag, _digestSet *DigestSet, _auditingReq *AuditingReq, _pubKey *pbc.Element) (bool, error) {

	chal := &_auditingReq.Chal
	proof := &_auditingReq.Proof

	left  := _param.Pairing.NewG1().Set1()
	right := _param.Pairing.NewG1().Set1()
	n := _tag.Size

	setA := GenA(chal, n)
	setV := GenV(chal, _param)

	for i := uint32(0); i < chal.C; i++ {
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

	u := _param.PowZn(_param.U, proof.Mu)
	right = _param.Mul(right, u)

	lhs := _param.Pairing.NewGT().Pair(left, _param.G)
	rhs := _param.Pairing.NewGT().Pair(right, _pubKey)

	return lhs.Equals(rhs), nil
}
