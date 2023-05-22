package devearn

import (
	"sidechain/x/devearn/keeper"
	"sidechain/x/devearn/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the assets
	for _, elem := range genState.AssetsList {
		k.SetAssets(ctx, elem)
	}

	// Set assets count
	k.SetAssetsCount(ctx, genState.AssetsCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.DevEarnInfos = k.GetAllDevEarnInfos(ctx)
	genesis.AssetsList = k.GetAllAssets(ctx)
	genesis.AssetsCount = k.GetAssetsCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
