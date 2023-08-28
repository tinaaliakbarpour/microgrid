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

func CmdDeleteDevice() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-device [grid-id] [address]",
		Short: "Broadcast message delete-device",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argGridId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argAddress := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteDevice(
				clientCtx.GetFromAddress().String(),
				argGridId,
				argAddress,
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
