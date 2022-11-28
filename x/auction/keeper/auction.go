package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"auction/x/auction/types"
)

func (k Keeper) GetAuction(ctx sdk.Context, id uint64) (val types.Auction, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuctionKey))

	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)

	b := store.Get(bz)
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) UpdateAuctionHighestBidID(ctx sdk.Context, id uint64, bidID uint64) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuctionKey))

	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)

	b := store.Get(bz)
	if b == nil {
		return types.AuctionNotFound
	}

	var auction types.Auction
	k.cdc.MustUnmarshal(b, &auction)

	auction.HighestBidIsPresent = true
	auction.CurrentHighestBidID = bidID

	appendedValue := k.cdc.MustMarshal(&auction)

	store.Set(bz, appendedValue)
	return nil
}

func (k Keeper) EndAuction(ctx sdk.Context, id uint64) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuctionKey))

	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)

	b := store.Get(bz)
	if b == nil {
		return types.AuctionNotFound
	}

	var auction types.Auction
	k.cdc.MustUnmarshal(b, &auction)
	auction.Ended = true

	appendedValue := k.cdc.MustMarshal(&auction)

	store.Set(bz, appendedValue)
	return nil
}

func (k Keeper) GetAuctionCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.AuctionCountKey))

	byteKey := []byte(types.AuctionCountKey)

	bz := store.Get(byteKey)

	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetAuctionCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.AuctionCountKey))

	byteKey := []byte(types.AuctionCountKey)

	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)

	store.Set(byteKey, bz)
}

func (k Keeper) AppendAuction(ctx sdk.Context, auction types.Auction) uint64 {
	count := k.GetAuctionCount(ctx)

	auction.Id = count
	auction.Ended = false

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.AuctionKey))

	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, auction.Id)

	appendedValue := k.cdc.MustMarshal(&auction)

	store.Set(byteKey, appendedValue)

	k.SetAuctionCount(ctx, count+1)
	return count
}
