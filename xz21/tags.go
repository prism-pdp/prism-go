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

func (this *TagDataSet) ImportByIndex(_param *PairingParam, _targetList []uint32) TagSet {
	var obj TagSet
	obj.Size = this.Size
	obj.Tag = make(map[uint32]*pbc.Element)
	for _, i := range _targetList {
		obj.Tag[i] = _param.Pairing.NewG1().SetBytes(this.TagData[i])
	}
	return obj
}

func (this *TagDataSet) Import(_param *PairingParam) TagSet {
	targetList := make([]uint32, 0, this.Size)
	for k, _ := range this.TagData {
		targetList = append(targetList, k)
	}
	return this.ImportByIndex(_param, targetList)
}

func (this *TagDataSet) ImportSubset(_param *PairingParam, _chunkNum uint32, _chal *Chal) TagSet {
	setA := _chal.GenA(_chunkNum)
	tagSet := this.ImportByIndex(_param, setA)
	return tagSet
}

func (this *TagDataSet) DuplicateSubset(_chunkNum uint32, _chal *Chal) *TagDataSet {
	setA := _chal.GenA(_chunkNum)
	return this.DuplicateByIndex(setA)
}

func (this *TagDataSet) DuplicateByIndex(_listIndex []uint32) *TagDataSet {
	newSet := NewTagDataSet()
	// Size
	newSet.Size = this.Size
	// Tag
	for _, i := range _listIndex {
		if _, ok := newSet.TagData[i]; !ok {
			newSet.TagData[i] = make([]byte, len(this.TagData[i]))
			copy(newSet.TagData[i], this.TagData[i])
		}
	}
	return newSet
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