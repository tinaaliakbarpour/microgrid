package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tinaaliakbarpour/microgrid/x/iot/types"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				GridList: []types.Grid{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				GridCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated grid",
			genState: &types.GenesisState{
				GridList: []types.Grid{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid grid count",
			genState: &types.GenesisState{
				GridList: []types.Grid{
					{
						Id: 1,
					},
				},
				GridCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
