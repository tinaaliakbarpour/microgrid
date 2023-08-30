package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateDeviceStatus = "update_device_status"

var _ sdk.Msg = &MsgUpdateDeviceStatus{}

func NewMsgUpdateDeviceStatus(creator string, voltage uint64, power uint64, others string, addres string, gridId uint64) *MsgUpdateDeviceStatus {
	return &MsgUpdateDeviceStatus{
		Creator: creator,
		Voltage: voltage,
		Power:   power,
		Others:  others,
		GridId:  gridId,
	}
}

func (msg *MsgUpdateDeviceStatus) Route() string {
	return RouterKey
}

func (msg *MsgUpdateDeviceStatus) Type() string {
	return TypeMsgUpdateDeviceStatus
}

func (msg *MsgUpdateDeviceStatus) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateDeviceStatus) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateDeviceStatus) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

//TODO: we can have all the validations and acl here
