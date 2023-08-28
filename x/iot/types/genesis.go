package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		GridList:   []Grid{},
		DeviceList: []Device{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in grid
	gridIdMap := make(map[uint64]bool)
	gridCount := gs.GetGridCount()
	for _, elem := range gs.GridList {
		if _, ok := gridIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for grid")
		}
		if elem.Id >= gridCount {
			return fmt.Errorf("grid id should be lower or equal than the last id")
		}
		gridIdMap[elem.Id] = true
	}
	// Check for duplicated index in device
	deviceIndexMap := make(map[string]struct{})

	for _, elem := range gs.DeviceList {
		index := string(DeviceKey(elem.GridId))
		if _, ok := deviceIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for device")
		}
		deviceIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
