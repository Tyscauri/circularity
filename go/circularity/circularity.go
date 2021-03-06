// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package circularity

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib"

func funcAddMaterial(ctx wasmlib.ScFuncContext, f *AddMaterialContext) {
}

func funcAddPPToFraction(ctx wasmlib.ScFuncContext, f *AddPPToFractionContext) {
}

func funcCreateFraction(ctx wasmlib.ScFuncContext, f *CreateFractionContext) {
}

func funcCreatePP(ctx wasmlib.ScFuncContext, f *CreatePPContext) {
}

func funcCreateRecyclate(ctx wasmlib.ScFuncContext, f *CreateRecyclateContext) {
}

func funcInit(ctx wasmlib.ScFuncContext, f *InitContext) {
    if f.Params.Owner().Exists() {
        f.State.Owner().SetValue(f.Params.Owner().Value())
        return
    }
    f.State.Owner().SetValue(ctx.ContractCreator())
}

func funcPayoutFrac(ctx wasmlib.ScFuncContext, f *PayoutFracContext) {
}

func funcSetMaterials(ctx wasmlib.ScFuncContext, f *SetMaterialsContext) {
}

func funcSetOwner(ctx wasmlib.ScFuncContext, f *SetOwnerContext) {
	f.State.Owner().SetValue(f.Params.Owner().Value())
}

func viewGetAmountOfRequiredFunds(ctx wasmlib.ScViewContext, f *GetAmountOfRequiredFundsContext) {
}

func viewGetMaterials(ctx wasmlib.ScViewContext, f *GetMaterialsContext) {
}

func viewGetOwner(ctx wasmlib.ScViewContext, f *GetOwnerContext) {
	f.Results.Owner().SetValue(f.State.Owner().Value())
}

func viewGetPP(ctx wasmlib.ScViewContext, f *GetPPContext) {
}

func viewGetTokenPerPackage(ctx wasmlib.ScViewContext, f *GetTokenPerPackageContext) {
}

func funcTickerPPToSameFraction(ctx wasmlib.ScFuncContext, f *TickerPPToSameFractionContext) {
}

func funcPayoutProducer(ctx wasmlib.ScFuncContext, f *PayoutProducerContext) {
}

func funcDeletePP(ctx wasmlib.ScFuncContext, f *DeletePPContext) {
}
