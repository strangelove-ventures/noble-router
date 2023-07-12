package keeper

import (
	"context"
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/strangelove-ventures/noble-router/x/router/types"
)

// TODO
func (k Keeper) HandleMessage(goCtx context.Context, msg []byte, attestation []byte) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// try to parse internal message into IBCForward
	var ibcForward types.IBCForward
	err := json.Unmarshal([]byte(msg.Message), &ibcForward)

	if err == nil { // message is an IBC forward

		mint, found := k.GetMint(ctx, msg.Attestation)
		if found {
			// check for an inflight packet - look up by _______
			// if found, error

			err := k.ics4Wrapper.SendPacket(ctx, nil, nil)
			if err != nil {
				return nil, err
			}

		} else {
			k.SetIBCForward(ctx, ibcForward)

		}
	}

	// try to parse internal message into Mint
	var mint types.Mints
	err := json.Unmarshal([]byte(msg.Message), &mint)

	// if msg is forward
	//	// if existing ibc forward info
	//		// if ack error on ibc forward info
	//			// if mint info set
	//				// if inflightpacket set
	//				 	// panic - unexpected state
	//				// else
	//					// send packet, set inflightpacket for chan/port/seq
	//		// else
	//			// error (previous operation still in progress)
	//	// else
	//		// set forward info
	//		// if mint info set
	//			// if inflightpacket set
	//			 	// error
	//			// else
	//				// send packet, set inflightpacket for chan/port/seq
	// else if msg is mint
	//	// set mint info
	//	// if forward info set
	//		// if inflightpacket set
	//			// error
	//		// else
	//			// send packet, set inflightpacket for chan/port/seq

	// # in ibc middleware
	//
	// OnAcknowledgement
	//	// if inflightpacket exists for chan/port/seq
	// 		// if ack success
	//			// clear inflightpacket, mint, and forward (happy path, success)
	//		// else (error on destination)
	//			// clear inflightpacket
	//			// retain mint info (so that replaceDepositForBurnWithMetadata can come back through and retry IBC forward with new metadata)
	//			// update ibcforward info to indicate ack error (so that replaceDepositForBurnWithMetadata can come back through and retry IBC forward with new metadata)
	//
	// OnTimeoutPacket
	//     	// if inflightpacket exists for chan/port/seq
	//			// handle timeout (call downstream OnTimeoutPacket) to refund funds on the chain, so they can be sent again.
	//			// sendpacket, update inflightpacket (already impld in PFM in keeper.RetryTimeout)

	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgHandleMessageResponse{}, err
}
