package router

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v3/modules/core/05-port/types"
	"github.com/strangelove-ventures/noble-router/x/router/keeper"
)

var _ porttypes.ICS4Wrapper = keeper.Keeper{}

type IBCMiddleware struct {
	app         porttypes.IBCModule
	keeper      *keeper.Keeper
	ics4Wrapper porttypes.ICS4Wrapper
}

// NewIBCMiddleware creates a new IBCMiddleware given the keeper and underlying application.
func NewIBCMiddleware(app porttypes.IBCModule, k *keeper.Keeper, ics4Wrapper porttypes.ICS4Wrapper) IBCMiddleware {
	return IBCMiddleware{
		app:         app,
		keeper:      k,
		ics4Wrapper: ics4Wrapper,
	}
}

// OnAcknowledgementPacket implements the IBCModule interface.
func (im IBCMiddleware) OnAcknowledgementPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	inFlightPacket := im.keeper.RetrieveInFlightPacket(ctx, packet.SourceChannel, packet.SourcePort, packet.Sequence)
	if inFlightPacket != nil {
		var ack channeltypes.Acknowledgement
		if err := channeltypes.SubModuleCdc.UnmarshalJSON(acknowledgement, &ack); err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal ICS-20 transfer packet acknowledgement: %v", err)
		}

		if ack.Success() {
			im.keeper.RemoveInFlightPacket(ctx, packet)
			im.keeper.DeleteInFlightPacket(ctx, inFlightPacket.SourceDomainSender, inFlightPacket.Nonce) // TODO probably don't track these twice
			im.keeper.DeleteMint(ctx, inFlightPacket.SourceDomainSender, inFlightPacket.Nonce)
			im.keeper.DeleteIBCForward(ctx, inFlightPacket.SourceDomainSender, inFlightPacket.Nonce)

		} else { // error on destination
			im.keeper.RemoveInFlightPacket(ctx, packet)
			im.keeper.DeleteInFlightPacket(ctx, inFlightPacket.SourceDomainSender, inFlightPacket.Nonce) // TODO probably don't track these twice

			// keep mint and mark IBCForward to indicate ack error for retry for future replaceDepositForBurnWithMetadata
			existingIBCForward, found := im.keeper.GetIBCForward(ctx, inFlightPacket.SourceDomainSender, inFlightPacket.Nonce)
			if found {
				existingIBCForward.AckError = true
				im.keeper.SetIBCForward(ctx, existingIBCForward)
			}

		}
	}

	return im.app.OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer)
}

// OnTimeoutPacket implements the IBCModule interface.
func (im IBCMiddleware) OnTimeoutPacket(ctx sdk.Context, packet channeltypes.Packet, relayer sdk.AccAddress) error {

	var data transfertypes.FungibleTokenPacketData
	if err := transfertypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		im.keeper.Logger(ctx).Error("packetForwardMiddleware error parsing packet data from timeout packet",
			"sequence", packet.Sequence,
			"src-channel", packet.SourceChannel, "src-port", packet.SourcePort,
			"dst-channel", packet.DestinationChannel, "dst-port", packet.DestinationPort,
			"error", err,
		)
		return im.app.OnTimeoutPacket(ctx, packet, relayer)
	}

	inFlightPacket, err := im.keeper.TimeoutShouldRetry(ctx, packet)

	if inFlightPacket != nil {
		if err != nil {
			im.keeper.RemoveInFlightPacket(ctx, packet)
			im.keeper.DeleteInFlightPacket(ctx, inFlightPacket.SourceDomainSender, inFlightPacket.Nonce) // TODO probably don't track these twice
			// this is a forwarded packet, so override handling to avoid refund from being processed on this chain.
			// WriteAcknowledgement with proxied ack to return success/fail to previous chain.
			// TODO
			return im.keeper.WriteAcknowledgementForForwardedPacket(ctx, packet, data, inFlightPacket, channeltypes.NewErrorAcknowledgement(err.Error()))
		}
		// timeout should be retried. In order to do that, we need to handle this timeout to refund on this chain first.
		if err := im.app.OnTimeoutPacket(ctx, packet, relayer); err != nil {
			return err
		}
		// TODO
		return im.keeper.RetryTimeout(ctx, packet.SourceChannel, packet.SourcePort, data, inFlightPacket)
	}

	// if inflight packet exists for chan/port/seq
	//      handle timeout (call downstream OnTimeoutPacket) to refund funds on the chain, so they can be sent again.
	//		sendpacket, update inflightpacket (already implemented in PFM in keeper.RetryTimeout)

	return im.app.OnTimeoutPacket(ctx, packet, relayer)
}
