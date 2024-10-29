package xz21

type DigestSet struct {
	blocks map[uint32][]byte
	TotalNum uint32
}

func NewDigestSet(_totalNum uint32) *DigestSet {
	this := DigestSet{}
	this.blocks = make(map[uint32][]byte)
	this.TotalNum = _totalNum
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
