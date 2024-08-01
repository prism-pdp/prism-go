package xz21

import (
	"encoding/base64"
	"encoding/json"
	"os"
	"github.com/Nik-U/pbc" // v0.0.0-20181205041846-3e516ca0c5d6
)

type PairingKey struct {
	PrivateKey *pbc.Element
	PublicKey *pbc.Element
}

type PairingKeyStr struct {
	PrivateKey string `json:private_key`
	PublicKey string `json:public_key`
}

func (this *PairingKey) Gen(_param *PairingParam) {
	this.PrivateKey = _param.Pairing.NewZr().Rand()
	this.PublicKey  = _param.Pairing.NewG1().PowZn(_param.G, this.PrivateKey)
}

func (this *PairingKey) ToString() PairingKeyStr {
	return PairingKeyStr{
		PrivateKey: base64.StdEncoding.EncodeToString(this.PrivateKey.Bytes()),
		PublicKey:  base64.StdEncoding.EncodeToString(this.PublicKey.Bytes()),
	}
}

func (this *PairingKey) Save(_path string) {
	tmp1 := this.ToString()

	tmp2, err := json.MarshalIndent(tmp1, "", "\t")
	if err != nil { panic(err) }

	f, err := os.Create(_path)
	if err != nil { panic(err) }
	defer f.Close()

	_, err = f.WriteString(string(tmp2))
	if err != nil { panic(err) }
}

func (this *PairingKey) Load(_path string, _param *PairingParam) {
	tmp1, err := os.ReadFile(_path)
	if err != nil { panic(err) }

	var keyStr PairingKeyStr
	err = json.Unmarshal(tmp1, &keyStr)
	if err != nil { panic(err) }

	bytesPrivateKey, err := base64.StdEncoding.DecodeString(keyStr.PrivateKey)
	if err != nil { panic(err) }
	this.PrivateKey = _param.Pairing.NewZr().SetBytes(bytesPrivateKey)

	bytesPublicKey, err  := base64.StdEncoding.DecodeString(keyStr.PublicKey)
	if err != nil { panic(err) }
	this.PublicKey = _param.Pairing.NewG1().SetBytes(bytesPublicKey)
}