package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	legacy.RegisterAminoMsg(cdc, &MsgAddIBCRateLimit{}, "ratelimit/MsgAddIBCRateLimit")
	legacy.RegisterAminoMsg(cdc, &MsgUpdateIBCRateLimit{}, "ratelimit/MsgUpdateIBCRateLimit")
	legacy.RegisterAminoMsg(cdc, &MsgRemoveIBCRateLimit{}, "ratelimit/MsgRemoveIBCRateLimit")
	legacy.RegisterAminoMsg(cdc, &MsgResetIBCRateLimit{}, "ratelimit/MsgResetIBCRateLimit")
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddIBCRateLimit{},
		&MsgUpdateIBCRateLimit{},
		&MsgRemoveIBCRateLimit{},
		&MsgResetIBCRateLimit{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
