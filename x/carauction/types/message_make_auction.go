package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMakeAuction = "make_auction"

var _ sdk.Msg = &MsgMakeAuction{}

func NewMsgMakeAuction(creator string, name string, initialPrice uint64, duration uint64) *MsgMakeAuction {
	return &MsgMakeAuction{
		Creator:      creator,
		Name:         name,
		InitialPrice: initialPrice,
		Duration:     duration,
	}
}

func (msg *MsgMakeAuction) Route() string {
	return RouterKey
}

func (msg *MsgMakeAuction) Type() string {
	return TypeMsgMakeAuction
}

func (msg *MsgMakeAuction) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMakeAuction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMakeAuction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
