package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/strangelove-ventures/noble-router/x/router/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetInFlightPacket sets a InFlightPacket in the store
func (k Keeper) SetInFlightPacket(ctx sdk.Context, key types.InFlightPacket) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&key)
	store.Set(types.LookupKey(key.SourceDomainSender, key.Nonce), b)
}

// GetInFlightPacket returns InFlightPacket
func (k Keeper) GetInFlightPacket(ctx sdk.Context, sourceContractAddress string, nonce uint64) (val types.InFlightPacket, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.InFlightPacketPrefix(types.InFlightPacketKeyPrefix))

	b := store.Get(types.LookupKey(sourceContractAddress, nonce))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// DeleteInFlightPacket removes a InFlightPacket from the store
func (k Keeper) DeleteInFlightPacket(ctx sdk.Context, sourceContractAddress string, nonce uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.InFlightPacketPrefix(types.InFlightPacketKeyPrefix))
	store.Delete(types.InFlightPacketPrefix(string(types.LookupKey(sourceContractAddress, nonce))))
}

// GetAllInFlightPackets returns all InFlightPackets
func (k Keeper) GetAllInFlightPackets(ctx sdk.Context) (list []types.InFlightPacket) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.InFlightPacketPrefix(types.InFlightPacketKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.InFlightPacket
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
