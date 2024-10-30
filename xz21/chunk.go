package xz21

import(
	"slices"
	"crypto/sha256"
	"encoding/binary"
)

type ChunkSet map[uint32][]byte

func GenChunkSet(_data []byte, _chunkNum uint32) ChunkSet {
	keepList := MakeList(_chunkNum)	
	return GenChunkSubsetByIndex(_data, _chunkNum, keepList)
}

func GenChunkSubset(_data []byte, _chunkNum uint32, _chal *Chal) ChunkSet {
	keepList := _chal.GenA(_chunkNum)
	return GenChunkSubsetByIndex(_data, _chunkNum, keepList)
}

func GenChunkSubsetByIndex(_data []byte, _chunkNum uint32, _keepList []uint32) ChunkSet {
	setChunk := make(ChunkSet)
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
		setChunk.Set(i, _data[s:e])
	}

	return setChunk
}

func (this ChunkSet) Base() map[uint32][]byte {
	return (map[uint32][]byte)(this)
}

func (this ChunkSet) Set(_index uint32, _data []byte) {
	this[_index] = make([]byte, len(_data))
	copy(this[_index], _data)
}

func (this ChunkSet) Fill(_data [][]byte) {
	for i := 0; i < len(_data); i++ {
		this.Set(uint32(i), _data[i])
	}
}

func (this ChunkSet) Size() uint32 {
	return uint32(len(this.Base()))
}

func (this ChunkSet) IndexList() []uint32 {
	return MapKeys(this.Base())
}

func (this ChunkSet) Hash() DigestSet {
	listIndex := MapKeys(this.Base())
	return this.HashByIndex(listIndex)
}

func (this ChunkSet) HashByIndex(_listIndex []uint32) DigestSet {
	setDigest := make(DigestSet)
	b := make([]byte, 4)
	for _, i := range _listIndex {
		binary.LittleEndian.PutUint32(b, i)
		digest := sha256.Sum256(slices.Concat(this[i], b))
		setDigest.Set(i, digest[:])
	}
	return setDigest
}