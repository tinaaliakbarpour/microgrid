package types

import (
	"encoding/binary"
)

var _ binary.ByteOrder

const (
	// DeviceKeyPrefix is the prefix to retrieve all Device
	DeviceKeyPrefix = "Device/value/"
)

// DeviceKey returns the store key to retrieve a Device from the index fields
func DeviceKey(
	gridId uint64,
	address string,
) []byte {
	var key []byte

	gridIdBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(gridIdBytes, gridId)

	key = append(key, gridIdBytes...)
	key = append(key, []byte("/")...)

	key = append(key, []byte(address)...)
	key = append(key, []byte("/")...)

	return key
}

// DevicesKeyByGridId returns a key that returns all the devices belongs to a specific grid id
func DevicesKeyByGridId(gridID uint64) []byte {
	var key []byte

	gridIdBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(gridIdBytes, gridID)

	key = append(key, KeyPrefix(DeviceKeyPrefix)...)
	key = append(key, gridIdBytes...)

	key = append(key, []byte("/")...)
	return key
}
