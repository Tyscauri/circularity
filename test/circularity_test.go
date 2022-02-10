// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"testing"

	"home/tim/Schreibtisch/circ_proto/test3/go/circularity"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmsolo"
	"github.com/stretchr/testify/require"
)

func TestDeploy(t *testing.T) {
	ctx := wasmsolo.NewSoloContext(t, circularity.ScName, circularity.OnLoad)
	require.NoError(t, ctx.ContractExists(circularity.ScName))
}
