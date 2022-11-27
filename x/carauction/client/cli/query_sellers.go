package cli

import (
    "strconv"
	
	"github.com/spf13/cobra"
    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"car-auction/x/carauction/types"
)

var _ = strconv.Itoa(0)

func CmdSellers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sellers [initial-price] [min-increment]",
		Short: "Query sellers",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			 reqInitialPrice := args[0]
			 reqMinIncrement := args[1]
			
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QuerySellersRequest{
				
                InitialPrice: reqInitialPrice, 
                MinIncrement: reqMinIncrement, 
            }

            

			res, err := queryClient.Sellers(cmd.Context(), params)
            if err != nil {
                return err
            }

            return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

    return cmd
}
