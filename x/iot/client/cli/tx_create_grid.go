package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/tinaaliakbarpour/microgrid/x/iot/types"
)

var _ = strconv.Itoa(0)

func CmdCreateGrid() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-grid [name] [center-lat] [center-lon] [side]",
		Short: "Broadcast message create-grid",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argCenterLat, err := cast.ToInt32E(args[1])
			if err != nil {
				return err
			}
			argCenterLon, err := cast.ToInt32E(args[2])
			if err != nil {
				return err
			}
			argSide, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateGrid(
				clientCtx.GetFromAddress().String(),
				argName,
				argCenterLat,
				argCenterLon,
				argSide,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
