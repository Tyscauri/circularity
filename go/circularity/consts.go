// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package test3

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

const (
	ScName        = "test3"
	ScDescription = "test3 description"
	HScName       = wasmtypes.ScHname(0x1ffeb901)
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
	ResultFracID          = "fracID"
	ResultFraction        = "fraction"
	ResultId              = "id"
	ResultOwner           = "owner"
	ResultPp              = "pp"
	ResultPpname          = "ppname"
	ResultPpresult        = "ppresult"
	ResultRecyclateID     = "recyclateID"
	ResultTokenPerPackage = "tokenPerPackage"
	ResultTokenRequired   = "tokenRequired"
)

const (
	StateCompositions     = "compositions"
	StateFracCompositions = "fracCompositions"
	StateFractions        = "fractions"
	StateLastPayout       = "lastPayout"
	StateOwner            = "owner"
	StatePayoffKeysFrac   = "payoffKeysFrac"
	StatePayoffsFrac      = "payoffsFrac"
	StatePricePerMg       = "pricePerMg"
	StateProductpasses    = "productpasses"
	StateRecyCompositions = "recyCompositions"
	StateRecyclates       = "recyclates"
	StateShareRecycler    = "shareRecycler"
)

const (
	FuncAddMaterial              = "addMaterial"
	FuncAddPPToFraction          = "addPPToFraction"
	FuncCreateFraction           = "createFraction"
	FuncCreatePP                 = "createPP"
	FuncCreateRecyclate          = "createRecyclate"
	FuncInit                     = "init"
	FuncPayoutFrac               = "payoutFrac"
	FuncSetMaterials             = "setMaterials"
	FuncSetOwner                 = "setOwner"
	ViewGetAmountOfRequiredFunds = "getAmountOfRequiredFunds"
	ViewGetMaterials             = "getMaterials"
	ViewGetOwner                 = "getOwner"
	ViewGetPP                    = "getPP"
	ViewGetTokenPerPackage       = "getTokenPerPackage"
)

const (
	HFuncAddMaterial              = wasmtypes.ScHname(0xdfeea1ff)
	HFuncAddPPToFraction          = wasmtypes.ScHname(0x50ac50a7)
	HFuncCreateFraction           = wasmtypes.ScHname(0x59842fc3)
	HFuncCreatePP                 = wasmtypes.ScHname(0x673fc3d7)
	HFuncCreateRecyclate          = wasmtypes.ScHname(0x5066d840)
	HFuncInit                     = wasmtypes.ScHname(0x1f44d644)
	HFuncPayoutFrac               = wasmtypes.ScHname(0x363bbfc5)
	HFuncSetMaterials             = wasmtypes.ScHname(0x7f0ebcae)
	HFuncSetOwner                 = wasmtypes.ScHname(0x2a15fe7b)
	HViewGetAmountOfRequiredFunds = wasmtypes.ScHname(0xbcbaf2d6)
	HViewGetMaterials             = wasmtypes.ScHname(0x1dca8ddb)
	HViewGetOwner                 = wasmtypes.ScHname(0x137107a6)
	HViewGetPP                    = wasmtypes.ScHname(0xbf711e08)
	HViewGetTokenPerPackage       = wasmtypes.ScHname(0x522a14c0)
)
