// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package circularity

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type MapHashToImmutableCompositions struct {
	proxy wasmtypes.Proxy
}

func (m MapHashToImmutableCompositions) GetCompositions(key wasmtypes.ScHash) ImmutableCompositions {
	return ImmutableCompositions{proxy: m.proxy.Key(wasmtypes.HashToBytes(key))}
}

type MapHashToImmutableFracCompositions struct {
	proxy wasmtypes.Proxy
}

func (m MapHashToImmutableFracCompositions) GetFracCompositions(key wasmtypes.ScHash) ImmutableFracCompositions {
	return ImmutableFracCompositions{proxy: m.proxy.Key(wasmtypes.HashToBytes(key))}
}

type MapHashToImmutableFraction struct {
	proxy wasmtypes.Proxy
}

func (m MapHashToImmutableFraction) GetFraction(key wasmtypes.ScHash) ImmutableFraction {
	return ImmutableFraction{proxy: m.proxy.Key(wasmtypes.HashToBytes(key))}
}

type MapHashToImmutableProductPass struct {
	proxy wasmtypes.Proxy
}

func (m MapHashToImmutableProductPass) GetProductPass(key wasmtypes.ScHash) ImmutableProductPass {
	return ImmutableProductPass{proxy: m.proxy.Key(wasmtypes.HashToBytes(key))}
}

type MapHashToImmutableRecyCompositions struct {
	proxy wasmtypes.Proxy
}

func (m MapHashToImmutableRecyCompositions) GetRecyCompositions(key wasmtypes.ScHash) ImmutableRecyCompositions {
	return ImmutableRecyCompositions{proxy: m.proxy.Key(wasmtypes.HashToBytes(key))}
}

type MapHashToImmutableRecyclate struct {
	proxy wasmtypes.Proxy
}

func (m MapHashToImmutableRecyclate) GetRecyclate(key wasmtypes.ScHash) ImmutableRecyclate {
	return ImmutableRecyclate{proxy: m.proxy.Key(wasmtypes.HashToBytes(key))}
}

type ImmutablecircularityState struct {
	proxy wasmtypes.Proxy
}

func (s ImmutablecircularityState) Compositions() MapHashToImmutableCompositions {
	return MapHashToImmutableCompositions{proxy: s.proxy.Root(StateCompositions)}
}

func (s ImmutablecircularityState) DonationAddress() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(StateDonationAddress))
}

func (s ImmutablecircularityState) FracCompositions() MapHashToImmutableFracCompositions {
	return MapHashToImmutableFracCompositions{proxy: s.proxy.Root(StateFracCompositions)}
}

func (s ImmutablecircularityState) Fractions() MapHashToImmutableFraction {
	return MapHashToImmutableFraction{proxy: s.proxy.Root(StateFractions)}
}

func (s ImmutablecircularityState) Owner() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(StateOwner))
}

func (s ImmutablecircularityState) PricePerMg() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.proxy.Root(StatePricePerMg))
}

func (s ImmutablecircularityState) Productpasses() MapHashToImmutableProductPass {
	return MapHashToImmutableProductPass{proxy: s.proxy.Root(StateProductpasses)}
}

func (s ImmutablecircularityState) RecyCompositions() MapHashToImmutableRecyCompositions {
	return MapHashToImmutableRecyCompositions{proxy: s.proxy.Root(StateRecyCompositions)}
}

func (s ImmutablecircularityState) Recyclates() MapHashToImmutableRecyclate {
	return MapHashToImmutableRecyclate{proxy: s.proxy.Root(StateRecyclates)}
}

func (s ImmutablecircularityState) ShareRecycler() wasmtypes.ScImmutableUint8 {
	return wasmtypes.NewScImmutableUint8(s.proxy.Root(StateShareRecycler))
}

func (s ImmutablecircularityState) TokenToDonate() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.proxy.Root(StateTokenToDonate))
}

type MapHashToMutableCompositions struct {
	proxy wasmtypes.Proxy
}

func (m MapHashToMutableCompositions) Clear() {
	m.proxy.ClearMap()
}

