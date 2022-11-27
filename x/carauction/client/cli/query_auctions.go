package cli

import (
    "strconv"
	
	"github.com/spf13/cobra"
    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"car-auction/x/carauction/types"
)

var _ = strconv.Itoa(0)

func CmdAuctions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auctions [initial-price] [min-increment]",
		Short: "Query auctions",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			 reqInitialPrice := args[0]
			 reqMinIncrement := args[1]
			
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAuctionsRequest{
				
                InitialPrice: reqInitialPrice, 
                MinIncrement: reqMinIncrement, 
            }

            

			res, err := queryClient.Auctions(cmd.Context(), params)
            if err != nil {
                return err
            }

            return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

    return cmd
}
