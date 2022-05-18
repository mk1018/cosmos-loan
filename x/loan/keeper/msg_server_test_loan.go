package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/username/loan/x/loan/types"
)

func (k msgServer) TestLoan(goCtx context.Context, msg *types.MsgTestLoan) (*types.MsgTestLoanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgTestLoanResponse{}, nil
}
