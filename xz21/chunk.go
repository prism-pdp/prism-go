package xz21

import(
	"slices"
	"crypto/sha256"
	"encoding/binary"
)

type ChunkSet struct {
	blocks map[uint32][]byte
	TotalNum uint32
}

func NewChunkSet(_totalNum uint32) *ChunkSet {
	this := ChunkSet{}
	this.blocks = make(map[uint32][]byte)
	this.TotalNum = _totalNum
	return &this
}

func GenChunkSet(_data []byte, _chunkNum uint32) *ChunkSet {
	keepList := MakeList(_chunkNum)	
	return GenChunkSubset(_data, _chunkNum, keepList)
}

func GenChunkSubsetByChal(_data []byte, _chunkNum uint32, _chal *Chal) *ChunkSet {
	setA := GenA(_chal, _chunkNum)
	return GenChunkSubset(_data, _chunkNum, setA)
}

func GenChunkSubset(_data []byte, _chunkNum uint32, _keepList []uint32) *ChunkSet {
	chunkSet := NewChunkSet(_chunkNum)
	dataSize := uint32(len(_data))

	chunkSize  := dataSize / _chunkNum
	chunkSizeR := dataSize % _chunkNum

	var s, e uint32

	for _, i := range _keepList {
		s = i * chunkSize
		e = s + chunkSize
		if i == (_chunkNum - 1) {
			e = e + chunkSizeR
		}
		chunkSet.Set(i, _data[s:e])
	}

	return chunkSet
}

func (this *ChunkSet) Set(_index uint32, _data []byte) {
	this.blocks[_index] = make([]byte, len(_data))
	copy(this.blocks[_index], _data)
}

func (this *ChunkSet) Fill(_data [][]byte) {
	for i := 0; i < len(_data); i++ {
		this.Set(uint32(i), _data[i])
	}
}

func (this *ChunkSet) Get(_index uint32) []byte {
	return this.blocks[_index]
}

func (this *ChunkSet) Size() uint32 {
	return uint32(len(this.blocks))
}

func (this *ChunkSet) IndexList() []uint32 {
	return MapKeys(this.blocks)
}

func (this *ChunkSet) Hash() *DigestSet {
	listIndex := MapKeys(this.blocks)
	return this.HashByIndex(listIndex)
}

func (this *ChunkSet) HashByChal(_chal *Chal) *DigestSet {
	setA := GenA(_chal, this.Size())
	return this.HashByIndex(setA)
}

func (this *ChunkSet) HashByIndex(_listIndex []uint32) *DigestSet {
	digestSet := NewDigestSet(this.TotalNum)
	b := make([]byte, 4)

	for _, i := range _listIndex {
		binary.LittleEndian.PutUint32(b, i)
		digest := sha256.Sum256(slices.Concat(this.Get(i), b))
		digestSet.Set(i, digest[:])
	}

	return digestSet
}