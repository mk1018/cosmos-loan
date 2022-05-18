package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgTestLoan = "test_loan"

var _ sdk.Msg = &MsgTestLoan{}

func NewMsgTestLoan(creator string, id uint64) *MsgTestLoan {
	return &MsgTestLoan{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgTestLoan) Route() string {
	return RouterKey
}

func (msg *MsgTestLoan) Type() string {
	return TypeMsgTestLoan
}

func (msg *MsgTestLoan) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTestLoan) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTestLoan) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
