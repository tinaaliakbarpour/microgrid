package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	"github.com/tinaaliakbarpour/microgrid/x/iot/types"
)

func CmdListDevice() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-device",
		Short: "list all device",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllDeviceRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.DeviceAll(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowDevice() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-device [grid-id]",
		Short: "shows a device",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			argGridId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			params := &types.QueryGetDeviceRequest{
				GridId: argGridId,
			}

			res, err := queryClient.Device(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
