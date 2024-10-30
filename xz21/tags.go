package xz21

import (
	"github.com/Nik-U/pbc"
)

type TagSet map[uint32]*pbc.Element
type TagDataSet map[uint32][]byte

func (this TagSet) Base() map[uint32]*pbc.Element {
	return (map[uint32]*pbc.Element)(this)
}

func (this TagSet) Size() uint32 {
	return (uint32)(len(this.Base()))
}

func (this TagDataSet) Base() map[uint32][]byte {
	return (map[uint32][]byte)(this)
}

func (this TagDataSet) Size() uint32 {
	return (uint32)(len(this.Base()))
}

func (this TagSet) Export() TagDataSet {
	data := make(TagDataSet)
	for i, v := range this.Base() {
		data[i] = v.Bytes()
	}
	return data
}

func (this TagDataSet) ImportByIndex(_param *PairingParam, _targetList []uint32) TagSet {
	setTag := make(TagSet)
	// obj.Size = this.Size
	// obj = make(map[uint32]*pbc.Element)
	for _, i := range _targetList {
		setTag[i] = _param.Pairing.NewG1().SetBytes(this[i])
	}
	return setTag
}

func (this TagDataSet) Import(_param *PairingParam) TagSet {
	targetList := make([]uint32, 0, this.Size())
	for k, _ := range this.Base() {
		targetList = append(targetList, k)
	}
	return this.ImportByIndex(_param, targetList)
}

func (this TagDataSet) ImportSubset(_param *PairingParam, _chunkNum uint32, _chal *Chal) TagSet {
	setA := _chal.GenA(_chunkNum)
	setTag := this.ImportByIndex(_param, setA)
	return setTag
}

func (this TagDataSet) DuplicateSubset(_chunkNum uint32, _chal *Chal) TagDataSet {
	setA := _chal.GenA(_chunkNum)
	return this.DuplicateByIndex(setA)
}

func (this TagDataSet) DuplicateByIndex(_listIndex []uint32) TagDataSet {
	newSet := make(TagDataSet)
	for _, i := range _listIndex {
		if _, ok := newSet[i]; !ok {
			newSet[i] = make([]byte, len(this[i]))
			copy(newSet[i], this[i])
		}
	}
	return newSet
}

func GenTags(_param *PairingParam, _privKey *PrivateKey, _setChunk ChunkSet) (TagSet, *DigestSet) {
	digestSet := _setChunk.Hash()
	setTag := make(TagSet)

	for i := uint32(0); i < _setChunk.Size(); i++ {
		e1 := _param.SetFromHash(digestSet.Get(i))
		e2 := _param.SetFromHash(digestSet.Get(i))
		e3 := _param.PowBig(_param.U, e2.X())
		e4 := _param.Mul(e1, e3)
		setTag[i] = _param.PowZn(e4, _privKey.Elem())
	}

	return setTag, digestSet
}