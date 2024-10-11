package xz21

import (
	"crypto/sha256"
	"encoding/binary"
	"slices"

	"github.com/Nik-U/pbc"
)

type Tag struct {
	Size uint32
	G map[uint32]*pbc.Element // gamma
}

type TagData struct {
	Size uint32
	G map[uint32][]byte // gamma
}

func NewTagData() *TagData {
	this := &TagData{Size: 0, G: make(map[uint32][]byte)}
	return this
}

func (this *Tag) Export() TagData {
	var data TagData
	data.Size = this.Size
	data.G = make(map[uint32][]byte)
	for i, v := range this.G {
		data.G[i] = v.Bytes()
	}
	return data
}

func (this *TagData) _import(_param *PairingParam, _targetList []uint32) Tag {
	var obj Tag
	obj.Size = this.Size
	obj.G = make(map[uint32]*pbc.Element)
	for _, i := range _targetList {
		obj.G[i] = _param.Pairing.NewG1().SetBytes(this.G[i])
	}
	return obj
}

func (this *TagData) ImportAll(_param *PairingParam) Tag {
	targetList := make([]uint32, 0, this.Size)
	for k, _ := range this.G {
		targetList = append(targetList, k)
	}
	return this._import(_param, targetList)
}

func (this *TagData) ImportSubset(_param *PairingParam, _chal *Chal) Tag {
	setA := GenA(_chal.K1, _chal.C, this.Size)
	tag := this._import(_param, setA)
	return tag
}

func (this *TagData) Copy(_from *TagData, _indexList []uint32) {
	// Size
	this.Size = _from.Size
	// G
	for _, i := range _indexList {
		if _, ok := this.G[i]; !ok {
			this.G[i] = make([]byte, len(_from.G[i]))
			copy(this.G[i], _from.G[i])
		}
	}
}

// https://github.com/es3ku/z22m2azuma/blob/main/user/src/interfaces/crypt/content.go#L23
func HashChunks(_chunks [][]byte, _targetList []uint32) map[uint32][]byte {

	sampledDigests := make(map[uint32][]byte)
	b := make([]byte, 4)

	for _, i := range _targetList {
		binary.LittleEndian.PutUint32(b, i)
		sampledDigest := sha256.Sum256(slices.Concat(_chunks[i], b))

		sampledDigests[i] = make([]byte, 32)
		copy(sampledDigests[i], sampledDigest[:])
	}

	return sampledDigests
}

func HashAllChunks(_chunks [][]byte) map[uint32][]byte {
	numChunk := uint32(len(_chunks))
	targetList := make([]uint32, numChunk)
	for i := uint32(0); i < numChunk; i++ {
		targetList[i] = i
	}
	return HashChunks(_chunks, targetList)
}

func HashSampledChunks(_chunks [][]byte, _chal *Chal) map[uint32][]byte {
	numChunk := uint32(len(_chunks))
	setA := GenA(_chal.K1, _chal.C, numChunk)
	return HashChunks(_chunks, setA)
}

func GenTag(_param *PairingParam, _privKey *pbc.Element, _chunks [][]byte) (Tag, map[uint32][]byte) {
	hashChunks := HashAllChunks(_chunks)

	var tag Tag
	tag.Size = uint32(len(_chunks))
	tag.G = make(map[uint32]*pbc.Element, tag.Size)

	for i := uint32(0); i < tag.Size; i++ {
		e1 := _param.SetFromHash(hashChunks[i])
		e2 := _param.SetFromHash(_chunks[i])
		e3 := _param.PowBig(_param.U, e2.X())
		e4 := _param.Mul(e1, e3)
		tag.G[i] = _param.PowZn(e4, _privKey)
	}

	return tag, hashChunks
}