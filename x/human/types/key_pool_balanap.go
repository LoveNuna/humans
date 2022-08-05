package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// PoolBalanapKeyPrefix is the prefix to retrieve all PoolBalanap
	PoolBalanapKeyPrefix = "PoolBalanap/value/"
)

// PoolBalanapKey returns the store key to retrieve a PoolBalanap from the index fields
func PoolBalanapKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
