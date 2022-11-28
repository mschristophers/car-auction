package keeper

import (
	"encoding/binary"
	"errors"

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

func (k Keeper) AppendEndAuction(ctx sdk.Context, endAuction types.EndAuction) (uint64, uint64, error) {

	auctionStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.AuctionKey))

	byteAuctionId := make([]byte, 8)
	binary.BigEndian.PutUint64(byteAuctionId, endAuction.AuctionID)

	targetAuctionByte := auctionStore.Get(byteAuctionId)
	var targetAuction types.Auction

	if err := k.cdc.Unmarshal(targetAuctionByte, &targetAuction); err != nil {
		return 0, 0, err
	}

	if targetAuction.Creator != endAuction.Creator {
		return 0, 0, errors.New("It is not a target auction creator.")
	}

	if targetAuction.Ended {
		return 0, 0, errors.New("The target auction already ended.")
	}

	bidStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BidKey))

	byteBidId := make([]byte, 8)
	binary.BigEndian.PutUint64(byteBidId, endAuction.BidID)

	targetBidByte := bidStore.Get(byteBidId)
	var targetBid types.Bid

	if err := k.cdc.Unmarshal(targetBidByte, &targetBid); err != nil {
		return 0, 0, err
	}

	count := k.GetEndAuctionCount(ctx)

	endAuction.Id = count
	endAuction.FinalPrice = targetBid.BidPrice

	// Update target auction status to ended
	targetAuction.Ended = true
	updatedTargetAuction := k.cdc.MustMarshal(&targetAuction)
	auctionStore.Set(byteAuctionId, updatedTargetAuction)

	// Save the Endd auction
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.EndAuctionKey))

	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, endAuction.Id)

	appendedValue := k.cdc.MustMarshal(&endAuction)

	store.Set(byteKey, appendedValue)

	k.SetEndAuctionCount(ctx, count+1)
	return count, endAuction.FinalPrice, nil
}
