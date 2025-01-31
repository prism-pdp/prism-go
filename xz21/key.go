package xz21

import (
	"github.com/Nik-U/pbc" // v0.0.0-20181205041846-3e516ca0c5d6
)

type PublicKey pbc.Element
type PrivateKey pbc.Element

type PublicKeyData []byte
type PrivateKeyData []byte

func (this *PublicKey) Base() *pbc.Element {
	return (*pbc.Element)(this)
}

func (this *PrivateKey) Base() *pbc.Element {
	return (*pbc.Element)(this)
}

func (this *PublicKeyData) Base() []byte {
	return ([]byte)(*this)
}

func (this *PrivateKeyData) Base() []byte {
	return ([]byte)(*this)
}

func (this *PublicKey) Export() PublicKeyData {
	var data PublicKeyData
	data = this.Base().Bytes()
	return data
}

func (this *PrivateKey) Export() PrivateKeyData {
	var data PrivateKeyData
	data = this.Base().Bytes()
	return data
}

func (this *PublicKeyData) Load(_key []byte) {
	*this = (PublicKeyData)(make([]byte, len(_key)))
	copy(this.Base(), _key)
}

func (this *PublicKeyData) Import(_param *PairingParam) *PublicKey {
	key := _param.Pairing.NewG1().SetBytes([]byte(*this))
	return (*PublicKey)(key)
}

func (this *PrivateKeyData) Import(_param *PairingParam) *PrivateKey {
	key := _param.Pairing.NewZr().SetBytes([]byte(*this))
	return (*PrivateKey)(key)
}

func GenPairingKey(_param *PairingParam) (*PublicKey, *PrivateKey) {
	sk := _param.Pairing.NewZr().Rand()
	pk := _param.Pairing.NewG1().PowZn(_param.G, sk)
	return (*PublicKey)(pk), (*PrivateKey)(sk)
}