func (m MapHashToMutableCompositions) GetCompositions(key wasmtypes.ScHash) MutableCompositions {
	return MutableCompositions{proxy: m.proxy.Key(wasmtypes.HashToBytes(key))}
}

type MapHashToMutableFracCompositions struct {
	proxy wasmtypes.Proxy
}

func (m MapHashToMutableFracCompositions) Clear() {
	m.proxy.ClearMap()
}

func (m MapHashToMutableFracCompositions) GetFracCompositions(key wasmtypes.ScHash) MutableFracCompositions {
	return MutableFracCompositions{proxy: m.proxy.Key(wasmtypes.HashToBytes(key))}
}

type MapHashToMutableFraction struct {
	proxy wasmtypes.Proxy
}

func (m MapHashToMutableFraction) Clear() {
	m.proxy.ClearMap()
}

func (m MapHashToMutableFraction) GetFraction(key wasmtypes.ScHash) MutableFraction {
	return MutableFraction{proxy: m.proxy.Key(wasmtypes.HashToBytes(key))}
}

type MapHashToMutableProductPass struct {
	proxy wasmtypes.Proxy
}

func (m MapHashToMutableProductPass) Clear() {
	m.proxy.ClearMap()
}

func (m MapHashToMutableProductPass) GetProductPass(key wasmtypes.ScHash) MutableProductPass {
	return MutableProductPass{proxy: m.proxy.Key(wasmtypes.HashToBytes(key))}
}

type MapHashToMutableRecyCompositions struct {
	proxy wasmtypes.Proxy
}

func (m MapHashToMutableRecyCompositions) Clear() {
	m.proxy.ClearMap()
}

func (m MapHashToMutableRecyCompositions) GetRecyCompositions(key wasmtypes.ScHash) MutableRecyCompositions {
	return MutableRecyCompositions{proxy: m.proxy.Key(wasmtypes.HashToBytes(key))}
}

type MapHashToMutableRecyclate struct {
	proxy wasmtypes.Proxy
}

func (m MapHashToMutableRecyclate) Clear() {
	m.proxy.ClearMap()
}

func (m MapHashToMutableRecyclate) GetRecyclate(key wasmtypes.ScHash) MutableRecyclate {
	return MutableRecyclate{proxy: m.proxy.Key(wasmtypes.HashToBytes(key))}
}

type MutablecircularityState struct {
	proxy wasmtypes.Proxy
}

func (s MutablecircularityState) AsImmutable() ImmutablecircularityState {
	return ImmutablecircularityState(s)
}

func (s MutablecircularityState) Compositions() MapHashToMutableCompositions {
	return MapHashToMutableCompositions{proxy: s.proxy.Root(StateCompositions)}
}

func (s MutablecircularityState) DonationAddress() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(StateDonationAddress))
}

func (s MutablecircularityState) FracCompositions() MapHashToMutableFracCompositions {
	return MapHashToMutableFracCompositions{proxy: s.proxy.Root(StateFracCompositions)}
}

func (s MutablecircularityState) Fractions() MapHashToMutableFraction {
	return MapHashToMutableFraction{proxy: s.proxy.Root(StateFractions)}
}

func (s MutablecircularityState) Owner() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(StateOwner))
}

func (s MutablecircularityState) PricePerMg() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.proxy.Root(StatePricePerMg))
}

func (s MutablecircularityState) Productpasses() MapHashToMutableProductPass {
	return MapHashToMutableProductPass{proxy: s.proxy.Root(StateProductpasses)}
}

func (s MutablecircularityState) RecyCompositions() MapHashToMutableRecyCompositions {
	return MapHashToMutableRecyCompositions{proxy: s.proxy.Root(StateRecyCompositions)}
}

func (s MutablecircularityState) Recyclates() MapHashToMutableRecyclate {
	return MapHashToMutableRecyclate{proxy: s.proxy.Root(StateRecyclates)}
}

func (s MutablecircularityState) ShareRecycler() wasmtypes.ScMutableUint8 {
	return wasmtypes.NewScMutableUint8(s.proxy.Root(StateShareRecycler))
}

func (s MutablecircularityState) TokenToDonate() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.proxy.Root(StateTokenToDonate))
}
