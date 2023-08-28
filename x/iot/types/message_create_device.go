package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateDevice = "create_device"

var _ sdk.Msg = &MsgCreateDevice{}

func NewMsgCreateDevice(creator string, gridId uint64, lat int32, lon int32, voltage uint64, power uint64, others string) *MsgCreateDevice {
	return &MsgCreateDevice{
		Creator: creator,
		GridId:  gridId,
		Lat:     lat,
		Lon:     lon,
		Voltage: voltage,
		Power:   power,
		Others:  others,
	}
}

func (msg *MsgCreateDevice) Route() string {
	return RouterKey
}

func (msg *MsgCreateDevice) Type() string {
	return TypeMsgCreateDevice
}

func (msg *MsgCreateDevice) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateDevice) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateDevice) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
