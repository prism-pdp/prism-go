package xz21

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
	"slices"

	"github.com/Nik-U/pbc"
)

func SplitData(_data []byte, _chunkSize uint32) ([][]byte, error) {
	if _chunkSize <= 0 { return nil, fmt.Errorf("invalid split num: %d", _chunkSize) }

	var chunk [][]byte
	dataSize := uint32(len(_data))
	rows := uint32(math.Ceil(float64(dataSize) / float64(_chunkSize)))

	var s, e, processedSize uint32

	processedSize = 0
	for i := uint32(0); i < rows; i++ {
		if i == (rows - 1) {
			s = i * _chunkSize
			e = s + (dataSize - processedSize)
		} else {
			s = i * _chunkSize
			e = s + _chunkSize
		}
		chunk = append(chunk, _data[s:e])
		processedSize += (e - s)
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

func GenTags(_param *PairingParam, _privKey *pbc.Element, _chunks [][]byte) ([]*pbc.Element, [][]byte, uint32) {

	numTags := uint32(len(_chunks))
	hashChunks := HashChunks(_chunks)
	tags := make([]*pbc.Element, numTags)

	for i := uint32(0); i < numTags; i++ {
		e1 := _param.SetFromHash(hashChunks[i])
		e2 := _param.SetFromHash(_chunks[i])
		e3 := _param.PowBig(_param.U, e2.X())
		e4 := _param.Mul(e1, e3)
		tags[i] = _param.PowZn(e4, _privKey)
	}

	return tags, hashChunks, numTags
}