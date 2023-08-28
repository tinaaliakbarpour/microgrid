package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// DeviceKeyPrefix is the prefix to retrieve all Device
	DeviceKeyPrefix = "Device/value/"
)

// DeviceKey returns the store key to retrieve a Device from the index fields
func DeviceKey(
	gridId uint64,
) []byte {
	var key []byte

	gridIdBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(gridIdBytes, gridId)
	key = append(key, gridIdBytes...)
	key = append(key, []byte("/")...)

	return key
}
