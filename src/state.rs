// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]
#![allow(unused_imports)]

use wasmlib::*;

use crate::*;

#[derive(Clone)]
pub struct MapHashToImmutableCompositions {
	pub(crate) proxy: Proxy,
}

impl MapHashToImmutableCompositions {
    pub fn get_compositions(&self, key: &ScHash) -> ImmutableCompositions {
        ImmutableCompositions { proxy: self.proxy.key(&hash_to_bytes(key)) }
    }
}

#[derive(Clone)]
pub struct MapHashToImmutableFracCompositions {
	pub(crate) proxy: Proxy,
}

impl MapHashToImmutableFracCompositions {
    pub fn get_frac_compositions(&self, key: &ScHash) -> ImmutableFracCompositions {
        ImmutableFracCompositions { proxy: self.proxy.key(&hash_to_bytes(key)) }
    }
}

#[derive(Clone)]
pub struct MapHashToImmutableFraction {
	pub(crate) proxy: Proxy,
}

impl MapHashToImmutableFraction {
    pub fn get_fraction(&self, key: &ScHash) -> ImmutableFraction {
        ImmutableFraction { proxy: self.proxy.key(&hash_to_bytes(key)) }
    }
}

#[derive(Clone)]
pub struct MapHashToImmutableProductPass {
	pub(crate) proxy: Proxy,
}

impl MapHashToImmutableProductPass {
    pub fn get_product_pass(&self, key: &ScHash) -> ImmutableProductPass {
        ImmutableProductPass { proxy: self.proxy.key(&hash_to_bytes(key)) }
    }
}

#[derive(Clone)]
pub struct MapHashToImmutableRecyCompositions {
	pub(crate) proxy: Proxy,
}

impl MapHashToImmutableRecyCompositions {
    pub fn get_recy_compositions(&self, key: &ScHash) -> ImmutableRecyCompositions {
        ImmutableRecyCompositions { proxy: self.proxy.key(&hash_to_bytes(key)) }
    }
}

#[derive(Clone)]
pub struct MapHashToImmutableRecyclate {
	pub(crate) proxy: Proxy,
}

impl MapHashToImmutableRecyclate {
    pub fn get_recyclate(&self, key: &ScHash) -> ImmutableRecyclate {
        ImmutableRecyclate { proxy: self.proxy.key(&hash_to_bytes(key)) }
    }
}

#[derive(Clone)]
pub struct ImmutablecircularityState {
	pub(crate) proxy: Proxy,
}

impl ImmutablecircularityState {
    pub fn compositions(&self) -> MapHashToImmutableCompositions {
		MapHashToImmutableCompositions { proxy: self.proxy.root(STATE_COMPOSITIONS) }
	}

    pub fn donation_address(&self) -> ScImmutableAgentID {
		ScImmutableAgentID::new(self.proxy.root(STATE_DONATION_ADDRESS))
	}

    pub fn frac_compositions(&self) -> MapHashToImmutableFracCompositions {
		MapHashToImmutableFracCompositions { proxy: self.proxy.root(STATE_FRAC_COMPOSITIONS) }
	}

    pub fn fractions(&self) -> MapHashToImmutableFraction {
		MapHashToImmutableFraction { proxy: self.proxy.root(STATE_FRACTIONS) }
	}

    pub fn owner(&self) -> ScImmutableAgentID {
		ScImmutableAgentID::new(self.proxy.root(STATE_OWNER))
	}

    pub fn price_per_mg(&self) -> ScImmutableUint64 {
		ScImmutableUint64::new(self.proxy.root(STATE_PRICE_PER_MG))
	}

    pub fn productpasses(&self) -> MapHashToImmutableProductPass {
		MapHashToImmutableProductPass { proxy: self.proxy.root(STATE_PRODUCTPASSES) }
	}

    pub fn recy_compositions(&self) -> MapHashToImmutableRecyCompositions {
		MapHashToImmutableRecyCompositions { proxy: self.proxy.root(STATE_RECY_COMPOSITIONS) }
	}

    pub fn recyclates(&self) -> MapHashToImmutableRecyclate {
		MapHashToImmutableRecyclate { proxy: self.proxy.root(STATE_RECYCLATES) }
	}

    pub fn share_recycler(&self) -> ScImmutableUint8 {
		ScImmutableUint8::new(self.proxy.root(STATE_SHARE_RECYCLER))
	}

    pub fn token_to_donate(&self) -> ScImmutableUint64 {
		ScImmutableUint64::new(self.proxy.root(STATE_TOKEN_TO_DONATE))
	}
}

#[derive(Clone)]
pub struct MapHashToMutableCompositions {
	pub(crate) proxy: Proxy,
}

