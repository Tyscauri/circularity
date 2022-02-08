// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]

use wasmlib::*;

pub const SC_NAME        : &str = "test3";
pub const SC_DESCRIPTION : &str = "test3 description";
pub const HSC_NAME       : ScHname = ScHname(0x1ffeb901);

pub const PARAM_CHARGE_WEIGHT : &str = "chargeWeight";
pub const PARAM_COMP          : &str = "comp";
pub const PARAM_FRAC_ID       : &str = "fracID";
pub const PARAM_ID            : &str = "id";
pub const PARAM_MAT           : &str = "mat";
pub const PARAM_NAME          : &str = "name";
pub const PARAM_OWNER         : &str = "owner";
pub const PARAM_PP_ID         : &str = "ppID";
pub const PARAM_PROD_PASS     : &str = "prodPass";
pub const PARAM_PROP          : &str = "prop";

pub const RESULT_COMPOSITIONS      : &str = "compositions";
pub const RESULT_FRAC_ID           : &str = "fracID";
pub const RESULT_FRACTION          : &str = "fraction";
pub const RESULT_ID                : &str = "id";
pub const RESULT_OWNER             : &str = "owner";
pub const RESULT_PP                : &str = "pp";
pub const RESULT_PPNAME            : &str = "ppname";
pub const RESULT_PPRESULT          : &str = "ppresult";
pub const RESULT_RECYCLATE_ID      : &str = "recyclateID";
pub const RESULT_TOKEN_PER_PACKAGE : &str = "tokenPerPackage";
pub const RESULT_TOKEN_REQUIRED    : &str = "tokenRequired";

pub const STATE_COMPOSITIONS      : &str = "compositions";
pub const STATE_FRAC_COMPOSITIONS : &str = "fracCompositions";
pub const STATE_FRACTIONS         : &str = "fractions";
pub const STATE_OWNER             : &str = "owner";
pub const STATE_PRICE_PER_MG      : &str = "pricePerMg";
pub const STATE_PRODUCTPASSES     : &str = "productpasses";
pub const STATE_RECY_COMPOSITIONS : &str = "recyCompositions";
pub const STATE_RECYCLATES        : &str = "recyclates";
pub const STATE_SHARE_RECYCLER    : &str = "shareRecycler";

pub const FUNC_ADD_MATERIAL                 : &str = "addMaterial";
pub const FUNC_ADD_PP_TO_FRACTION           : &str = "addPPToFraction";
pub const FUNC_CREATE_FRACTION              : &str = "createFraction";
pub const FUNC_CREATE_PP                    : &str = "createPP";
pub const FUNC_CREATE_RECYCLATE             : &str = "createRecyclate";
pub const FUNC_INIT                         : &str = "init";
pub const FUNC_SET_MATERIALS                : &str = "setMaterials";
pub const FUNC_SET_OWNER                    : &str = "setOwner";
pub const VIEW_GET_AMOUNT_OF_REQUIRED_FUNDS : &str = "getAmountOfRequiredFunds";
pub const VIEW_GET_MATERIALS                : &str = "getMaterials";
pub const VIEW_GET_OWNER                    : &str = "getOwner";
pub const VIEW_GET_PP                       : &str = "getPP";
pub const VIEW_GET_TOKEN_PER_PACKAGE        : &str = "getTokenPerPackage";

pub const HFUNC_ADD_MATERIAL                 : ScHname = ScHname(0xdfeea1ff);
pub const HFUNC_ADD_PP_TO_FRACTION           : ScHname = ScHname(0x50ac50a7);
pub const HFUNC_CREATE_FRACTION              : ScHname = ScHname(0x59842fc3);
pub const HFUNC_CREATE_PP                    : ScHname = ScHname(0x673fc3d7);
pub const HFUNC_CREATE_RECYCLATE             : ScHname = ScHname(0x5066d840);
pub const HFUNC_INIT                         : ScHname = ScHname(0x1f44d644);
pub const HFUNC_SET_MATERIALS                : ScHname = ScHname(0x7f0ebcae);
pub const HFUNC_SET_OWNER                    : ScHname = ScHname(0x2a15fe7b);
pub const HVIEW_GET_AMOUNT_OF_REQUIRED_FUNDS : ScHname = ScHname(0xbcbaf2d6);
pub const HVIEW_GET_MATERIALS                : ScHname = ScHname(0x1dca8ddb);
pub const HVIEW_GET_OWNER                    : ScHname = ScHname(0x137107a6);
pub const HVIEW_GET_PP                       : ScHname = ScHname(0xbf711e08);
pub const HVIEW_GET_TOKEN_PER_PACKAGE        : ScHname = ScHname(0x522a14c0);