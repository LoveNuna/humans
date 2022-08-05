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

func CmdKeysignVote() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "keysign-vote [tx-hash] [pub-key]",
		Short: "Broadcast message keysignVote",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTxHash := args[0]
			argPubKey := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgKeysignVote(
				clientCtx.GetFromAddress().String(),
				argTxHash,
				argPubKey,
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
