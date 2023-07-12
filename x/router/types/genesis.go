package types

import (
	"fmt"
)

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		InFlightPackets: []InFlightPacket{},
		Mints:           []Mint{},
		IbcForwards:     []IBCForward{},
		Params:          DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {

	// Check for duplicated index in denoms
	inFlightPacketsIndexMap := make(map[string]struct{})
	for _, elem := range gs.InFlightPackets {
		index := string(InFlightPacket([]byte(elem.Denom)))
		if _, ok := inFlightPacketsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for denom")
		}
		inFlightPacketsIndexMap[index] = struct{}{}
	}

	// Check for duplicated index in mints
	mintsIndexMap := make(map[string]struct{})
	for _, elem := range gs.Mints {
		index := string(Mint([]byte(elem.Attestation)))
		if _, ok := mintsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for attestation")
		}
		mintsIndexMap[index] = struct{}{}
	}

	// Check for duplicated index in ibcForwards
	ibcForwardsIndexMap := make(map[string]struct{})
	for _, elem := range gs.IbcForwards {
		index := string(IBCForward(elem.SourceDomainSender))
		if _, ok := ibcForwardsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for ibcForward")
		}
		ibcForwardsIndexMap[index] = struct{}{}
	}

	return gs.Params.Validate()
}
