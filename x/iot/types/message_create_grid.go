package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateGrid = "create_grid"

var _ sdk.Msg = &MsgCreateGrid{}

func NewMsgCreateGrid(creator string, name string, centerLat int32, centerLon int32, side uint64) *MsgCreateGrid {
	return &MsgCreateGrid{
		Creator:   creator,
		Name:      name,
		CenterLat: centerLat,
		CenterLon: centerLon,
		Side:      side,
	}
}

func (msg *MsgCreateGrid) Route() string {
	return RouterKey
}

func (msg *MsgCreateGrid) Type() string {
	return TypeMsgCreateGrid
}

func (msg *MsgCreateGrid) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateGrid) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateGrid) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	//todo : we can add basic validation for type create grid on lat lon and side and also name
	return nil
}
