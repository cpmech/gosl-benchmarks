package main

import (
	"_/lib/bmark"
	"gosl/chk"
	"gosl/io"
	"gosl/la"
	"gosl/mpi"
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
	// fnkey := "audikw_1"
	// fnkey := "Flan_1565"
	// fnkey := "atmosmodl"
	fnkey := "tmt_unsym"

	// allocate communicator and solver
	mpi.Start()
	comm := mpi.NewCommunicator(nil)
	kind := "mumps"
	solver := la.NewSparseSolver(kind)
	defer func() {
		solver.Free()
		mpi.Stop()
	}()

	io.Pf("reading matrix (%s)\n", fnkey)
	T := new(la.Triplet)
	symmetric := T.ReadSmat(io.Sf("../data/%s.mtx", fnkey))
	if symmetric {
		io.Pf("... is symmetric\n")
	}
	results()

	m, n := T.Size()
	if m != n {
		chk.Panic("matrix must be square. m=%d, n=%d\n", m, n)
	}
	x := la.NewVector(m)
	b := la.NewVector(m)
	b.Fill(1)

	sw := bmark.StartNewStopwatch()
	io.Pf("initializing (%s)\n", kind)
	args := la.NewSparseConfig(comm)
	args.Symmetric = symmetric
	args.MumpsMaxMemoryPerProcessor = 3000
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
