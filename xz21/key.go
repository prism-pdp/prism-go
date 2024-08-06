package xz21

import (
	"github.com/Nik-U/pbc" // v0.0.0-20181205041846-3e516ca0c5d6
)

type PublicKey struct {
	Key *pbc.Element
}

type PrivateKey struct {
	Key *pbc.Element
}

type PublicKeyData struct {
	Key []byte
}

type PrivateKeyData struct {
	Key []byte
}

func (this *PublicKey) Export() PublicKeyData {
	var data PublicKeyData
	data.Key = this.Key.Bytes()
	return data
}

func (this *PrivateKey) Export() PrivateKeyData {
	var data PrivateKeyData
	data.Key = this.Key.Bytes()
	return data
}

func (this *PublicKeyData) Import(_param *PairingParam) PublicKey {
	var obj PublicKey
	obj.Key  = _param.Pairing.NewG1().SetBytes(this.Key)
	return obj
}

func (this *PrivateKeyData) Import(_param *PairingParam) PrivateKey {
	var obj PrivateKey
	obj.Key = _param.Pairing.NewZr().SetBytes(this.Key)
	return obj
}

func GenPairingKey(_param *PairingParam) (PublicKey, PrivateKey) {
	var sk PrivateKey
	var pk PublicKey
	sk.Key = _param.Pairing.NewZr().Rand()
	pk.Key  = _param.Pairing.NewG1().PowZn(_param.G, sk.Key)
	return pk, sk
}
