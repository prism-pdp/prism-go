package xz21

func SplitData(_data []byte, _chunkNum uint32) (*ChunkSet, error) {
	chunkSet := NewChunkSet()
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
		// chunkSet = append(chunkSet, _data[s:e])
		chunkSet.Set(i, _data[s:e])
	}

	return chunkSet, nil
}

func MakeSubset(_data []byte, _tagDataSet *TagDataSet, _chal *Chal) (*DigestSet, *TagDataSet, error) {
	chunkSet, err := SplitData(_data, _tagDataSet.Size)
	if err != nil { return nil, nil, err}

	numChunk := chunkSet.Size()
	setA := GenA(_chal, numChunk)

	digestSubSet := HashChunks(chunkSet, setA)

	tagDataSubset := NewTagDataSet()
	tagDataSubset.Size = _tagDataSet.Size
	for _, i := range setA {
		if _, ok := tagDataSubset.TagData[i]; !ok {
			tagDataSubset.TagData[i] = _tagDataSet.TagData[i]
		}
	}

	return digestSubSet, tagDataSubset, nil
}

func MapKeys[K comparable, V any](m map[K]V) []K {
    result := make([]K, 0, len(m))
    for key := range m {
        result = append(result, key)
    }
    return result
}