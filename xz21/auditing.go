package xz21

import (
	"fmt"
	"errors"
	"time"
)

const (
	WaitingProof = iota
	WaitingResult
	DoneAuditing
)

type AuditingLog struct {
	Chal *Chal `json:'chal'`
	Proof *Proof `json:'proof'`
	Result bool `json:'result'`
	Date   *time.Time `json:'date'`
	Stage  uint8 `json:'stage'`
}

type AuditingLogData struct {
	ChalData *ChalData `json:'chal'`
	ProofData ProofData `json:'proof'`
	Result bool `json:'result'`
	Date   string `json:'date'`
	Stage  uint8 `json:'stage'`
}

func (this *AuditingLog) Export() *AuditingLogData {
	var data AuditingLogData
	data.ChalData = this.Chal.Export()
	data.ProofData = this.Proof.Export()
	data.Result = this.Result
	data.Date = this.Date.Format(time.RFC3339)
	data.Stage = this.Stage
	return &data
}

func (this *AuditingLogData) Import(_param *PairingParam) *AuditingLog {
	var obj AuditingLog
	obj.Chal = this.ChalData.Import(_param)
	obj.Proof = this.ProofData.Import(_param)
	obj.Result = this.Result
	if len(this.Date) > 0 {
		t, err := time.Parse(time.RFC3339, this.Date)
		if err != nil { panic(err) }
		obj.Date = &t
	} else {
		obj.Date = nil
	}
	obj.Stage = this.Stage
	return &obj
}


func (this *AuditingLogData) LoadFromXZ21(_src *XZ21AuditingLog) error {
	var err error
	this.ChalData, err = DecodeToChalData(_src.Chal)
	if err != nil { return err }
	this.ProofData = LoadProofData(_src.Proof)
	if err != nil { return err }
	this.Result = _src.Result
	if _src.Date != nil {
		this.Date = time.Unix(_src.Date.Int64(), 0).Format(time.RFC3339)
	} else {
		this.Date = ""
	}
	this.Stage = _src.Stage

	return nil
}

func VerifyProof(_param *PairingParam, _chunkNum uint32, _setTag TagSet, _setDigest DigestSet, _chal *Chal, _proof *Proof, _pubKey *PublicKey) (bool, error) {
	left  := _param.Pairing.NewG1().Set1()
	right := _param.Pairing.NewG1().Set1()

	setA := _chal.GenA(_chunkNum)
	setV := _chal.GenV(_param)

	for i := uint32(0); i < _chal.C; i++ {
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
		if digest := _setDigest[a]; digest != nil {
			m1 := _param.SetFromHash(digest)
			m2 := _param.PowZn(m1, v)
			right  = _param.Mul(right, m2)
		} else {
			return false, errors.New(fmt.Sprintf("Invalid digests (index:%d)\n", a))
		}
	}

	u := _param.PowZn(_param.U, _proof.Base())
	right = _param.Mul(right, u)

	lhs := _param.Pairing.NewGT().Pair(left, _param.G)
	rhs := _param.Pairing.NewGT().Pair(right, _pubKey.Base())

	return lhs.Equals(rhs), nil
}
