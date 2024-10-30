package xz21

import (
	"fmt"
	"errors"
)

type AuditingReq struct {
	Chal *Chal `json:'chal'`
	Proof *Proof `json:'proof'`
}

type AuditingReqData struct {
	ChalData *ChalData `json:'chal'`
	ProofData *ProofData `json:'proof'`
}

type AuditingLog struct {
	Req *AuditingReq `json:'req'`
	Result bool `json:'result'`
}

type AuditingLogData struct {
	ReqData *AuditingReqData `json:'req'`
	Result bool `json:'result'`
}

func (this *AuditingReq) Export() *AuditingReqData {
	var data AuditingReqData
	data.ChalData = this.Chal.Export()
	data.ProofData = this.Proof.Export()
	return &data
}

func (this *AuditingReqData) Import(_param *PairingParam) *AuditingReq {
	var obj AuditingReq
	obj.Chal = this.ChalData.Import(_param)
	obj.Proof = this.ProofData.Import(_param)
	return &obj
}

func (this *AuditingReqData) LoadFromXZ21(_src *XZ21AuditingReq) error {
	var err error
	this.ChalData, err = DecodeToChalData(_src.Chal)
	if err != nil { return err }
	this.ProofData, err = DecodeToProofData(_src.Proof)
	if err != nil { return err }

	return nil
}

func (this *AuditingLog) Export() *AuditingLogData {
	var data AuditingLogData
	data.ReqData = this.Req.Export()
	data.Result = this.Result
	return &data
}

func (this *AuditingLogData) Import(_param *PairingParam) *AuditingLog {
	var obj AuditingLog
	obj.Req = this.ReqData.Import(_param)
	obj.Result = this.Result
	return &obj
}


func (this *AuditingLogData) LoadFromXZ21(_src *XZ21AuditingLog) error {
	var err error
	this.ReqData.ChalData, err = DecodeToChalData(_src.Req.Chal)
	if err != nil { return err }
	this.ReqData.ProofData, err = DecodeToProofData(_src.Req.Proof)
	if err != nil { return err }
	this.Result = _src.Result

	return nil
}

func (this *AuditingReq) VerifyProof(_param *PairingParam, _chunkNum uint32, _setTag TagSet, _digestSet *DigestSet, _pubKey *PublicKey) (bool, error) {
	left  := _param.Pairing.NewG1().Set1()
	right := _param.Pairing.NewG1().Set1()

	setA := this.Chal.GenA(_chunkNum)
	setV := this.Chal.GenV(_param)

	for i := uint32(0); i < this.Chal.C; i++ {
		a := setA[i]
		v := setV[i]

		// Left
		if t1, ok := _setTag[a]; ok {
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

	u := _param.PowZn(_param.U, this.Proof.Mu)
	right = _param.Mul(right, u)

	lhs := _param.Pairing.NewGT().Pair(left, _param.G)
	rhs := _param.Pairing.NewGT().Pair(right, _pubKey.Elem())

	return lhs.Equals(rhs), nil
}
