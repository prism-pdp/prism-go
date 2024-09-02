package xz21

import (
	"encoding/json"
	"math/big"
	"os"
	"github.com/Nik-U/pbc" // v0.0.0-20181205041846-3e516ca0c5d6
)

type PairingParam struct {
	P       *pbc.Params
	Pairing *pbc.Pairing
	G       *pbc.Element
	U       *pbc.Element
}

func GenPairingParam() PairingParam {
	var obj PairingParam
	obj.P = pbc.GenerateA(uint32(160), uint32(512))
	obj.Pairing = obj.P.NewPairing()
	obj.G = obj.Pairing.NewG1().Rand()
	obj.U = obj.Pairing.NewG1().Rand()
	return obj
}

func (this *PairingParam) ToXZ21Param() XZ21Param {
	var xz21Param XZ21Param
	xz21Param.P = this.P.String()
	xz21Param.U = this.U.Bytes()
	xz21Param.G = this.G.Bytes()
	return xz21Param
}

func GenParamFromXZ21Param(_xz21Param *XZ21Param) PairingParam {
	var err error
	var param PairingParam

	if err != nil { panic(err) }

	param.P, err = pbc.NewParamsFromString(_xz21Param.P)
	if err != nil { panic(err) }
	param.Pairing = pbc.NewPairing(param.P)
	param.G = param.Pairing.NewG1().SetBytes(_xz21Param.G)
	param.U = param.Pairing.NewG1().SetBytes(_xz21Param.U)

	return param
}

func (this *PairingParam) Save(_path string) {
	tmp1 := this.ToXZ21Param()

	tmp2, err := json.MarshalIndent(&tmp1, "", "\t")
	if err != nil { panic(err) }

	f, err := os.Create(_path)
	if err != nil { panic(err) }
	defer f.Close()

	_, err = f.WriteString(string(tmp2))
	if err != nil { panic(err) }
}

func (this *PairingParam) SetFromHash(_hash []byte) *pbc.Element {
	return this.Pairing.NewG1().SetFromHash(_hash)
}

func (this *PairingParam) SetBytes(_buf []byte) *pbc.Element {
	return this.Pairing.NewG1().SetBytes(_buf)
}

func (this *PairingParam) PowBig(_target *pbc.Element, _i *big.Int) *pbc.Element {
	return this.Pairing.NewG1().PowBig(_target, _i)
}

func (this *PairingParam) PowZn(_target, _i *pbc.Element) *pbc.Element {
	return this.Pairing.NewG1().PowZn(_target, _i)
}

func (this *PairingParam) Mul(_x, _i *pbc.Element) *pbc.Element {
	return this.Pairing.NewG1().Mul(_x, _i)
}

func (this *PairingParam) MulZn(_x, _i *pbc.Element) *pbc.Element {
	return this.Pairing.NewG1().MulZn(_x, _i)
}
