package xz21

import(
	"slices"
	"crypto/sha256"
	"encoding/binary"
)

type DigestSet struct {
	blocks map[uint32][]byte
}

func NewDigestSet() *DigestSet {
	this := DigestSet{}
	this.blocks = make(map[uint32][]byte)
	return &this
}

func (this *DigestSet) Set(_index uint32, _data []byte) {
	this.blocks[_index] = make([]byte, len(_data))
	copy(this.blocks[_index], _data)
}

func (this *DigestSet) Get(_index uint32) []byte {
	return this.blocks[_index]
}

func (this *DigestSet) Size() uint32 {
	return uint32(len(this.blocks))
}

func HashChunks(_chunkSet *ChunkSet, _targetList []uint32) *DigestSet {

	digestSet := NewDigestSet()
	b := make([]byte, 4)

	for _, i := range _targetList {
		binary.LittleEndian.PutUint32(b, i)
		digest := sha256.Sum256(slices.Concat(_chunkSet.Get(i), b))
		digestSet.Set(i, digest[:])
	}

	return digestSet
}

func HashAllChunks(_chunkSet *ChunkSet) *DigestSet {
	numChunk := _chunkSet.Size()
	targetList := make([]uint32, numChunk)
	for i := uint32(0); i < numChunk; i++ {
		targetList[i] = i
	}
	return HashChunks(_chunkSet, targetList)
}

func HashSampledChunks(_chunkSet *ChunkSet, _chal *Chal) *DigestSet {
	numChunk := _chunkSet.Size()
	setA := GenA(_chal, numChunk)
	return HashChunks(_chunkSet, setA)
}
