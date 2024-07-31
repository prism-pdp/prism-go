package xz21

import (
	"encoding/base64"
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

type PairingParamStr struct {
	Params string `json:params`
	G      string `json:g`
	U      string `json:u`
}

func (this *PairingParam) Gen() {
	this.Params = pbc.GenerateA(uint32(160), uint32(512))
	this.Pairing = this.Params.NewPairing()
	this.G = this.Pairing.NewG1().Rand()
	this.U = this.Pairing.NewG1().Rand()
}

func (this *PairingParam) ToString() PairingParamStr {
	return PairingParamStr{
		Params: this.Params.String(),
		G: base64.StdEncoding.EncodeToString(this.G.Bytes()),
		U: base64.StdEncoding.EncodeToString(this.U.Bytes()),
	}
}

func (this *PairingParam) FromString(_string string) {
	var paramStr PairingParamStr
	var err error

	err = json.Unmarshal([]byte(_string), &paramStr)
	if err != nil { panic(err) }

	this.Params, err = pbc.NewParamsFromString(paramStr.Params)
	if err != nil { panic(err) }

	this.Pairing = this.Params.NewPairing()

	tmpG, err := base64.StdEncoding.DecodeString(paramStr.G)
	if err != nil { panic(err) }
	this.G = this.Pairing.NewG1().SetBytes(tmpG)

	tmpU, err := base64.StdEncoding.DecodeString(paramStr.U)
	if err != nil { panic(err) }
	this.U = this.Pairing.NewG1().SetBytes(tmpU)
}

func (this *PairingParam) ToXZ21Para() *XZ21Para {
	para := make(XZ21Para)
	para.Pairing = this.Params.String()
	para.U = this.U.Bytes()
	para.G = this.G.Bytes()
	return para
}

func (this *PairingParam) Save(_path string) {
	tmp1 := this.ToString()

	tmp2, err := json.MarshalIndent(tmp1, "", "\t")
	if err != nil { panic(err) }

	f, err := os.Create(_path)
	if err != nil { panic(err) }
	defer f.Close()

	_, err = f.WriteString(string(tmp2))
	if err != nil { panic(err) }
}

func (this *PairingParam) Load(_path string) {
	tmp, err := os.ReadFile(_path)
	if err != nil { panic(err) }

	this.FromString(string(tmp))
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