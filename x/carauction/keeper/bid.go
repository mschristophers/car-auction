package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"car-auction/x/carauction/types"
)

func (k Keeper) GetBidCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BidCountKey))

	byteKey := []byte(types.BidCountKey)

	bz := store.Get(byteKey)

	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetBidCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BidCountKey))

	byteKey := []byte(types.BidCountKey)

	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)

	store.Set(byteKey, bz)
}

func (k Keeper) AppendBid(ctx sdk.Context, bid types.Bid) uint64 {
	count := k.GetBidCount(ctx)

	bid.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BidKey))

	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, bid.Id)

	appendedValue := k.cdc.MustMarshal(&bid)

	store.Set(byteKey, appendedValue)

	k.SetBidCount(ctx, count+1)
	return count
}
