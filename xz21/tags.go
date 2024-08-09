package xz21

import (
	"crypto/sha256"
	"encoding/binary"
	"slices"

	"github.com/Nik-U/pbc"
)

type Tag struct {
	Size uint32
	Mu []*pbc.Element
}

type TagData struct {
	Size uint32
	Mu [][]byte
}

func (this *Tag) Export() TagData {
	var data TagData
	data.Size = uint32(len(this.Mu))
	data.Mu = make([][]byte, data.Size)
	for i, v := range this.Mu {
		data.Mu[i] = v.Bytes()
	}
	return data
}

func (this *TagData) Import(_param *PairingParam) Tag {
	var obj Tag
	obj.Size = uint32(len(this.Mu))
	obj.Mu = make([]*pbc.Element, obj.Size)
	for i, v := range this.Mu {
		obj.Mu[i] = _param.Pairing.NewG1().SetBytes(v)
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
func HashChunks(_chunks [][]byte) [][]byte {
	numChunk := uint32(len(_chunks))
	hash := make([][]byte, numChunk)

	for i := uint32(0); i < numChunk; i++ {
		hash[i]  = make([]byte, 32)
	}

	b := make([]byte, 4)
	for i := uint32(0); i < numChunk; i++ {
		binary.LittleEndian.PutUint32(b, i)
		h := sha256.Sum256(slices.Concat(_chunks[i], b))
		copy(hash[i], h[:])
	}

	return hash
}

func GenTag(_param *PairingParam, _privKey *pbc.Element, _chunks [][]byte) (Tag, [][]byte) {
	hashChunks := HashChunks(_chunks)

	var tag Tag
	tag.Size = uint32(len(_chunks))
	tag.Mu = make([]*pbc.Element, tag.Size)

	for i := uint32(0); i < tag.Size; i++ {
		e1 := _param.SetFromHash(hashChunks[i])
		e2 := _param.SetFromHash(_chunks[i])
		e3 := _param.PowBig(_param.U, e2.X())
		e4 := _param.Mul(e1, e3)
		tag.Mu[i] = _param.PowZn(e4, _privKey)
	}

	return tag, hashChunks
}