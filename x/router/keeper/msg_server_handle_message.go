package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	"github.com/strangelove-ventures/noble-router/x/router/types"
	"strconv"
)

func (k Keeper) HandleMessage(goCtx context.Context, msg []byte) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// parse outer message
	outerMessage, err := decodeMessage(msg)
	if err != nil {
		return err
	}

	// parse internal message into IBCForward
	ibcForward, err := decodeIBCForward(outerMessage.MessageBody)
	if err == nil { // message is an IBC forward

		existingIBCForward, existingIBCForwardFound := k.GetIBCForward(ctx, string(outerMessage.Sender), outerMessage.Nonce)
		if existingIBCForwardFound {
			if ibcForward.AckError {
				_, existingMintFound := k.GetMint(ctx, string(outerMessage.Sender), outerMessage.Nonce)
				if existingMintFound {
					_, existingInFlightPacketFound := k.GetInFlightPacket(ctx, string(outerMessage.Sender), outerMessage.Nonce)
					if existingInFlightPacketFound {
						panic("unexpected state")
					} else {
						packet := buildPacket(existingIBCForward)
						err := k.SendPacket(ctx, nil, packet) // TODO add channel capabilities
						if err != nil {
							return err
						}

						inFlightPacket := buildInFlightPacket(outerMessage, ibcForward)

						k.SetInFlightPacket(ctx, inFlightPacket)
					}
				} else {
					panic("unexpected state")
				}
			} else {
				return sdkerrors.Wrapf(types.ErrHandleMessage, "previous operation still in progress")
			}
		} else {
			k.SetIBCForward(ctx, ibcForward)
			_, existingMintFound := k.GetMint(ctx, string(outerMessage.Sender), outerMessage.Nonce)
			if existingMintFound {
				_, existingInFlightPacketFound := k.GetInFlightPacket(ctx, string(outerMessage.Sender), outerMessage.Nonce)
				if existingInFlightPacketFound {
					panic("unexpected state")
				} else {
					packet := buildPacket(existingIBCForward)
					k.SendPacket(ctx, nil, packet) // TODO add channel capabilities

					inFlightPacket := buildInFlightPacket(outerMessage, ibcForward)

					k.SetInFlightPacket(ctx, inFlightPacket)
				}
			}
		}
	} else {
		// try to parse internal message into burn (representing a remote burn -> local mint)
		burnMessage, err := decodeBurnMessage(outerMessage.MessageBody)
		if err == nil { // message is a Mint
			mint := types.Mint{
				SourceDomainSender: string(outerMessage.Sender),
				Nonce:              outerMessage.Nonce,
				Amount: &sdk.Coin{
					Denom:  string(burnMessage.BurnToken),
					Amount: sdk.NewInt(int64(burnMessage.Amount)),
				},
				DestinationDomain: strconv.Itoa(int(outerMessage.DestinationDomain)),
				MintRecipient:     string(burnMessage.MintRecipient),
			}
			k.SetMint(ctx, mint)
			existingIBCForward, found := k.GetIBCForward(ctx, string(burnMessage.MessageSender), outerMessage.Nonce)
			if found {
				_, found := k.GetInFlightPacket(ctx, mint.SourceDomainSender, mint.Nonce)
				if found {
					panic("unexpected state")
				} else {
					packet := buildPacket(existingIBCForward)
					k.SendPacket(ctx, nil, packet) // TODO add channel capabilities

					inFlightPacket := buildInFlightPacket(outerMessage, ibcForward)
					k.SetInFlightPacket(ctx, inFlightPacket)
				}
			}
		}
	}

	return err
}

func buildInFlightPacket(outerMessage Message, ibcForward types.IBCForward) types.InFlightPacket {
	inFlightPacket := types.InFlightPacket{
		SourceDomainSender:     string(outerMessage.Sender),
		Nonce:                  outerMessage.Nonce,
		OriginalSenderAddress:  string(outerMessage.Recipient),
		RefundChannelId:        ibcForward.Channel,
		RefundPortId:           ibcForward.Port,
		PacketSrcChannelId:     ibcForward.Channel,
		PacketSrcPortId:        ibcForward.Port,
		PacketTimeoutTimestamp: 0,  // TODO
		PacketTimeoutHeight:    "", // TODO
		PacketData:             []byte(ibcForward.Data),
		RefundSequence:         0, // TODO
		RetriesRemaining:       0, // TODO
		Timeout:                0, // TODO
		Nonrefundable:          false,
	}
	return inFlightPacket
}

func buildPacket(ibcForward types.IBCForward) channeltypes.Packet {
	packet := channeltypes.Packet{
		SourcePort:    ibcForward.Port,
		SourceChannel: ibcForward.Channel,
		Data:          []byte(ibcForward.Data),
	}
	return packet
}
