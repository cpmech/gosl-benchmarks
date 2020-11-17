package main

import (
	"_/lib/bmark"
	"gosl/chk"
	"gosl/io"
	"gosl/la"
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
	fnkey := "pre2"

	// allocate solver
	kind := "umfpack"
	solver := la.NewSparseSolver(kind)
	defer func() {
		solver.Free()
	}()

	io.Pf("reading matrix (%s)\n", fnkey)
	T := new(la.Triplet)
	symmetric := T.ReadSmat(io.Sf("../data/%s.mtx", fnkey))
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
	args.Symmetric = symmetric
	solver.Init(T, args)
	results()

	io.Pf("factorizing (%s)\n", kind)
	solver.Fact()
	results()

	io.Pf("solving (%s)\n", kind)
	solver.Solve(x, b, false)
	results()
	sw.Stop("... total (without reading) ")
}
