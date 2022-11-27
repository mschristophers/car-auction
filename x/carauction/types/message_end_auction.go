package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgEndAuction = "end_auction"

var _ sdk.Msg = &MsgEndAuction{}

func NewMsgEndAuction(creator string, auctionID uint64, bidID uint64) *MsgEndAuction {
	return &MsgEndAuction{
		Creator:   creator,
		AuctionID: auctionID,
		BidID:     bidID,
	}
}

func (msg *MsgEndAuction) Route() string {
	return RouterKey
}

func (msg *MsgEndAuction) Type() string {
	return TypeMsgEndAuction
}

func (msg *MsgEndAuction) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEndAuction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEndAuction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
