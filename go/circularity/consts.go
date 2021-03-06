// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package circularity

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

const (
	ScName        = "circularity"
	ScDescription = "iota based circularity approach for plastic packaging, project former name was test3"
	HScName       = wasmtypes.ScHname(0x74fbdd8f)
)

const (
	ParamChargeWeight = "chargeWeight"
	ParamComp         = "comp"
	ParamExpiryDate   = "expiryDate"
	ParamFracID       = "fracID"
	ParamId           = "id"
	ParamName         = "name"
	ParamOwner        = "owner"
	ParamPpID         = "ppID"
	ParamProdPass     = "prodPass"
	ParamPurpose      = "purpose"
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
	StateDonationAddress  = "donationAddress"
	StateFracCompositions = "fracCompositions"
	StateFractions        = "fractions"
	StateOwner            = "owner"
	StatePricePerMg       = "pricePerMg"
	StateProductpasses    = "productpasses"
	StateRecyCompositions = "recyCompositions"
	StateRecyclates       = "recyclates"
	StateShareRecycler    = "shareRecycler"
	StateTokenToDonate    = "tokenToDonate"
)

const (
	FuncAddPPToFraction          = "addPPToFraction"
	FuncCreateFraction           = "createFraction"
	FuncCreatePP                 = "createPP"
	FuncCreateRecyclate          = "createRecyclate"
	FuncDeletePP                 = "deletePP"
	FuncInit                     = "init"
	FuncPayoutProducer           = "payoutProducer"
	FuncSetMaterials             = "setMaterials"
	FuncSetOwner                 = "setOwner"
	ViewGetAmountOfRequiredFunds = "getAmountOfRequiredFunds"
	ViewGetMaterials             = "getMaterials"
	ViewGetOwner                 = "getOwner"
	ViewGetPP                    = "getPP"
	ViewGetTokenPerPackage       = "getTokenPerPackage"
)

const (
	HFuncAddPPToFraction          = wasmtypes.ScHname(0x50ac50a7)
	HFuncCreateFraction           = wasmtypes.ScHname(0x59842fc3)
	HFuncCreatePP                 = wasmtypes.ScHname(0x673fc3d7)
	HFuncCreateRecyclate          = wasmtypes.ScHname(0x5066d840)
	HFuncDeletePP                 = wasmtypes.ScHname(0x56dedc36)
	HFuncInit                     = wasmtypes.ScHname(0x1f44d644)
	HFuncPayoutProducer           = wasmtypes.ScHname(0x3a56494b)
	HFuncSetMaterials             = wasmtypes.ScHname(0x7f0ebcae)
	HFuncSetOwner                 = wasmtypes.ScHname(0x2a15fe7b)
	HViewGetAmountOfRequiredFunds = wasmtypes.ScHname(0xbcbaf2d6)
	HViewGetMaterials             = wasmtypes.ScHname(0x1dca8ddb)
	HViewGetOwner                 = wasmtypes.ScHname(0x137107a6)
	HViewGetPP                    = wasmtypes.ScHname(0xbf711e08)
	HViewGetTokenPerPackage       = wasmtypes.ScHname(0x522a14c0)
)
