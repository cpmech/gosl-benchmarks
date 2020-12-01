// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"testing"

	"_/lib/bmark"
	"_/lib/reference"

	"github.com/cpmech/gosl/chk"
	"github.com/cpmech/gosl/io"
	"github.com/cpmech/gosl/la"
)

var globalSW *bmark.Stopwatch

func init() {
	globalSW = bmark.StartNewStopwatch()
}

func results() {
	globalSW.Stop("... ")
	bmark.MemoryUsage("... ")
	globalSW.Reset()
}

func main() {
	// fnkey := "inline_1"
	// fnkey := "audikw_1" // fail
	// fnkey := "Flan_1565" // fail
	// fnkey := "atmosmodl" // fail
	// fnkey := "tmt_unsym"
	// fnkey := "Hamrle3"
	// fnkey := "pre2"
	fnkey := "bfwb62"

	// allocate solver
	kind := "umfpack"
	solver := la.NewSparseSolver(kind)
	defer func() {
		solver.Free()
	}()

	io.Pf("reading matrix (%s)\n", fnkey)
	T := new(la.Triplet)
	symmetric := T.ReadSmat(io.Sf("../data/%s.mtx", fnkey), true, nil)
	m, n := T.Size()
	if m != n {
		chk.Panic("matrix must be square. m=%d, n=%d\n", m, n)
	}
	io.Pf("... symmetric = %v\n", symmetric)
	io.Pf("... number of rows (equal to columns) = %d\n", m)
	io.Pf("... number of non-zeros (pattern entries) = %d\n", T.Len())
	results()

	x := la.NewVector(m)
	b := la.NewVector(m)
	b.Fill(1)

	sw := bmark.StartNewStopwatch()
	io.Pf("initializing (%s)\n", kind)
	args := la.NewSparseConfig(nil)
	if symmetric {
		args.SetUmfpackSymmetry()
	}
	solver.Init(T, args)
	results()

	io.Pf("factorizing (%s)\n", kind)
	solver.Fact()
	results()

	io.Pf("solving (%s)\n", kind)
	solver.Solve(x, b, false)
	results()
	sw.Stop("... total (without reading) ")

	if fnkey == "bfwb62" {
		chk.Verbose = true
		tst := new(testing.T)
		chk.Array(tst, "x", 1e-15, x, reference.XCorrectBfwb62)
		xx := la.NewVector(m)
		la.DenSolve(xx, T.ToDense(), b, false)
		chk.Array(tst, "xx", 1e-10, xx, reference.XCorrectBfwb62)
	}
}
