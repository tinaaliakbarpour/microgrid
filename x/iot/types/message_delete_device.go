package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeleteDevice = "delete_device"

var _ sdk.Msg = &MsgDeleteDevice{}

func NewMsgDeleteDevice(creator string, gridId uint64, address string) *MsgDeleteDevice {
	return &MsgDeleteDevice{
		Creator: creator,
		GridId:  gridId,
		Address: address,
	}
}

func (msg *MsgDeleteDevice) Route() string {
	return RouterKey
}

func (msg *MsgDeleteDevice) Type() string {
	return TypeMsgDeleteDevice
}

func (msg *MsgDeleteDevice) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteDevice) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteDevice) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
