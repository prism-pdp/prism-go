package xz21

func MakeSubset(_data []byte, _tagDataSet *TagDataSet, _chal *Chal) (*DigestSet, *TagDataSet, error) {
	chunkSet := GenChunkSet(_data, _tagDataSet.Size)

	numChunk := chunkSet.Size()
	setA := GenA(_chal, numChunk)

	digestSubSet := chunkSet.HashByIndex(setA)

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

func MakeList[T uint32](_num T) []T {
	list := make([]T, _num)
	for i := T(0); i < _num; i++ {
		list[i] = i
	}
	return list
}