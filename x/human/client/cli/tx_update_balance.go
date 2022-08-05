package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"human/x/human/types"
)

var _ = strconv.Itoa(0)

func CmdUpdateBalance() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-balance [chain-name] [balance] [decimal]",
		Short: "Broadcast message updateBalance",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argChainName := args[0]
			argBalance := args[1]
			argDecimal := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateBalance(
				clientCtx.GetFromAddress().String(),
				argChainName,
				argBalance,
				argDecimal,
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
