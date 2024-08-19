package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/Int3facechain/ibc-rate-limiting/ratelimit/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the ratelimit MsgServer interface
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// Adds a new rate limit. Fails if the rate limit already exists or the channel value is 0
func (k msgServer) AddIBCRateLimit(goCtx context.Context, msg *types.MsgAddIBCRateLimit) (*types.MsgAddIBCRateLimitResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if k.authority != msg.Authority {
		return nil, errorsmod.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.authority, msg.Authority)
	}

	if err := k.Keeper.AddRateLimit(ctx, msg); err != nil {
		return nil, err
	}

	return &types.MsgAddIBCRateLimitResponse{}, nil
}

// Updates an existing rate limit. Fails if the rate limit doesn't exist
func (k msgServer) UpdateIBCRateLimit(goCtx context.Context, msg *types.MsgUpdateIBCRateLimit) (*types.MsgUpdateIBCRateLimitResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if k.authority != msg.Authority {
		return nil, errorsmod.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.authority, msg.Authority)
	}

	if err := k.Keeper.UpdateRateLimit(ctx, msg); err != nil {
		return nil, err
	}

	return &types.MsgUpdateIBCRateLimitResponse{}, nil
}

// Removes a rate limit. Fails if the rate limit doesn't exist
func (k msgServer) RemoveIBCRateLimit(goCtx context.Context, msg *types.MsgRemoveIBCRateLimit) (*types.MsgRemoveIBCRateLimitResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if k.authority != msg.Authority {
		return nil, errorsmod.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.authority, msg.Authority)
	}

	_, found := k.Keeper.GetRateLimit(ctx, msg.Denom, msg.ChannelId)
	if !found {
		return nil, types.ErrRateLimitNotFound
	}

	k.Keeper.RemoveRateLimit(ctx, msg.Denom, msg.ChannelId)
	return &types.MsgRemoveIBCRateLimitResponse{}, nil
}

// Resets the flow on a rate limit. Fails if the rate limit doesn't exist
func (k msgServer) ResetIBCRateLimit(goCtx context.Context, msg *types.MsgResetIBCRateLimit) (*types.MsgResetIBCRateLimitResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if k.authority != msg.Authority {
		return nil, errorsmod.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.authority, msg.Authority)
	}

	if err := k.Keeper.ResetRateLimit(ctx, msg.Denom, msg.ChannelId); err != nil {
		return nil, err
	}

	return &types.MsgResetIBCRateLimitResponse{}, nil
}
