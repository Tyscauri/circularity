// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"testing"

	"home/tim/Schreibtisch/circ_proto/test3/go/test3"
	"github.com/iotaledger/wasp/packages/vm/wasmsolo"
	"github.com/stretchr/testify/require"
)

func TestDeploy(t *testing.T) {
	ctx := wasmsolo.NewSoloContext(t, test3.ScName, test3.OnLoad)
	require.NoError(t, ctx.ContractExists(test3.ScName))
}

func setupTest(t *testing.T) *wasmsolo.SoloContext {
	return wasmsolo.NewSoloContext(t, test3.ScName, test3.OnLoad)
}

func TestCreateAndGetPP(t * testing.T) {
        ctx:= setupTest(t)
        createPP:= test3.ScFuncs.CreatePP(ctx.Sign(ctx.NewSoloAgent())) 	//for funcs
        createPP.Params.Name().SetValue("Chips")
        createPP.Params.Id().SetValue("123")
        createPP.Func.TransferIotas(1).Post()
        
        addMaterial:=test3.ScFuncs.AddMaterial(ctx.Sign(ctx.NewSoloAgent()))
        addMaterial.Params.Id().SetValue("123")
        addMaterial.Params.Mat().SetValue("PP")
        addMaterial.Params.Prop().SetValue(60)
        addMaterial.Func.TransferIotas(1).Post()
        
        addMaterial2:=test3.ScFuncs.AddMaterial(ctx.Sign(ctx.NewSoloAgent()))
        addMaterial2.Params.Id().SetValue("123")
        addMaterial2.Params.Mat().SetValue("PE")
        addMaterial2.Params.Prop().SetValue(40)
        addMaterial2.Func.TransferIotas(1).Post()
                
        createPP2:= test3.ScFuncs.CreatePP(ctx.Sign(ctx.NewSoloAgent()))
        createPP2.Params.Name().SetValue("wurscht")
        createPP2.Params.Id().SetValue("1234")
        createPP2.Func.TransferIotas(1).Post()
        
        //test getPP
        getPP := test3.ScFuncs.GetPP(ctx)				//for views
        getPP.Params.Id().SetValue("123")                            // for views
        getPP.Func.Call()
        require.NoError(t, ctx.Err)
        require.EqualValues(t, "Chips", getPP.Results.Ppname().Value())
        
        getPP2 := test3.ScFuncs.GetPP(ctx)			
        getPP2.Params.Id().SetValue("1234")
        getPP2.Func.Call()
        require.NoError(t, ctx.Err)
        require.NotEqualValues(t, "Chips", getPP.Results.Ppname().Value())
        require.NotEqualValues(t, "Chips", getPP.Results.Ppresult().Value().Name)
        
        getMat1 := test3.ScFuncs.GetMaterials(ctx)
        getMat1.Params.Id().SetValue("123")
        getMat1.Func.Call()
        require.NoError(t, ctx.Err)
        require.EqualValues(t, "PP", getMat1.Results.Compositions().GetComposition(0).Value().Material)
        require.EqualValues(t, 60, getMat1.Results.Compositions().GetComposition(0).Value().Proportion)
        require.EqualValues(t, "PE", getMat1.Results.Compositions().GetComposition(1).Value().Material)
        require.EqualValues(t, 40, getMat1.Results.Compositions().GetComposition(1).Value().Proportion)

} 	


func PowerTest(t* testing.T) {
}
