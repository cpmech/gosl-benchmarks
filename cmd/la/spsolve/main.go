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
	if comm == nil {
		io.Pf(msg, prm...)
	} else if comm.Rank() == 0 {
		io.Pf(msg, prm...)
	}
}

func main() {
	// parse flags
	var kind string // "mumps" or "umfpack"
	var fnkey string
	var ordering string // mumps' ordering
	flag.StringVar(&kind, "kind", "mumps", "\"mumps\" or \"umfpack\"")
	flag.StringVar(&fnkey, "fnkey", "bfwb62", "fnkey = matrix name")
	flag.StringVar(&ordering, "ordering", "", "ordering for MUMPS")
	flag.Parse()

	// check kind
	if kind != "mumps" && kind != "umfpack" {
		chk.Panic("kind must be \"mumps\" or \"umfpack\"")
	}

	// allocate communicator
	var comm *mpi.Communicator
	if kind == "mumps" {
		mpi.Start()
		comm = mpi.NewCommunicator(nil)
		defer mpi.Stop()
	}

	// allocate SolverResults
	res := new(bmark.SolverResults)
	res.Kind = kind
	res.Ordering = ordering

	// whether to load the full matrix or not, if symmetric
	mirrorIfSym := false
	if kind == "umfpack" {
		mirrorIfSym = true
	}

	// read matrix
	pf(comm, "reading matrix (%s)\n", fnkey)
	T := new(la.Triplet)
	res.Symmetric = T.ReadSmat(io.Sf("./data/%s.mtx", fnkey), mirrorIfSym, comm)
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

	// allocate solver
	solver := la.NewSparseSolver(kind)
	defer solver.Free()

	// initialize solver
	pf(comm, "initializing (%s)\n", kind)
	args := la.NewSparseConfig(comm)
	if res.Symmetric {
		args.SetMumpsSymmetry(true, false)
	}
	if kind == "mumps" {
		args.SetMumpsOrdering(ordering)
		args.MumpsMaxMemoryPerProcessor = 30000
	}
	args.Verbose = true
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

	// calc solver time
	res.CalcSolverTime()

	// save results
	outFnkey := io.Sf("%s_%s", kind, fnkey)
	if ordering != "" {
		outFnkey += "_" + ordering
	}
	if comm == nil {
		res.Save("./results", outFnkey)
	} else if comm.Rank() == 0 {
		res.MpiSize = comm.Size()
		res.Save("./results", io.Sf("%s_np%d", outFnkey, comm.Size()))
	}

	// check
	if fnkey == "bfwb62" {
		chk.Verbose = true
		tst := new(testing.T)
		chk.Array(tst, "x", 1e-10, x, reference.XCorrectBfwb62)
		if comm != nil {
			if comm.Size() == 1 {
				xx := la.NewVector(res.NumberOfRows)
				Td := new(la.Triplet)
				Td.ReadSmat(io.Sf("./data/%s.mtx", fnkey), true, nil)
				la.DenSolve(xx, Td.ToDense(), b, false)
				chk.Array(tst, "xx", 1e-10, xx, reference.XCorrectBfwb62)
			}
		}
	}
}
