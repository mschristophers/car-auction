package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"car-auction/x/carauction/types"
)

func (k Keeper) GetAuctionCount(ctx sdk.Context) uint64 {
	// Get the store using storeKey (aka "auction") and AuctionCountKey (aka "Auction/count/")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.AuctionCountKey))

	// Converts the AuctionCountKey to bytes
	byteKey := []byte(types.AuctionCountKey)

	// Gets the value of the count
	bz := store.Get(byteKey)

	// Returns zero if the count value can't be found (e.g., it's the first auction)
	if bz == nil {
		return 0
	}

	// Converts the count into uint64
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetAuctionCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.AuctionCountKey))

	byteKey := []byte(types.AuctionCountKey)

	// Converts coutn from uint64 to string and get bytes
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)

	// Sets the value of Auction/count/ to count
	store.Set(byteKey, bz)
}

func (k Keeper) AppendAuction(ctx sdk.Context, auction types.Auction) uint64 {
	count := k.GetAuctionCount(ctx)

	auction.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.AuctionKey))

	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, auction.Id)

	appendedValue := k.cdc.MustMarshal(&auction)

	store.Set(byteKey, appendedValue)

	k.SetAuctionCount(ctx, count+1)
	return count
}