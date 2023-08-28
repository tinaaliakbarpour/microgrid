package cli

// import (
//     "strconv"

// 	 "github.com/spf13/cast"
// 	"github.com/spf13/cobra"
//     "github.com/cosmos/cosmos-sdk/client"
// 	"github.com/cosmos/cosmos-sdk/client/flags"
// 	"github.com/tinaaliakbarpour/microgrid/x/iot/types"
// )

// var _ = strconv.Itoa(0)

// func CmdShowDevice() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "show-device [grid-id] [address]",
// 		Short: "Query show-device",
// 		Args:  cobra.ExactArgs(2),
// 		RunE: func(cmd *cobra.Command, args []string) (err error) {
// 			 reqGridId, err := cast.ToUint64E(args[0])
//             		if err != nil {
//                 		return err
//             		}
// 			 reqAddress := args[1]

// 			clientCtx, err := client.GetClientQueryContext(cmd)
// 			if err != nil {
// 				return err
// 			}

// 			queryClient := types.NewQueryClient(clientCtx)

// 			params := &types.QueryShowDeviceRequest{

//                 GridId: reqGridId,
//                 Address: reqAddress,
//             }

// 			res, err := queryClient.ShowDevice(cmd.Context(), params)
//             if err != nil {
//                 return err
//             }

//             return clientCtx.PrintProto(res)
// 		},
// 	}

// 	flags.AddQueryFlagsToCmd(cmd)

//     return cmd
// }
