package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddBid = "add_bid"

var _ sdk.Msg = &MsgAddBid{}

func NewMsgAddBid(creator string, auctionID uint64, bidPrice uint64) *MsgAddBid {
	return &MsgAddBid{
		Creator:   creator,
		AuctionID: auctionID,
		BidPrice:  bidPrice,
	}
}

func (msg *MsgAddBid) Route() string {
	return RouterKey
}

func (msg *MsgAddBid) Type() string {
	return TypeMsgAddBid
}

func (msg *MsgAddBid) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddBid) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddBid) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
