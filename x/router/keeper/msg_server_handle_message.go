package keeper

import (
	"context"
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/strangelove-ventures/noble-router/x/router/types"
)

func (k Keeper) HandleMessage(goCtx context.Context, msg []byte) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// try to parse outer message
	outerMessage := decodeMessage(msg)

	// try to parse internal message into IBCForward
	var ibcForward types.IBCForward
	err := json.Unmarshal(outerMessage.MessageBody, &ibcForward)
	if err == nil { // message is an IBC forward
		existingIBCForward, found := k.GetIBCForward(ctx, ibcForward.SourceDomainSender, ibcForward.Nonce)
		if found {
			// if ack error
			if true { // TODO
				existingMint, found := k.GetMint(ctx, ibcForward.SourceDomainSender, ibcForward.Nonce)
				if found {
					_, found := k.GetInFlightPacket(ctx, ibcForward.SourceDomainSender, ibcForward.Nonce)
					if found {
						panic("unexpected state")
					} else {
						// TODO sendPacket

						// TODO
						inFlightPacket := types.InFlightPacket{
							SourceDomainSender:     "",
							Nonce:                  0,
							OriginalSenderAddress:  "",
							RefundChannelId:        "",
							RefundPortId:           "",
							PacketSrcChannelId:     "",
							PacketSrcPortId:        "",
							PacketTimeoutTimestamp: 0,
							PacketTimeoutHeight:    "",
							PacketData:             nil,
							RefundSequence:         0,
							RetriesRemaining:       0,
							Timeout:                0,
							Nonrefundable:          false,
						}
						k.SetInFlightPacket(ctx, inFlightPacket)
					}
				}
			} else {
				// error (previous operation still in progress)
			}
		}
	}

	// try to parse internal message into mint TODO else block?
	var mint types.Mint
	err = json.Unmarshal(outerMessage.MessageBody, &mint)
	if err == nil { // message is a Mint
		k.SetMint(ctx, mint)
		_, found := k.GetInFlightPacket(ctx, mint.SourceDomainSender, mint.Nonce)
		if found {
			// err
		} else {
			// TODO sendpacket

			// TODO
			inFlightPacket := types.InFlightPacket{
				SourceDomainSender:     "",
				Nonce:                  0,
				OriginalSenderAddress:  "",
				RefundChannelId:        "",
				RefundPortId:           "",
				PacketSrcChannelId:     "",
				PacketSrcPortId:        "",
				PacketTimeoutTimestamp: 0,
				PacketTimeoutHeight:    "",
				PacketData:             nil,
				RefundSequence:         0,
				RetriesRemaining:       0,
				Timeout:                0,
				Nonrefundable:          false,
			}
			k.SetInFlightPacket(ctx, inFlightPacket)
		}
	}

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
