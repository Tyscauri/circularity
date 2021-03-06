// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package test3

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

type Composition struct {
	Material   string 
	Proportion uint8 
}

func NewCompositionFromBytes(bytes []byte) *Composition {
	decode := wasmlib.NewBytesDecoder(bytes)
	data := &Composition{}
	data.Material   = decode.String()
	data.Proportion = decode.Uint8()
	decode.Close()
	return data
}

func (o *Composition) Bytes() []byte {
	return wasmlib.NewBytesEncoder().
		String(o.Material).
		Uint8(o.Proportion).
		Data()
}

type ImmutableComposition struct {
	objID int32
	keyID wasmlib.Key32
}

func (o ImmutableComposition) Exists() bool {
	return wasmlib.Exists(o.objID, o.keyID, wasmlib.TYPE_BYTES)
}

func (o ImmutableComposition) Value() *Composition {
	return NewCompositionFromBytes(wasmlib.GetBytes(o.objID, o.keyID, wasmlib.TYPE_BYTES))
}

type MutableComposition struct {
	objID int32
	keyID wasmlib.Key32
}

func (o MutableComposition) Delete() {
	wasmlib.DelKey(o.objID, o.keyID, wasmlib.TYPE_BYTES)
}

func (o MutableComposition) Exists() bool {
	return wasmlib.Exists(o.objID, o.keyID, wasmlib.TYPE_BYTES)
}

func (o MutableComposition) SetValue(value *Composition) {
	wasmlib.SetBytes(o.objID, o.keyID, wasmlib.TYPE_BYTES, value.Bytes())
}

func (o MutableComposition) Value() *Composition {
	return NewCompositionFromBytes(wasmlib.GetBytes(o.objID, o.keyID, wasmlib.TYPE_BYTES))
}

type FracComposition struct {
	Material string 
	Weight   uint64  //in mg
}

func NewFracCompositionFromBytes(bytes []byte) *FracComposition {
	decode := wasmlib.NewBytesDecoder(bytes)
	data := &FracComposition{}
	data.Material = decode.String()
	data.Weight   = decode.Uint64()
	decode.Close()
	return data
}

func (o *FracComposition) Bytes() []byte {
	return wasmlib.NewBytesEncoder().
		String(o.Material).
		Uint64(o.Weight).
		Data()
}

type ImmutableFracComposition struct {
	objID int32
	keyID wasmlib.Key32
}

func (o ImmutableFracComposition) Exists() bool {
	return wasmlib.Exists(o.objID, o.keyID, wasmlib.TYPE_BYTES)
}

func (o ImmutableFracComposition) Value() *FracComposition {
	return NewFracCompositionFromBytes(wasmlib.GetBytes(o.objID, o.keyID, wasmlib.TYPE_BYTES))
}

type MutableFracComposition struct {
	objID int32
	keyID wasmlib.Key32
}

func (o MutableFracComposition) Delete() {
	wasmlib.DelKey(o.objID, o.keyID, wasmlib.TYPE_BYTES)
}

func (o MutableFracComposition) Exists() bool {
	return wasmlib.Exists(o.objID, o.keyID, wasmlib.TYPE_BYTES)
}

func (o MutableFracComposition) SetValue(value *FracComposition) {
	wasmlib.SetBytes(o.objID, o.keyID, wasmlib.TYPE_BYTES, value.Bytes())
}

func (o MutableFracComposition) Value() *FracComposition {
	return NewFracCompositionFromBytes(wasmlib.GetBytes(o.objID, o.keyID, wasmlib.TYPE_BYTES))
}

type Fraction struct {
	DecFood    bool 
	DecHygiene bool 
	Did        string 
	FracId     wasmlib.ScHash 
	Issuer     wasmlib.ScAgentID 
	Name       string 
}

func NewFractionFromBytes(bytes []byte) *Fraction {
	decode := wasmlib.NewBytesDecoder(bytes)
	data := &Fraction{}
	data.DecFood    = decode.Bool()
	data.DecHygiene = decode.Bool()
	data.Did        = decode.String()
	data.FracId     = decode.Hash()
	data.Issuer     = decode.AgentID()
	data.Name       = decode.String()
	decode.Close()
	return data
}

