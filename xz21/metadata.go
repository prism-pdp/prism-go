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
	Chunk [][]byte // m
	Hash  [][]byte // M
	Tag   []*pbc.Element // gamma
}

func SplitData(_data []byte, _chunkSize int) ([][]byte, error) {
	if _chunkSize <= 0 { return nil, fmt.Errorf("invalid split num: %d", _chunkSize) }

	var splitedData [][]byte
	dataSize := len(_data)
	rows := int(math.Ceil(float64(dataSize) / float64(_chunkSize)))

	var s, e, processedSize int

	processedSize = 0
	for i := 0; i < rows; i++ {
		if i == (rows - 1) {
			s = i * _chunkSize
			e = s + (dataSize - processedSize)
		} else {
			s = i * _chunkSize
			e = s + _chunkSize
		}
		splitedData = append(splitedData, _data[s:e])
		processedSize += (e - s)
	}

	return splitedData, nil
}

// https://github.com/es3ku/z22m2azuma/blob/main/user/src/interfaces/crypt/content.go#L23
func GenMetadata(_param *PairingParam, _key *PairingKey, _splitedData [][]byte) *Metadata {
	numChunk := uint32(len(_splitedData))

	// metadata := make([]Metadata, numChunk)
	var metadata Metadata
	metadata.Hash = make([][]byte, numChunk)
	metadata.Chunk = make([][]byte, numChunk)
	metadata.Tag = make([]*pbc.Element, numChunk)

	for i := uint32(0); i < numChunk; i++ {
		metadata.Hash[i]  = make([]byte, 32)
		metadata.Chunk[i] = make([]byte, len(_splitedData[i]))

		copy(metadata.Chunk[i], _splitedData[i])
	}

	b := make([]byte, 4)
	for i := uint32(0); i < numChunk; i++ {
		binary.LittleEndian.PutUint32(b, i)
		hash := sha256.Sum256(slices.Concat(metadata.Chunk[i], b))
		copy(metadata.Hash[i], hash[:])
	}

	metadata.GenTag(_param, _key)

	return &metadata
}

func (this *Metadata) GenTag(_param *PairingParam, _key *PairingKey) {
	for i := range this.Tag {
		e1 := _param.SetFromHash(this.Hash[i])
		e2 := _param.SetFromHash(this.Chunk[i])
		e3 := _param.PowBig(_param.U, e2.X())
		e4 := _param.Mul(e1, e3)
		this.Tag[i] = _param.PowZn(e4, _key.PrivateKey)
	}
}