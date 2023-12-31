package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/ipgit2860/vcoa/x/vnft/keeper"
	"github.com/ipgit2860/vcoa/x/vnft/types"
)

func SimulateMsgMintNft(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgMintNft{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the MintNft simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "MintNft simulation not implemented"), nil, nil
	}
}
