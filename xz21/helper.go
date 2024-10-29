package xz21

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