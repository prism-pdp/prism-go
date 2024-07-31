package xz21

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
	"slices"

	"github.com/Nik-U/pbc" // v0.0.0-20181205041846-3e516ca0c5d6
)

type Metadata struct {
	Size uint32
	Hash [][]byte
	Tags []*pbc.Element
}

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
func HashChunk(_chunk [][]byte) [][]byte {
	numChunk := uint32(len(_chunk))
	hash := make([][]byte, numChunk)

	for i := uint32(0); i < numChunk; i++ {
		hash[i]  = make([]byte, 32)
	}

	b := make([]byte, 4)
	for i := uint32(0); i < numChunk; i++ {
		binary.LittleEndian.PutUint32(b, i)
		h := sha256.Sum256(slices.Concat(_chunk[i], b))
		copy(hash[i], h[:])
	}

	return hash
}

func GenMetadata(_param *PairingParam, _key *PairingKey, _chunk [][]byte) *Metadata {

	meta := new(Metadata)
	meta.Size = uint32(len(_chunk))
	meta.Hash = HashChunk(_chunk)
	meta.Tags = make([]*pbc.Element, meta.Size)

	for i := uint32(0); i < meta.Size; i++ {
		e1 := _param.SetFromHash(meta.Hash[i])
		e2 := _param.SetFromHash(_chunk[i])
		e3 := _param.PowBig(_param.U, e2.X())
		e4 := _param.Mul(e1, e3)
		meta.Tags[i] = _param.PowZn(e4, _key.PrivateKey)
	}

	return meta
}