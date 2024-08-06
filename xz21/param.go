package xz21

import (
	"encoding/json"
	"math/big"
	"os"
	"github.com/Nik-U/pbc" // v0.0.0-20181205041846-3e516ca0c5d6
)

type PairingParam struct {
	Params  *pbc.Params
	Pairing *pbc.Pairing
	G       *pbc.Element
	U       *pbc.Element
}

func GenPairingParam() PairingParam {
	var obj PairingParam
	obj.Params = pbc.GenerateA(uint32(160), uint32(512))
	obj.Pairing = obj.Params.NewPairing()
	obj.G = obj.Pairing.NewG1().Rand()
	obj.U = obj.Pairing.NewG1().Rand()
	return obj
}

func (this *PairingParam) ToXZ21Para() XZ21Para {
	var xz21Para XZ21Para
	xz21Para.Param = this.Params.String()
	xz21Para.U = this.U.Bytes()
	xz21Para.G = this.G.Bytes()
	return xz21Para
}

func GenParamFromXZ21Para(_xz21Para *XZ21Para) PairingParam {
	var err error
	var para PairingParam

	if err != nil { panic(err) }

	para.Params, err = pbc.NewParamsFromString(_xz21Para.Param)
	if err != nil { panic(err) }
	para.Pairing = pbc.NewPairing(para.Params)
	para.G = para.Pairing.NewG1().SetBytes(_xz21Para.G)
	para.U = para.Pairing.NewG1().SetBytes(_xz21Para.U)

	return para
}

func (this *PairingParam) Save(_path string) {
	tmp1 := this.ToXZ21Para()

	tmp2, err := json.MarshalIndent(&tmp1, "", "\t")
	if err != nil { panic(err) }

	f, err := os.Create(_path)
	if err != nil { panic(err) }
	defer f.Close()

	_, err = f.WriteString(string(tmp2))
	if err != nil { panic(err) }
}

func LoadPairingParam(_data []byte) PairingParam {
	var xz21Para XZ21Para
	err := json.Unmarshal(_data, &xz21Para)
	if err != nil { panic(err) }
	return GenParamFromXZ21Para(&xz21Para)
}

func LoadPairingParamFromFile(_path string) PairingParam {
	tmp, err := os.ReadFile(_path)
	if err != nil { panic(err) }

	return LoadPairingParam(tmp)
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