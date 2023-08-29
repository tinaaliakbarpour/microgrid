package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/tinaaliakbarpour/microgrid/x/iot/types"
)

var _ = strconv.Itoa(0)

func CmdListDeviceByGridId() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-device-by-grid-id [grid-id]",
		Short: "Query list-device-by-grid-id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqGridId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryListDeviceByGridIdRequest{

				GridId: reqGridId,
			}

			res, err := queryClient.ListDeviceByGridId(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
