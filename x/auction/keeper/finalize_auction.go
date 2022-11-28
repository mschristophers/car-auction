package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"auction/x/auction/types"
)

func (k Keeper) GetFinalizeAuctionCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.FinalizeAuctionCountKey))

	byteKey := []byte(types.FinalizeAuctionCountKey)

	bz := store.Get(byteKey)

	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetFinalizeAuctionCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.FinalizeAuctionCountKey))

	byteKey := []byte(types.FinalizeAuctionCountKey)

	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)

	store.Set(byteKey, bz)
}

func (k Keeper) AppendFinalizeAuction(ctx sdk.Context, finalizeAuction types.FinalizeAuction) (uint64, string, error) {

	count := k.GetFinalizeAuctionCount(ctx)

	finalizeAuction.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.FinalizeAuctionKey))

	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, finalizeAuction.Id)

	appendedValue := k.cdc.MustMarshal(&finalizeAuction)

	store.Set(byteKey, appendedValue)

	k.SetFinalizeAuctionCount(ctx, count+1)
	return count, finalizeAuction.FinalPrice, nil
}
