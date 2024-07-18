package xz21

import (
	"encoding/base64"
	"encoding/json"
	"os"
	"github.com/Nik-U/pbc" // v0.0.0-20181205041846-3e516ca0c5d6
)

type PairingParam struct {
	Params  *pbc.Params
	Pairing *pbc.Pairing
	G       *pbc.Element
}

type PairingParamStr struct {
	Params string `json:params`
	G      string `json:g`
}

func (this *PairingParam) Gen() {
	this.Params = pbc.GenerateA(uint32(160), uint32(512))
	this.Pairing = this.Params.NewPairing()
	this.G = this.Pairing.NewG1().Rand()
}

func (this *PairingParam) ToString() PairingParamStr {
	return PairingParamStr{
		Params: this.Params.String(),
		G: base64.StdEncoding.EncodeToString(this.G.Bytes()),
	}
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
	tmp1, err := os.ReadFile(_path)
	if err != nil { panic(err) }

	var paramStr PairingParamStr
	err = json.Unmarshal(tmp1, &paramStr)
	if err != nil { panic(err) }

	this.Params, err = pbc.NewParamsFromString(paramStr.Params)
	if err != nil { panic(err) }

	this.Pairing = this.Params.NewPairing()

	tmpG, err := base64.StdEncoding.DecodeString(paramStr.G)
	if err != nil { panic(err) }

	this.G = this.Pairing.NewG1().SetBytes(tmpG)
}