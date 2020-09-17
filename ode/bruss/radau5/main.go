// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"gosl/io"
	"gosl/mpi"
	"gosl/ode"
)

func main() {

	// constants
	N := 6
	ndim := 2 * N
	xend := 10.0

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
	fcn, jac, y := problem(N, comm)

	// configurations
	conf := ode.NewConfig("radau5", "", comm)
	conf.SetStepOut(true, nil)
	conf.SetTol(1e-2)

	// solver
	sol := ode.NewSolver(ndim, conf, fcn, jac, nil)

	// solve
	sol.Solve(y, 0, xend)

	// output root
	if mpi.WorldRank() == 0 {
		io.WriteTableVD("results", "yend.txt", []string{"y"}, y)
		sol.Stat.Print(true)
	}
}
