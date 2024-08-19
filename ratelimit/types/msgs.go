package types

import (
	"regexp"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
)

const (
	TypeMsgAddIBCRateLimit    = "AddRateLimit"
	TypeMsgUpdateIBCRateLimit = "UpdateRateLimit"
	TypeMsgRemoveIBCRateLimit = "RemoveRateLimit"
	TypeMsgResetIBCRateLimit  = "ResetRateLimit"
)

var (
	_ sdk.Msg = &MsgAddIBCRateLimit{}
	_ sdk.Msg = &MsgUpdateIBCRateLimit{}
	_ sdk.Msg = &MsgRemoveIBCRateLimit{}
	_ sdk.Msg = &MsgResetIBCRateLimit{}

	// Implement legacy interface for ledger support
	_ legacytx.LegacyMsg = &MsgAddIBCRateLimit{}
	_ legacytx.LegacyMsg = &MsgUpdateIBCRateLimit{}
	_ legacytx.LegacyMsg = &MsgRemoveIBCRateLimit{}
	_ legacytx.LegacyMsg = &MsgResetIBCRateLimit{}
)

// ----------------------------------------------
//               MsgAddIBCRateLimit
// ----------------------------------------------

func NewMsgAddIBCRateLimit(denom, channelId string, maxPercentSend sdkmath.Int, maxPercentRecv sdkmath.Int, durationHours uint64) *MsgAddIBCRateLimit {
	return &MsgAddIBCRateLimit{
		Denom:          denom,
		ChannelId:      channelId,
		MaxPercentSend: maxPercentSend,
		MaxPercentRecv: maxPercentRecv,
		DurationHours:  durationHours,
	}
}

func (msg MsgAddIBCRateLimit) Type() string {
	return TypeMsgAddIBCRateLimit
}

func (msg MsgAddIBCRateLimit) Route() string {
	return RouterKey
}

func (msg *MsgAddIBCRateLimit) GetSigners() []sdk.AccAddress {
	staker, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{staker}
}

func (msg *MsgAddIBCRateLimit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddIBCRateLimit) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authority address (%s)", err)
	}

	if msg.Denom == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid denom (%s)", msg.Denom)
	}

	matched, err := regexp.MatchString(`^channel-\d+$`, msg.ChannelId)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "unable to verify channel-id (%s)", msg.ChannelId)
	}
	if !matched {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest,
			"invalid channel-id (%s), must be of the format 'channel-{N}'", msg.ChannelId)
	}

	if msg.MaxPercentSend.GT(sdkmath.NewInt(100)) || msg.MaxPercentSend.LT(sdkmath.ZeroInt()) {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest,
			"max-percent-send percent must be between 0 and 100 (inclusively), Provided: %v", msg.MaxPercentSend)
	}

	if msg.MaxPercentRecv.GT(sdkmath.NewInt(100)) || msg.MaxPercentRecv.LT(sdkmath.ZeroInt()) {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest,
			"max-percent-recv percent must be between 0 and 100 (inclusively), Provided: %v", msg.MaxPercentRecv)
	}

	if msg.MaxPercentRecv.IsZero() && msg.MaxPercentSend.IsZero() {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest,
			"either the max send or max receive threshold must be greater than 0")
	}

	if msg.DurationHours == 0 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "duration can not be zero")
	}

	return nil
}

// ----------------------------------------------
//               MsgUpdateIBCRateLimit
// ----------------------------------------------

func NewMsgUpdateIBCRateLimit(denom, channelId string, maxPercentSend sdkmath.Int, maxPercentRecv sdkmath.Int, durationHours uint64) *MsgUpdateIBCRateLimit {
	return &MsgUpdateIBCRateLimit{
		Denom:          denom,
		ChannelId:      channelId,
		MaxPercentSend: maxPercentSend,
		MaxPercentRecv: maxPercentRecv,
		DurationHours:  durationHours,
	}
}

