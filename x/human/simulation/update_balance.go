package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"human/x/human/keeper"
	"human/x/human/types"
)

func SimulateMsgUpdateBalance(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgUpdateBalance{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the UpdateBalance simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "UpdateBalance simulation not implemented"), nil, nil
	}
}
