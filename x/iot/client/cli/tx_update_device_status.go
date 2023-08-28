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

func CmdUpdateDeviceStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-device-status [voltage] [power] [others] [addres] [grid-id]",
		Short: "Broadcast message update-device-status",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argVoltage, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argPower, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			argOthers := args[2]
			argAddres := args[3]
			argGridId, err := cast.ToUint64E(args[4])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateDeviceStatus(
				clientCtx.GetFromAddress().String(),
				argVoltage,
				argPower,
				argOthers,
				argAddres,
				argGridId,
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
