// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package test3

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

const (
	ScName        = "test3"
	ScDescription = "test3 description"
	HScName       = wasmlib.ScHname(0x1ffeb901)
)

const (
	ParamChargeWeight = "chargeWeight"
	ParamComp         = "comp"
	ParamFracID       = "fracID"
	ParamId           = "id"
	ParamMat          = "mat"
	ParamName         = "name"
	ParamOwner        = "owner"
	ParamPpID         = "ppID"
	ParamProdPass     = "prodPass"
	ParamProp         = "prop"
)

const (
	ResultCompositions    = "compositions"
	ResultId              = "id"
	ResultOwner           = "owner"
	ResultPpname          = "ppname"
	ResultPpresult        = "ppresult"
	ResultTokenPerPackage = "tokenPerPackage"
	ResultTokenRequired   = "tokenRequired"
)

const (
	StateCompositions     = "compositions"
	StateFracCompositions = "fracCompositions"
	StateFractions        = "fractions"
	StateOwner            = "owner"
	StatePricePerMg       = "pricePerMg"
	StateProductpasses    = "productpasses"
	StateShareRecycler    = "shareRecycler"
)

const (
	FuncAddMaterial              = "addMaterial"
	FuncAddPPToFraction          = "addPPToFraction"
	FuncCreateFraction           = "createFraction"
	FuncCreatePP                 = "createPP"
	FuncInit                     = "init"
	FuncSetMaterials             = "setMaterials"
	FuncSetOwner                 = "setOwner"
	ViewGetAmountOfRequiredFunds = "getAmountOfRequiredFunds"
	ViewGetMaterials             = "getMaterials"
	ViewGetOwner                 = "getOwner"
	ViewGetPP                    = "getPP"
	ViewGetTokenPerPackage       = "getTokenPerPackage"
)

const (
	HFuncAddMaterial              = wasmlib.ScHname(0xdfeea1ff)
	HFuncAddPPToFraction          = wasmlib.ScHname(0x50ac50a7)
	HFuncCreateFraction           = wasmlib.ScHname(0x59842fc3)
	HFuncCreatePP                 = wasmlib.ScHname(0x673fc3d7)
	HFuncInit                     = wasmlib.ScHname(0x1f44d644)
	HFuncSetMaterials             = wasmlib.ScHname(0x7f0ebcae)
	HFuncSetOwner                 = wasmlib.ScHname(0x2a15fe7b)
	HViewGetAmountOfRequiredFunds = wasmlib.ScHname(0xbcbaf2d6)
	HViewGetMaterials             = wasmlib.ScHname(0x1dca8ddb)
	HViewGetOwner                 = wasmlib.ScHname(0x137107a6)
	HViewGetPP                    = wasmlib.ScHname(0xbf711e08)
	HViewGetTokenPerPackage       = wasmlib.ScHname(0x522a14c0)
)