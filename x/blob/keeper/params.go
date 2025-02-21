package keeper

import (
	"github.com/celestiaorg/celestia-app/v2/x/blob/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams gets all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.GasPerBlobByte(ctx),
		k.GovMaxSquareSize(ctx),
	)
}

// SetParams sets the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramStore.SetParamSet(ctx, &params)
}

// GasPerBlobByte returns the GasPerBlobByte param
func (k Keeper) GasPerBlobByte(ctx sdk.Context) (res uint32) {
	k.paramStore.Get(ctx, types.KeyGasPerBlobByte, &res)
	return res
}

// GovMaxSquareSize returns the GovMaxSquareSize param.
//
// Note: it is unlikely that you want to use this value directly because
// governance can modify it to be anything. Most use-cases will want to use
// squaresize.MaxEffective instead of GovMaxSquareSize.
func (k Keeper) GovMaxSquareSize(ctx sdk.Context) (res uint64) {
	k.paramStore.Get(ctx, types.KeyGovMaxSquareSize, &res)
	return res
}
