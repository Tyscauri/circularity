// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test3

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

func funcCreatePP(ctx wasmlib.ScFuncContext, f *CreatePPContext) {
}

func funcInit(ctx wasmlib.ScFuncContext, f *InitContext) {
    if f.Params.Owner().Exists() {
        f.State.Owner().SetValue(f.Params.Owner().Value())
        return
    }
    f.State.Owner().SetValue(ctx.ContractCreator())
}

func funcSetOwner(ctx wasmlib.ScFuncContext, f *SetOwnerContext) {
	f.State.Owner().SetValue(f.Params.Owner().Value())
}

func viewGetAmountOfRequiredFunds(ctx wasmlib.ScViewContext, f *GetAmountOfRequiredFundsContext) {
}

func viewGetOwner(ctx wasmlib.ScViewContext, f *GetOwnerContext) {
	f.Results.Owner().SetValue(f.State.Owner().Value())
}

func viewGetPP(ctx wasmlib.ScViewContext, f *GetPPContext) {
}

func viewGetTokenPerPackage(ctx wasmlib.ScViewContext, f *GetTokenPerPackageContext) {
}

func funcAddMaterial(ctx wasmlib.ScFuncContext, f *AddMaterialContext) {
}

func viewGetMaterials(ctx wasmlib.ScViewContext, f *GetMaterialsContext) {
}

func funcSetMaterials(ctx wasmlib.ScFuncContext, f *SetMaterialsContext) {
}

func funcAddPPToFraction(ctx wasmlib.ScFuncContext, f *AddPPToFractionContext) {
}

func funcCreateFraction(ctx wasmlib.ScFuncContext, f *CreateFractionContext) {
}
