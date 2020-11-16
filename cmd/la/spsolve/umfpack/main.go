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
	fnkey := "inline_1"

	// allocate solver
	solver := la.NewSparseSolver("umfpack")
	defer func() {
		solver.Free()
	}()

	io.Pf("reading matrix (%s)\n", fnkey)
	T := new(la.Triplet)
	symmetric := T.ReadSmat(io.Sf("../data/%s.mtx", fnkey))
	if symmetric {
		io.Pf("    is symmetric\n")
	}
	results()

	m, n := T.Size()
	if m != n {
		chk.Panic("matrix must be square. m=%d, n=%d\n", m, n)
	}
	x := la.NewVector(m)
	b := la.NewVector(m)

	io.Pf("initializing\n")
	args := la.NewSparseConfig(nil)
	args.Symmetric = symmetric
	solver.Init(T, args)
	results()

	io.Pf("factorizing\n")
	solver.Fact()
	results()

	io.Pf("solving\n")
	solver.Solve(x, b, false)
	results()
}
