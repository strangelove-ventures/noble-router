package keeper

import (
	"fmt"
	porttypes "github.com/cosmos/ibc-go/v3/modules/core/05-port/types"
	"github.com/strangelove-ventures/noble-router/x/router/types"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		cdc         codec.BinaryCodec
		storeKey    storetypes.StoreKey
		paramstore  paramtypes.Subspace
		cctp        types.CctpKeeper
		ics4Wrapper porttypes.ICS4Wrapper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	cctpfactory types.CctpKeeper,
	ics4Wrapper porttypes.ICS4Wrapper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:         cdc,
		storeKey:    storeKey,
		paramstore:  ps,
		cctp:        cctpfactory,
		ics4Wrapper: ics4Wrapper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// TODO delete
//// SendPacket implements the ICS4Wrapper interface.
//func (k Keeper) SendPacket(
//	ctx sdk.Context,
//	chanCap *capabilitytypes.Capability,
//	packet exported.PacketI,
//) error {
//	chanPacket, ok := packet.(chantypes.Packet)
//	if !ok {
//		// not channel packet, forward to next middleware
//		return k.ics4Wrapper.SendPacket(ctx, chanCap, packet)
//	}
//
//	var data transfertypes.FungibleTokenPacketData
//	if err := transfertypes.ModuleCdc.UnmarshalJSON(chanPacket.Data, &data); err != nil {
//		// not fungible token packet data, forward to next middleware
//		return k.ics4Wrapper.SendPacket(ctx, chanCap, packet)
//	}
//
//	newData, err := transfertypes.ModuleCdc.MarshalJSON(&data)
//	if err != nil {
//		return fmt.Errorf("failed to marshal new packet data: %w", err)
//	}
//
//	chanPacket.Data = newData
//
//	return k.ics4Wrapper.SendPacket(ctx, chanCap, chanPacket)
//}
//
//// WriteAcknowledgement implements the ICS4Wrapper interface.
//func (k Keeper) WriteAcknowledgement(
//	ctx sdk.Context,
//	chanCap *capabilitytypes.Capability,
//	packet exported.PacketI,
//	ack exported.Acknowledgement,
//) error {
//	return k.ics4Wrapper.WriteAcknowledgement(ctx, chanCap, packet, ack)
//}
