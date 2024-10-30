package xz21

type DigestSet map[uint32][]byte

func (this DigestSet) Base() map[uint32][]byte {
	return (map[uint32][]byte)(this)
}

func (this DigestSet) Set(_index uint32, _data []byte) {
	this[_index] = make([]byte, len(_data))
	copy(this[_index], _data)
}

func (this DigestSet) Size() uint32 {
	return uint32(len(this.Base()))
}
