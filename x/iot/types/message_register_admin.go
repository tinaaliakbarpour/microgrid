package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRegisterAdmin = "register_admin"

var _ sdk.Msg = &MsgRegisterAdmin{}

func NewMsgRegisterAdmin(creator string, id uint64, address string) *MsgRegisterAdmin {
	return &MsgRegisterAdmin{
		Creator: creator,
		Id:      id,
		Address: address,
	}
}

func (msg *MsgRegisterAdmin) Route() string {
	return RouterKey
}

func (msg *MsgRegisterAdmin) Type() string {
	return TypeMsgRegisterAdmin
}

func (msg *MsgRegisterAdmin) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegisterAdmin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegisterAdmin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
