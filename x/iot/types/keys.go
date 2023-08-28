package types

const (
	// ModuleName defines the module name
	ModuleName = "iot"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_iot"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	GridKey      = "Grid/value/"
	GridCountKey = "Grid/count/"
)
