package xz21

import (
	"crypto/sha256"
	"encoding/binary"
	"slices"

	"github.com/Nik-U/pbc"
)

type Tag struct {
	Size uint32
	G []*pbc.Element // gamma
}

type TagData struct {
	Size uint32
	G [][]byte // gamma
}

func (this *Tag) Export() TagData {
	var data TagData
	data.Size = uint32(len(this.G))
	data.G = make([][]byte, data.Size)
	for i, v := range this.G {
		data.G[i] = v.Bytes()
	}
	return data
}

func (this *TagData) Import(_param *PairingParam) Tag {
	var obj Tag
	obj.Size = uint32(len(this.G))
	obj.G = make([]*pbc.Element, obj.Size)
	for i, v := range this.G {
		obj.G[i] = _param.Pairing.NewG1().SetBytes(v)
	}
	return obj
}

func SplitData(_data []byte, _chunkNum uint32) ([][]byte, error) {
	var chunk [][]byte
	dataSize := uint32(len(_data))

	chunkSize  := dataSize / _chunkNum
	chunkSizeR := dataSize % _chunkNum

	var s, e uint32

	for i := uint32(0); i < _chunkNum; i++ {
		s = i * chunkSize
		e = s + chunkSize
		if i == (_chunkNum - 1) {
			e = e + chunkSizeR
		}
		chunk = append(chunk, _data[s:e])
	}

	return chunk, nil
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
	tag.G = make([]*pbc.Element, tag.Size)

	for i := uint32(0); i < tag.Size; i++ {
		e1 := _param.SetFromHash(hashChunks[i])
		e2 := _param.SetFromHash(_chunks[i])
		e3 := _param.PowBig(_param.U, e2.X())
		e4 := _param.Mul(e1, e3)
		tag.G[i] = _param.PowZn(e4, _privKey)
	}

	return tag, hashChunks
}