func (msg MsgUpdateIBCRateLimit) Type() string {
	return TypeMsgUpdateIBCRateLimit
}

func (msg MsgUpdateIBCRateLimit) Route() string {
	return RouterKey
}

func (msg *MsgUpdateIBCRateLimit) GetSigners() []sdk.AccAddress {
	staker, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{staker}
}

func (msg *MsgUpdateIBCRateLimit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateIBCRateLimit) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authority address (%s)", err)
	}

	if msg.Denom == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid denom (%s)", msg.Denom)
	}

	matched, err := regexp.MatchString(`^channel-\d+$`, msg.ChannelId)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "unable to verify channel-id (%s)", msg.ChannelId)
	}
	if !matched {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest,
			"invalid channel-id (%s), must be of the format 'channel-{N}'", msg.ChannelId)
	}

	if msg.MaxPercentSend.GT(sdkmath.NewInt(100)) || msg.MaxPercentSend.LT(sdkmath.ZeroInt()) {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest,
			"max-percent-send percent must be between 0 and 100 (inclusively), Provided: %v", msg.MaxPercentSend)
	}

	if msg.MaxPercentRecv.GT(sdkmath.NewInt(100)) || msg.MaxPercentRecv.LT(sdkmath.ZeroInt()) {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest,
			"max-percent-recv percent must be between 0 and 100 (inclusively), Provided: %v", msg.MaxPercentRecv)
	}

	if msg.MaxPercentRecv.IsZero() && msg.MaxPercentSend.IsZero() {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest,
			"either the max send or max receive threshold must be greater than 0")
	}

	if msg.DurationHours == 0 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "duration can not be zero")
	}

	return nil
}

// ----------------------------------------------
//               MsgRemoveIBCRateLimit
// ----------------------------------------------

func NewMsgRemoveIBCRateLimit(denom, channelId string) *MsgRemoveIBCRateLimit {
	return &MsgRemoveIBCRateLimit{
		Denom:     denom,
		ChannelId: channelId,
	}
}

func (msg MsgRemoveIBCRateLimit) Type() string {
	return TypeMsgRemoveIBCRateLimit
}

func (msg MsgRemoveIBCRateLimit) Route() string {
	return RouterKey
}

func (msg *MsgRemoveIBCRateLimit) GetSigners() []sdk.AccAddress {
	staker, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{staker}
}

func (msg *MsgRemoveIBCRateLimit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveIBCRateLimit) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authority address (%s)", err)
	}

	if msg.Denom == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid denom (%s)", msg.Denom)
	}

	matched, err := regexp.MatchString(`^channel-\d+$`, msg.ChannelId)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "unable to verify channel-id (%s)", msg.ChannelId)
	}
	if !matched {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest,
			"invalid channel-id (%s), must be of the format 'channel-{N}'", msg.ChannelId)
	}

	return nil
}

// ----------------------------------------------
//               MsgResetIBCRateLimit
// ----------------------------------------------

func NewMsgResetIBCRateLimit(denom, channelId string) *MsgResetIBCRateLimit {
	return &MsgResetIBCRateLimit{
		Denom:     denom,
		ChannelId: channelId,
	}
}

func (msg MsgResetIBCRateLimit) Type() string {
	return TypeMsgResetIBCRateLimit
}

func (msg MsgResetIBCRateLimit) Route() string {
	return RouterKey
}

func (msg *MsgResetIBCRateLimit) GetSigners() []sdk.AccAddress {
	staker, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{staker}
}

func (msg *MsgResetIBCRateLimit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgResetIBCRateLimit) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authority address (%s)", err)
	}

	if msg.Denom == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid denom (%s)", msg.Denom)
	}

	matched, err := regexp.MatchString(`^channel-\d+$`, msg.ChannelId)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "unable to verify channel-id (%s)", msg.ChannelId)
	}
	if !matched {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest,
			"invalid channel-id (%s), must be of the format 'channel-{N}'", msg.ChannelId)
	}

	return nil
}
