// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package test3

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

const (
	IdxParamChargeWeight = 0
	IdxParamComp         = 1
	IdxParamFracID       = 2
	IdxParamId           = 3
	IdxParamMat          = 4
	IdxParamName         = 5
	IdxParamOwner        = 6
	IdxParamPpID         = 7
	IdxParamProdPass     = 8
	IdxParamProp         = 9

	IdxResultCompositions    = 10
	IdxResultId              = 11
	IdxResultOwner           = 12
	IdxResultPpname          = 13
	IdxResultPpresult        = 14
	IdxResultTokenPerPackage = 15
	IdxResultTokenRequired   = 16

	IdxStateCompositions     = 17
	IdxStateFracCompositions = 18
	IdxStateFractions        = 19
	IdxStateOwner            = 20
	IdxStatePricePerMg       = 21
	IdxStateProductpasses    = 22
	IdxStateShareRecycler    = 23
)

const keyMapLen = 24

var keyMap = [keyMapLen]wasmlib.Key{
	ParamChargeWeight,
	ParamComp,
	ParamFracID,
	ParamId,
	ParamMat,
	ParamName,
	ParamOwner,
	ParamPpID,
	ParamProdPass,
	ParamProp,
	ResultCompositions,
	ResultId,
	ResultOwner,
	ResultPpname,
	ResultPpresult,
	ResultTokenPerPackage,
	ResultTokenRequired,
	StateCompositions,
	StateFracCompositions,
	StateFractions,
	StateOwner,
	StatePricePerMg,
	StateProductpasses,
	StateShareRecycler,
}

var idxMap [keyMapLen]wasmlib.Key32