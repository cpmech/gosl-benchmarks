// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"testing"

	"_/lib/bmark"
	"_/lib/reference"

	"github.com/cpmech/gosl/chk"
	"github.com/cpmech/gosl/io"
	"github.com/cpmech/gosl/la"
	"github.com/cpmech/gosl/mpi"
)

func pf(comm *mpi.Communicator, msg string, prm ...interface{}) {
	if comm.Rank() == 0 {
		io.Pf(msg, prm...)
	}
}

func main() {
	// parse flags
	var fnkey string
	flag.StringVar(&fnkey, "fnkey", "bfwb62", "fnkey = matrix name")
	flag.Parse()

	// allocate communicator and solver
	mpi.Start()
	comm := mpi.NewCommunicator(nil)
	kind := "mumps"
	solver := la.NewSparseSolver(kind)
	defer func() {
		solver.Free()
		mpi.Stop()
	}()

	// allocate SolverResults
	res := new(bmark.SolverResults)

	// read matrix
	pf(comm, "reading matrix (%s)\n", fnkey)
	T := new(la.Triplet)
	res.Symmetric = T.ReadSmat(io.Sf("../data/%s.mtx", fnkey), false, comm)
	res.NumberOfRows, res.NumberOfCols = T.Size()
	res.NumberOfNonZeros = T.Len()
	if res.NumberOfRows != res.NumberOfCols {
		chk.Panic("matrix must be square. m=%d, n=%d\n", res.NumberOfRows, res.NumberOfCols)
	}
	pf(comm, "... symmetric = %v\n", res.Symmetric)
	pf(comm, "... number of rows (equal to columns) = %d\n", res.NumberOfRows)
	pf(comm, "... number of non-zeros (pattern entries) = %d\n", res.NumberOfNonZeros)
	res.StepReadMatrix = bmark.MeasureTimeAndMemory(true, comm)

	// allocate vectors and set right-hand-side
	x := la.NewVector(res.NumberOfRows)
	b := la.NewVector(res.NumberOfRows)
	b.Fill(1)

	// initialize solver
	pf(comm, "initializing (%s)\n", kind)
	args := la.NewSparseConfig(comm)
	if res.Symmetric {
		args.SetMumpsSymmetry(true, false)
	}
	args.MumpsMaxMemoryPerProcessor = 30000
	solver.Init(T, args)
	res.StepInitialize = bmark.MeasureTimeAndMemory(true, comm)

	// perform factorization
	pf(comm, "factorizing (%s)\n", kind)
	solver.Fact()
	res.StepFactorize = bmark.MeasureTimeAndMemory(true, comm)

	// solve system
	pf(comm, "solving (%s)\n", kind)
	solver.Solve(x, b, false)
	res.StepSolve = bmark.MeasureTimeAndMemory(true, comm)

	// calc solver time and save results
	res.CalcSolverTime()
	res.Save("../results", io.Sf("%s_%s", kind, fnkey))

	// check
	if fnkey == "bfwb62" {
		chk.Verbose = true
		tst := new(testing.T)
		chk.Array(tst, "x", 1e-10, x, reference.XCorrectBfwb62)
		if comm.Size() == 1 {
			xx := la.NewVector(res.NumberOfRows)
			Td := new(la.Triplet)
			Td.ReadSmat(io.Sf("../data/%s.mtx", fnkey), true, nil)
			la.DenSolve(xx, Td.ToDense(), b, false)
			chk.Array(tst, "xx", 1e-10, xx, reference.XCorrectBfwb62)
		}
	}
}