func (o *Fraction) Bytes() []byte {
	return wasmlib.NewBytesEncoder().
		Bool(o.DecFood).
		Bool(o.DecHygiene).
		String(o.Did).
		Hash(o.FracId).
		AgentID(o.Issuer).
		String(o.Name).
		Data()
}

type ImmutableFraction struct {
	objID int32
	keyID wasmlib.Key32
}

func (o ImmutableFraction) Exists() bool {
	return wasmlib.Exists(o.objID, o.keyID, wasmlib.TYPE_BYTES)
}

func (o ImmutableFraction) Value() *Fraction {
	return NewFractionFromBytes(wasmlib.GetBytes(o.objID, o.keyID, wasmlib.TYPE_BYTES))
}

type MutableFraction struct {
	objID int32
	keyID wasmlib.Key32
}

func (o MutableFraction) Delete() {
	wasmlib.DelKey(o.objID, o.keyID, wasmlib.TYPE_BYTES)
}

func (o MutableFraction) Exists() bool {
	return wasmlib.Exists(o.objID, o.keyID, wasmlib.TYPE_BYTES)
}

func (o MutableFraction) SetValue(value *Fraction) {
	wasmlib.SetBytes(o.objID, o.keyID, wasmlib.TYPE_BYTES, value.Bytes())
}

func (o MutableFraction) Value() *Fraction {
	return NewFractionFromBytes(wasmlib.GetBytes(o.objID, o.keyID, wasmlib.TYPE_BYTES))
}

type ProductPass struct {
	Amount        int64 
	ChargeWeight  uint64 
	DecFood       bool 
	DecHygiene    bool 
	Did           string  //merged did:iota:id
	Id            wasmlib.ScHash 
	Issuer        wasmlib.ScAgentID  //packaging producer
	Name          string 
	PackageWeight uint64 
	Version       uint8 
}

func NewProductPassFromBytes(bytes []byte) *ProductPass {
	decode := wasmlib.NewBytesDecoder(bytes)
	data := &ProductPass{}
	data.Amount        = decode.Int64()
	data.ChargeWeight  = decode.Uint64()
	data.DecFood       = decode.Bool()
	data.DecHygiene    = decode.Bool()
	data.Did           = decode.String()
	data.Id            = decode.Hash()
	data.Issuer        = decode.AgentID()
	data.Name          = decode.String()
	data.PackageWeight = decode.Uint64()
	data.Version       = decode.Uint8()
	decode.Close()
	return data
}

func (o *ProductPass) Bytes() []byte {
	return wasmlib.NewBytesEncoder().
		Int64(o.Amount).
		Uint64(o.ChargeWeight).
		Bool(o.DecFood).
		Bool(o.DecHygiene).
		String(o.Did).
		Hash(o.Id).
		AgentID(o.Issuer).
		String(o.Name).
		Uint64(o.PackageWeight).
		Uint8(o.Version).
		Data()
}

type ImmutableProductPass struct {
	objID int32
	keyID wasmlib.Key32
}

func (o ImmutableProductPass) Exists() bool {
	return wasmlib.Exists(o.objID, o.keyID, wasmlib.TYPE_BYTES)
}

func (o ImmutableProductPass) Value() *ProductPass {
	return NewProductPassFromBytes(wasmlib.GetBytes(o.objID, o.keyID, wasmlib.TYPE_BYTES))
}

type MutableProductPass struct {
	objID int32
	keyID wasmlib.Key32
}

func (o MutableProductPass) Delete() {
	wasmlib.DelKey(o.objID, o.keyID, wasmlib.TYPE_BYTES)
}

func (o MutableProductPass) Exists() bool {
	return wasmlib.Exists(o.objID, o.keyID, wasmlib.TYPE_BYTES)
}

func (o MutableProductPass) SetValue(value *ProductPass) {
	wasmlib.SetBytes(o.objID, o.keyID, wasmlib.TYPE_BYTES, value.Bytes())
}

func (o MutableProductPass) Value() *ProductPass {
	return NewProductPassFromBytes(wasmlib.GetBytes(o.objID, o.keyID, wasmlib.TYPE_BYTES))
}