impl MapHashToMutableCompositions {
    pub fn clear(&self) {
        self.proxy.clear_map();
    }

    pub fn get_compositions(&self, key: &ScHash) -> MutableCompositions {
        MutableCompositions { proxy: self.proxy.key(&hash_to_bytes(key)) }
    }
}

#[derive(Clone)]
pub struct MapHashToMutableFracCompositions {
	pub(crate) proxy: Proxy,
}

impl MapHashToMutableFracCompositions {
    pub fn clear(&self) {
        self.proxy.clear_map();
    }

    pub fn get_frac_compositions(&self, key: &ScHash) -> MutableFracCompositions {
        MutableFracCompositions { proxy: self.proxy.key(&hash_to_bytes(key)) }
    }
}

#[derive(Clone)]
pub struct MapHashToMutableFraction {
	pub(crate) proxy: Proxy,
}

impl MapHashToMutableFraction {
    pub fn clear(&self) {
        self.proxy.clear_map();
    }

    pub fn get_fraction(&self, key: &ScHash) -> MutableFraction {
        MutableFraction { proxy: self.proxy.key(&hash_to_bytes(key)) }
    }
}

#[derive(Clone)]
pub struct MapHashToMutableProductPass {
	pub(crate) proxy: Proxy,
}

impl MapHashToMutableProductPass {
    pub fn clear(&self) {
        self.proxy.clear_map();
    }

    pub fn get_product_pass(&self, key: &ScHash) -> MutableProductPass {
        MutableProductPass { proxy: self.proxy.key(&hash_to_bytes(key)) }
    }
}

#[derive(Clone)]
pub struct MapHashToMutableRecyCompositions {
	pub(crate) proxy: Proxy,
}

impl MapHashToMutableRecyCompositions {
    pub fn clear(&self) {
        self.proxy.clear_map();
    }

    pub fn get_recy_compositions(&self, key: &ScHash) -> MutableRecyCompositions {
        MutableRecyCompositions { proxy: self.proxy.key(&hash_to_bytes(key)) }
    }
}

#[derive(Clone)]
pub struct MapHashToMutableRecyclate {
	pub(crate) proxy: Proxy,
}

impl MapHashToMutableRecyclate {
    pub fn clear(&self) {
        self.proxy.clear_map();
    }

    pub fn get_recyclate(&self, key: &ScHash) -> MutableRecyclate {
        MutableRecyclate { proxy: self.proxy.key(&hash_to_bytes(key)) }
    }
}

#[derive(Clone)]
pub struct MutablecircularityState {
	pub(crate) proxy: Proxy,
}

impl MutablecircularityState {
    pub fn as_immutable(&self) -> ImmutablecircularityState {
		ImmutablecircularityState { proxy: self.proxy.root("") }
	}

    pub fn compositions(&self) -> MapHashToMutableCompositions {
		MapHashToMutableCompositions { proxy: self.proxy.root(STATE_COMPOSITIONS) }
	}

    pub fn donation_address(&self) -> ScMutableAgentID {
		ScMutableAgentID::new(self.proxy.root(STATE_DONATION_ADDRESS))
	}

    pub fn frac_compositions(&self) -> MapHashToMutableFracCompositions {
		MapHashToMutableFracCompositions { proxy: self.proxy.root(STATE_FRAC_COMPOSITIONS) }
	}

    pub fn fractions(&self) -> MapHashToMutableFraction {
		MapHashToMutableFraction { proxy: self.proxy.root(STATE_FRACTIONS) }
	}

    pub fn owner(&self) -> ScMutableAgentID {
		ScMutableAgentID::new(self.proxy.root(STATE_OWNER))
	}

    pub fn price_per_mg(&self) -> ScMutableUint64 {
		ScMutableUint64::new(self.proxy.root(STATE_PRICE_PER_MG))
	}

    pub fn productpasses(&self) -> MapHashToMutableProductPass {
		MapHashToMutableProductPass { proxy: self.proxy.root(STATE_PRODUCTPASSES) }
	}

    pub fn recy_compositions(&self) -> MapHashToMutableRecyCompositions {
		MapHashToMutableRecyCompositions { proxy: self.proxy.root(STATE_RECY_COMPOSITIONS) }
	}

    pub fn recyclates(&self) -> MapHashToMutableRecyclate {
		MapHashToMutableRecyclate { proxy: self.proxy.root(STATE_RECYCLATES) }
	}

    pub fn share_recycler(&self) -> ScMutableUint8 {
		ScMutableUint8::new(self.proxy.root(STATE_SHARE_RECYCLER))
	}

    pub fn token_to_donate(&self) -> ScMutableUint64 {
		ScMutableUint64::new(self.proxy.root(STATE_TOKEN_TO_DONATE))
	}
}
