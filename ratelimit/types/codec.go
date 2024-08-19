package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	legacy.RegisterAminoMsg(cdc, &MsgAddRateLimit{}, "ratelimit/MsgAddRateLimit")
	legacy.RegisterAminoMsg(cdc, &MsgUpdateRateLimit{}, "ratelimit/MsgUpdateRateLimit")
	legacy.RegisterAminoMsg(cdc, &MsgRemoveRateLimit{}, "ratelimit/MsgRemoveRateLimit")
	legacy.RegisterAminoMsg(cdc, &MsgResetRateLimit{}, "ratelimit/MsgResetRateLimit")
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddRateLimit{},
		&MsgUpdateRateLimit{},
		&MsgRemoveRateLimit{},
		&MsgResetRateLimit{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
