// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"gosl/chk"
	"gosl/io"
	"gosl/la"
	"gosl/mpi"
	"gosl/num"
	"testing"
)

// TestJacobian tests the Jacobian function
func TestJacobian(tst *testing.T) {

	// flags
	io.Verbose = true
	chk.Verbose = true
	chk.PrintTitle("TestJacobian")

	// constants
	N := 6

	// start mpi
	mpi.Start()
	defer mpi.Stop()

	// check number of processors
	if mpi.WorldSize() > N {
		if mpi.WorldRank() == 0 {
			io.Pf("ERROR: the maximum number of processors is %d\n", N)
		}
		return
	}

	// communicator
	comm := mpi.NewCommunicator(nil)

	// problem
	fcn, jac, y := equations(N, comm)

	// test function
	tstFcn := func(fy, y la.Vector) {
		fcn(fy, 0, 0, y)
	}

	// test Jacobian function
	tstJac := func(dfydy *la.Triplet, y la.Vector) {
		jac(dfydy, 0, 0, y)
	}

	// check
	num.CompareJacMpi(tst, comm, tstFcn, tstJac, y, 1e-6, true)
}
