// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]

use wasmlib::*;

pub struct test3Events {
}

impl test3Events {

	pub fn payout(&self) {
		EventEncoder::new("test3.payout").emit();
	}
}
