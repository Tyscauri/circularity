// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package test3

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

type MapHashToImmutableCompositions struct {
	objID int32
}

func (m MapHashToImmutableCompositions) GetCompositions(key wasmlib.ScHash) ImmutableCompositions {
	subID := wasmlib.GetObjectID(m.objID, key.KeyID(), wasmlib.TYPE_ARRAY|wasmlib.TYPE_BYTES)
	return ImmutableCompositions{objID: subID}
}

type MapHashToImmutableFracCompositions struct {
	objID int32
}

func (m MapHashToImmutableFracCompositions) GetFracCompositions(key wasmlib.ScHash) ImmutableFracCompositions {
	subID := wasmlib.GetObjectID(m.objID, key.KeyID(), wasmlib.TYPE_ARRAY|wasmlib.TYPE_BYTES)
	return ImmutableFracCompositions{objID: subID}
}

type MapHashToImmutableFraction struct {
	objID int32
}

func (m MapHashToImmutableFraction) GetFraction(key wasmlib.ScHash) ImmutableFraction {
	return ImmutableFraction{objID: m.objID, keyID: key.KeyID()}
}

type MapHashToImmutableProductPass struct {
	objID int32
}

func (m MapHashToImmutableProductPass) GetProductPass(key wasmlib.ScHash) ImmutableProductPass {
	return ImmutableProductPass{objID: m.objID, keyID: key.KeyID()}
}

type Immutabletest3State struct {
	id int32
}

func (s Immutabletest3State) Compositions() MapHashToImmutableCompositions {
	mapID := wasmlib.GetObjectID(s.id, idxMap[IdxStateCompositions], wasmlib.TYPE_MAP)
	return MapHashToImmutableCompositions{objID: mapID}
}

func (s Immutabletest3State) FracCompositions() MapHashToImmutableFracCompositions {
	mapID := wasmlib.GetObjectID(s.id, idxMap[IdxStateFracCompositions], wasmlib.TYPE_MAP)
	return MapHashToImmutableFracCompositions{objID: mapID}
}

func (s Immutabletest3State) Fractions() MapHashToImmutableFraction {
	mapID := wasmlib.GetObjectID(s.id, idxMap[IdxStateFractions], wasmlib.TYPE_MAP)
	return MapHashToImmutableFraction{objID: mapID}
}

func (s Immutabletest3State) Owner() wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(s.id, idxMap[IdxStateOwner])
}

func (s Immutabletest3State) PricePerMg() wasmlib.ScImmutableUint64 {
	return wasmlib.NewScImmutableUint64(s.id, idxMap[IdxStatePricePerMg])
}

func (s Immutabletest3State) Productpasses() MapHashToImmutableProductPass {
	mapID := wasmlib.GetObjectID(s.id, idxMap[IdxStateProductpasses], wasmlib.TYPE_MAP)
	return MapHashToImmutableProductPass{objID: mapID}
}

func (s Immutabletest3State) ShareRecycler() wasmlib.ScImmutableUint8 {
	return wasmlib.NewScImmutableUint8(s.id, idxMap[IdxStateShareRecycler])
}

type MapHashToMutableCompositions struct {
	objID int32
}

func (m MapHashToMutableCompositions) Clear() {
	wasmlib.Clear(m.objID)
}

func (m MapHashToMutableCompositions) GetCompositions(key wasmlib.ScHash) MutableCompositions {
	subID := wasmlib.GetObjectID(m.objID, key.KeyID(), wasmlib.TYPE_ARRAY|wasmlib.TYPE_BYTES)
	return MutableCompositions{objID: subID}
}

type MapHashToMutableFracCompositions struct {
	objID int32
}

func (m MapHashToMutableFracCompositions) Clear() {
	wasmlib.Clear(m.objID)
}

func (m MapHashToMutableFracCompositions) GetFracCompositions(key wasmlib.ScHash) MutableFracCompositions {
	subID := wasmlib.GetObjectID(m.objID, key.KeyID(), wasmlib.TYPE_ARRAY|wasmlib.TYPE_BYTES)
	return MutableFracCompositions{objID: subID}
}

type MapHashToMutableFraction struct {
	objID int32
}

func (m MapHashToMutableFraction) Clear() {
	wasmlib.Clear(m.objID)
}

func (m MapHashToMutableFraction) GetFraction(key wasmlib.ScHash) MutableFraction {
	return MutableFraction{objID: m.objID, keyID: key.KeyID()}
}

type MapHashToMutableProductPass struct {
	objID int32
}

func (m MapHashToMutableProductPass) Clear() {
	wasmlib.Clear(m.objID)
}

func (m MapHashToMutableProductPass) GetProductPass(key wasmlib.ScHash) MutableProductPass {
	return MutableProductPass{objID: m.objID, keyID: key.KeyID()}
}

type Mutabletest3State struct {
	id int32
}

func (s Mutabletest3State) Compositions() MapHashToMutableCompositions {
	mapID := wasmlib.GetObjectID(s.id, idxMap[IdxStateCompositions], wasmlib.TYPE_MAP)
	return MapHashToMutableCompositions{objID: mapID}
}

func (s Mutabletest3State) FracCompositions() MapHashToMutableFracCompositions {
	mapID := wasmlib.GetObjectID(s.id, idxMap[IdxStateFracCompositions], wasmlib.TYPE_MAP)
	return MapHashToMutableFracCompositions{objID: mapID}
}

func (s Mutabletest3State) Fractions() MapHashToMutableFraction {
	mapID := wasmlib.GetObjectID(s.id, idxMap[IdxStateFractions], wasmlib.TYPE_MAP)
	return MapHashToMutableFraction{objID: mapID}
}

func (s Mutabletest3State) Owner() wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(s.id, idxMap[IdxStateOwner])
}

func (s Mutabletest3State) PricePerMg() wasmlib.ScMutableUint64 {
	return wasmlib.NewScMutableUint64(s.id, idxMap[IdxStatePricePerMg])
}

func (s Mutabletest3State) Productpasses() MapHashToMutableProductPass {
	mapID := wasmlib.GetObjectID(s.id, idxMap[IdxStateProductpasses], wasmlib.TYPE_MAP)
	return MapHashToMutableProductPass{objID: mapID}
}

func (s Mutabletest3State) ShareRecycler() wasmlib.ScMutableUint8 {
	return wasmlib.NewScMutableUint8(s.id, idxMap[IdxStateShareRecycler])
}
