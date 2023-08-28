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

func CmdCreateDevice() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-device [grid-id] [lat] [lon] [voltage] [power] [others]",
		Short: "Broadcast message create-device",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argGridId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argLat, err := cast.ToInt32E(args[1])
			if err != nil {
				return err
			}
			argLon, err := cast.ToInt32E(args[2])
			if err != nil {
				return err
			}
			argVoltage, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}
			argPower, err := cast.ToUint64E(args[4])
			if err != nil {
				return err
			}
			argOthers := args[5]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateDevice(
				clientCtx.GetFromAddress().String(),
				argGridId,
				argLat,
				argLon,
				argVoltage,
				argPower,
				argOthers,
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
