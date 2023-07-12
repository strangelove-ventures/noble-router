package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/strangelove-ventures/noble-router/x/router/types"
)

func CmdListInFlightPackets() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-in-flight-packets",
		Short: "lists all InFlightPackets",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllInFlightPacketsRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.InFlightPackets(context.Background(), params)
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

func CmdShowInFlightPacket() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-in-flight-packet [source-contract-address] [nonce]",
		Short: "shows an InFlightPacket",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			sourceContractAddress := args[0]
			nonceRaw := args[1]
			nonce, err := strconv.ParseUint(nonceRaw, 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetInFlightPacketRequest{
				SourceContractAddress: sourceContractAddress,
				Nonce:                 nonce,
			}

			res, err := queryClient.InFlightPacket(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
