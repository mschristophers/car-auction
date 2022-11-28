package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"car-auction/x/carauction/types"
)

func (k Keeper) GetEndAuctionCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.EndAuctionCountKey))

	byteKey := []byte(types.EndAuctionCountKey)

	bz := store.Get(byteKey)

	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetEndAuctionCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.EndAuctionCountKey))

	byteKey := []byte(types.EndAuctionCountKey)

	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)

	store.Set(byteKey, bz)
}

func (k Keeper) AppendEndAuction(ctx sdk.Context, endAuction types.EndAuction) (uint64, string, error) {
	count := k.GetEndAuctionCount(ctx)

	endAuction.Id = count

	// Save the End auction
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.EndAuctionKey))

	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, endAuction.Id)

	appendedValue := k.cdc.MustMarshal(&endAuction)

	store.Set(byteKey, appendedValue)

	k.SetEndAuctionCount(ctx, count+1)
	return count, endAuction.HammerPrice, nil
}
