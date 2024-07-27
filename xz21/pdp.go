package xz21

import (
	"fmt"
)

func Upload(_param *PairingParam, _key *PairingKey, _data []byte, _chunkSize int) (*Metadata, error) {
	splitedData, err := SplitData(_data, _chunkSize)
	if err != nil { return nil, fmt.Errorf("Failure: SplitData") }

	metadata := GenMetadata(_param, _key, splitedData)

	return metadata, nil
}

func DedupChallen() error {
	return nil
}