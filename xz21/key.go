package xz21

import (
	"encoding/json"
	"os"
	"github.com/Nik-U/pbc" // v0.0.0-20181205041846-3e516ca0c5d6
)

type PairingKey struct {
	PrivateKey *pbc.Element
	PublicKey *pbc.Element
}

type PairingKeyData struct {
	PrivateKey []byte
	PublicKey  []byte
}

func (this *PairingKey) Export() PairingKeyData {
	var keyData PairingKeyData
	keyData.PrivateKey = this.PrivateKey.Bytes()
	keyData.PublicKey = this.PublicKey.Bytes()
	return keyData
}

func (this *PairingKeyData) Import(_param *PairingParam) PairingKey {
	var keyObj PairingKey
	keyObj.PrivateKey = _param.Pairing.NewZr().SetBytes(this.PrivateKey)
	keyObj.PublicKey  = _param.Pairing.NewG1().SetBytes(this.PublicKey)
	return keyObj
}

func GenPairingKey(_param *PairingParam) PairingKey {
	var obj PairingKey
	obj.PrivateKey = _param.Pairing.NewZr().Rand()
	obj.PublicKey  = _param.Pairing.NewG1().PowZn(_param.G, obj.PrivateKey)
	return obj
}

func (this *PairingKey) Save(_path string) {
	tmp1 := this.Export()

	tmp2, err := json.MarshalIndent(&tmp1, "", "\t")
	if err != nil { panic(err) }

	f, err := os.Create(_path)
	if err != nil { panic(err) }
	defer f.Close()

	_, err = f.WriteString(string(tmp2))
	if err != nil { panic(err) }
}

func LoadPairingKey(_data []byte, _param *PairingParam) PairingKey {
	var pairingKeyData PairingKeyData
	err := json.Unmarshal(_data, &pairingKeyData)
	if err != nil { panic(err) }
	return pairingKeyData.Import(_param)
}

func LoadPairingKeyFromFile(_path string, _param *PairingParam) PairingKey {
	tmp, err := os.ReadFile(_path)
	if err != nil { panic(err) }

	return LoadPairingKey(tmp, _param)
}