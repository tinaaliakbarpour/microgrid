package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeleteGrid = "delete_grid"

var _ sdk.Msg = &MsgDeleteGrid{}

func NewMsgDeleteGrid(creator string, id uint64) *MsgDeleteGrid {
	return &MsgDeleteGrid{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgDeleteGrid) Route() string {
	return RouterKey
}

func (msg *MsgDeleteGrid) Type() string {
	return TypeMsgDeleteGrid
}

func (msg *MsgDeleteGrid) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteGrid) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteGrid) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
