// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"math"

	"github.com/cpmech/gosl/io"
	"github.com/cpmech/gosl/mpi"
	"github.com/cpmech/gosl/ode"
)

func main() {

	// constants
	N := 100000
	tolExponent := 2
	tol := math.Pow(0.1, float64(tolExponent))

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

	// configurations
	conf := ode.NewConfig("radau5", "mumps", comm)
	conf.SetStepOut(true, nil)
	conf.SetTol(tol)
	conf.IniH = 0.05
	conf.GoChan = false

	// solver
	ndim := 2 * N
	sol := ode.NewSolver(ndim, conf, fcn, jac, nil)

	// solve
	sol.Solve(y, 0.0, 0.1)

	// output root
	if mpi.WorldRank() == 0 {
		fn := io.Sf("n%d_gosl_tolM%d_np%d_%s.txt", N, tolExponent, comm.Size(), sol.Stat.LsKind)
		if N < 50 {
			io.WriteTableVD("results", fn, []string{"yend"}, y)
		}
		io.Pf("N                         = %v\n", N)
		io.Pf("tolerance                 = %v\n", tol)
		sol.Stat.Print(true)
	}
}
