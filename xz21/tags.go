package xz21

import (
	"github.com/Nik-U/pbc"
)

type TagSet struct {
	Size uint32
	Tag map[uint32]*pbc.Element // gamma
}

type TagDataSet struct {
	Size uint32
	TagData map[uint32][]byte // gamma
}

func NewTagDataSet() *TagDataSet {
	this := &TagDataSet{Size: 0, TagData: make(map[uint32][]byte)}
	return this
}

func (this *TagSet) Export() TagDataSet {
	var data TagDataSet
	data.Size = this.Size
	data.TagData = make(map[uint32][]byte)
	for i, v := range this.Tag {
		data.TagData[i] = v.Bytes()
	}
	return data
}

func (this *TagDataSet) _import(_param *PairingParam, _targetList []uint32) TagSet {
	var obj TagSet
	obj.Size = this.Size
	obj.Tag = make(map[uint32]*pbc.Element)
	for _, i := range _targetList {
		obj.Tag[i] = _param.Pairing.NewG1().SetBytes(this.TagData[i])
	}
	return obj
}

func (this *TagDataSet) ImportAll(_param *PairingParam) TagSet {
	targetList := make([]uint32, 0, this.Size)
	for k, _ := range this.TagData {
		targetList = append(targetList, k)
	}
	return this._import(_param, targetList)
}

func (this *TagDataSet) ImportSubset(_param *PairingParam, _chal *Chal) TagSet {
	setA := GenA(_chal, this.Size)
	tagSet := this._import(_param, setA)
	return tagSet
}

func (this *TagDataSet) Copy(_from *TagDataSet, _indexList []uint32) {
	// Size
	this.Size = _from.Size
	// Tag
	for _, i := range _indexList {
		if _, ok := this.TagData[i]; !ok {
			this.TagData[i] = make([]byte, len(_from.TagData[i]))
			copy(this.TagData[i], _from.TagData[i])
		}
	}
}

func GenTags(_param *PairingParam, _privKey *pbc.Element, _chunkSet *ChunkSet) (TagSet, *DigestSet) {
	digestSet := _chunkSet.Hash()

	var tagSet TagSet
	tagSet.Size = _chunkSet.Size()
	tagSet.Tag = make(map[uint32]*pbc.Element, tagSet.Size)

	for i := uint32(0); i < tagSet.Size; i++ {
		e1 := _param.SetFromHash(digestSet.Get(i))
		e2 := _param.SetFromHash(digestSet.Get(i))
		e3 := _param.PowBig(_param.U, e2.X())
		e4 := _param.Mul(e1, e3)
		tagSet.Tag[i] = _param.PowZn(e4, _privKey)
	}

	return tagSet, digestSet
}