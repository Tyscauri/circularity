// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package test3

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib"

type AddMaterialCall struct {
	Func    *wasmlib.ScFunc
	Params  MutableAddMaterialParams
}

type AddPPToFractionCall struct {
	Func    *wasmlib.ScFunc
	Params  MutableAddPPToFractionParams
	Results ImmutableAddPPToFractionResults
}

type CreateFractionCall struct {
	Func    *wasmlib.ScFunc
	Params  MutableCreateFractionParams
	Results ImmutableCreateFractionResults
}

type CreatePPCall struct {
	Func    *wasmlib.ScFunc
	Params  MutableCreatePPParams
	Results ImmutableCreatePPResults
}

type CreateRecyclateCall struct {
	Func    *wasmlib.ScFunc
	Params  MutableCreateRecyclateParams
	Results ImmutableCreateRecyclateResults
}

type InitCall struct {
	Func    *wasmlib.ScInitFunc
	Params  MutableInitParams
}

type PayoutFracCall struct {
	Func    *wasmlib.ScFunc
	Params  MutablePayoutFracParams
}

type SetMaterialsCall struct {
	Func    *wasmlib.ScFunc
	Params  MutableSetMaterialsParams
}

type SetOwnerCall struct {
	Func    *wasmlib.ScFunc
	Params  MutableSetOwnerParams
}

type GetAmountOfRequiredFundsCall struct {
	Func    *wasmlib.ScView
	Params  MutableGetAmountOfRequiredFundsParams
	Results ImmutableGetAmountOfRequiredFundsResults
}

type GetMaterialsCall struct {
	Func    *wasmlib.ScView
	Params  MutableGetMaterialsParams
	Results ImmutableGetMaterialsResults
}

type GetOwnerCall struct {
	Func    *wasmlib.ScView
	Results ImmutableGetOwnerResults
}

type GetPPCall struct {
	Func    *wasmlib.ScView
	Params  MutableGetPPParams
	Results ImmutableGetPPResults
}

type GetTokenPerPackageCall struct {
	Func    *wasmlib.ScView
	Params  MutableGetTokenPerPackageParams
	Results ImmutableGetTokenPerPackageResults
}

type Funcs struct{}

var ScFuncs Funcs

func (sc Funcs) AddMaterial(ctx wasmlib.ScFuncCallContext) *AddMaterialCall {
	f := &AddMaterialCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncAddMaterial)}
	f.Params.proxy = wasmlib.NewCallParamsProxy(&f.Func.ScView)
	return f
}

func (sc Funcs) AddPPToFraction(ctx wasmlib.ScFuncCallContext) *AddPPToFractionCall {
	f := &AddPPToFractionCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncAddPPToFraction)}
	f.Params.proxy = wasmlib.NewCallParamsProxy(&f.Func.ScView)
	wasmlib.NewCallResultsProxy(&f.Func.ScView, &f.Results.proxy)
	return f
}

func (sc Funcs) CreateFraction(ctx wasmlib.ScFuncCallContext) *CreateFractionCall {
	f := &CreateFractionCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncCreateFraction)}
	f.Params.proxy = wasmlib.NewCallParamsProxy(&f.Func.ScView)
	wasmlib.NewCallResultsProxy(&f.Func.ScView, &f.Results.proxy)
	return f
}

func (sc Funcs) CreatePP(ctx wasmlib.ScFuncCallContext) *CreatePPCall {
	f := &CreatePPCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncCreatePP)}
	f.Params.proxy = wasmlib.NewCallParamsProxy(&f.Func.ScView)
	wasmlib.NewCallResultsProxy(&f.Func.ScView, &f.Results.proxy)
	return f
}

func (sc Funcs) CreateRecyclate(ctx wasmlib.ScFuncCallContext) *CreateRecyclateCall {
	f := &CreateRecyclateCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncCreateRecyclate)}
	f.Params.proxy = wasmlib.NewCallParamsProxy(&f.Func.ScView)
	wasmlib.NewCallResultsProxy(&f.Func.ScView, &f.Results.proxy)
	return f
}

func (sc Funcs) Init(ctx wasmlib.ScFuncCallContext) *InitCall {
	f := &InitCall{Func: wasmlib.NewScInitFunc(ctx, HScName, HFuncInit)}
	f.Params.proxy = wasmlib.NewCallParamsProxy(&f.Func.ScView)
	return f
}

func (sc Funcs) PayoutFrac(ctx wasmlib.ScFuncCallContext) *PayoutFracCall {
	f := &PayoutFracCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncPayoutFrac)}
	f.Params.proxy = wasmlib.NewCallParamsProxy(&f.Func.ScView)
	return f
}

func (sc Funcs) SetMaterials(ctx wasmlib.ScFuncCallContext) *SetMaterialsCall {
	f := &SetMaterialsCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncSetMaterials)}
	f.Params.proxy = wasmlib.NewCallParamsProxy(&f.Func.ScView)
	return f
}

func (sc Funcs) SetOwner(ctx wasmlib.ScFuncCallContext) *SetOwnerCall {
	f := &SetOwnerCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncSetOwner)}
	f.Params.proxy = wasmlib.NewCallParamsProxy(&f.Func.ScView)
	return f
}

func (sc Funcs) GetAmountOfRequiredFunds(ctx wasmlib.ScViewCallContext) *GetAmountOfRequiredFundsCall {
	f := &GetAmountOfRequiredFundsCall{Func: wasmlib.NewScView(ctx, HScName, HViewGetAmountOfRequiredFunds)}
	f.Params.proxy = wasmlib.NewCallParamsProxy(f.Func)
	wasmlib.NewCallResultsProxy(f.Func, &f.Results.proxy)
	return f
}

func (sc Funcs) GetMaterials(ctx wasmlib.ScViewCallContext) *GetMaterialsCall {
	f := &GetMaterialsCall{Func: wasmlib.NewScView(ctx, HScName, HViewGetMaterials)}
	f.Params.proxy = wasmlib.NewCallParamsProxy(f.Func)
	wasmlib.NewCallResultsProxy(f.Func, &f.Results.proxy)
	return f
}

func (sc Funcs) GetOwner(ctx wasmlib.ScViewCallContext) *GetOwnerCall {
	f := &GetOwnerCall{Func: wasmlib.NewScView(ctx, HScName, HViewGetOwner)}
	wasmlib.NewCallResultsProxy(f.Func, &f.Results.proxy)
	return f
}

func (sc Funcs) GetPP(ctx wasmlib.ScViewCallContext) *GetPPCall {
	f := &GetPPCall{Func: wasmlib.NewScView(ctx, HScName, HViewGetPP)}
	f.Params.proxy = wasmlib.NewCallParamsProxy(f.Func)
	wasmlib.NewCallResultsProxy(f.Func, &f.Results.proxy)
	return f
}

func (sc Funcs) GetTokenPerPackage(ctx wasmlib.ScViewCallContext) *GetTokenPerPackageCall {
	f := &GetTokenPerPackageCall{Func: wasmlib.NewScView(ctx, HScName, HViewGetTokenPerPackage)}
	f.Params.proxy = wasmlib.NewCallParamsProxy(f.Func)
	wasmlib.NewCallResultsProxy(f.Func, &f.Results.proxy)
	return f
}