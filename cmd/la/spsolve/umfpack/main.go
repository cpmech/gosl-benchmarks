// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"testing"

	"_/lib/bmark"

	"github.com/cpmech/gosl/chk"
	"github.com/cpmech/gosl/io"
	"github.com/cpmech/gosl/la"
)

var xCorrectBfwb62 = []float64{
	-1.02570048377040759e+05,
	-1.08800418159713998e+05,
	-7.87848688672370918e+04,
	-6.12550631774225840e+04,
	-1.16611533352550643e+05,
	-8.91949258261042705e+04,
	-5.57584825429375196e+04,
	-3.37535346291137103e+04,
	-6.74159236038033268e+04,
	-5.61065283435406673e+04,
	-3.69561341372605821e+04,
	-2.67385128650871302e+04,
	-4.67349124343154253e+04,
	-4.18861901056076676e+04,
	-4.34393771636046149e+04,
	-1.11210692731083000e+04,
	-1.16010526640020762e+04,
	-4.31993854681577286e+04,
	-5.82924327463857844e+03,
	-2.42374319876188747e+04,
	-2.39432136682168457e+04,
	+5.27355041927211232e+02,
	-1.24769422505944240e+04,
	-1.47005934749971748e+04,
	-4.95701604733381391e+04,
	-1.38451884223610182e+03,
	-1.57972501695015781e+04,
	-5.19172705598900066e+04,
	-4.99494464999615593e+04,
	-1.19678659380488571e+04,
	-1.56190973892000347e+04,
	-6.18809904102459404e+03,
	-1.05693761694190998e+04,
	-2.93013328593191145e+04,
	-9.15514607143451940e+03,
	-1.27058094439569140e+04,
	-1.93936053067287430e+04,
	-6.84836276779992295e+03,
	-1.07869319688850719e+04,
	-4.61926223513438963e+04,
	-1.99579363156562504e+04,
	-7.83564896339727693e+03,
	-6.37173129434054590e+03,
	-1.88075622025074267e+03,
	-8.71648101674354621e+03,
	-1.21683775603205122e+04,
	-1.91184585274694587e+03,
	-5.64233479410600103e+03,
	-6.47747230904305070e+03,
	-4.47783973932844674e+03,
	-9.82971659947420812e+03,
	-1.95594295004403466e+04,
	-2.09457080830507803e+04,
	-5.46686114796283709e+03,
	-5.28888244321673483e+03,
	-2.07962090362636227e+04,
	-9.33272319073228937e+03,
	+1.96672299472196187e+02,
	-4.40813445835840230e+03,
	-4.87188111893421956e+03,
	-1.75640594405328884e+04,
	-1.77959327708208002e+04,
}

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
		chk.Array(tst, "x", 1e-15, x, xCorrectBfwb62)
		xx := la.NewVector(m)
		la.DenSolve(xx, T.ToDense(), b, false)
		chk.Array(tst, "xx", 1e-10, xx, xCorrectBfwb62)
	}
}
