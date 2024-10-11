package xz21

import (
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
	setA := GenA(_chal, this.Size)
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

func GenTag(_param *PairingParam, _privKey *pbc.Element, _chunkSet *ChunkSet) (Tag, *DigestSet) {
	digestSet := HashAllChunks(_chunkSet)

	var tag Tag
	tag.Size = _chunkSet.Size()
	tag.G = make(map[uint32]*pbc.Element, tag.Size)

	for i := uint32(0); i < tag.Size; i++ {
		e1 := _param.SetFromHash(digestSet.Get(i))
		e2 := _param.SetFromHash(_chunkSet.Get(i))
		e3 := _param.PowBig(_param.U, e2.X())
		e4 := _param.Mul(e1, e3)
		tag.G[i] = _param.PowZn(e4, _privKey)
	}

	return tag, digestSet
}