// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"gosl/la"
	"gosl/mpi"
	"math"
)

func equations(N int, comm *mpi.Communicator) (
	fcn func(f la.Vector, dx, x float64, y la.Vector),
	jac func(dfdy *la.Triplet, dx, x float64, y la.Vector),
	yini la.Vector,
) {
	// constants
	ndim := 2 * N
	gamma := 0.02 * math.Pow(float64(N+1), 2)

	// partition variables
	id, sz := comm.Rank(), comm.Size()
	start, endp1 := (id*N)/sz, ((id+1)*N)/sz

	// workspace
	w := la.NewVector(ndim)

	// dy/dx function
	fcn = func(f la.Vector, dx, x float64, y la.Vector) {
		w.Fill(0)
		for i := start; i < endp1; i++ {
			m := 2 * i
			n := m + 1
			ui := y[m]
			vi := y[n]
			prod := ui * ui * vi
			uim, vim := 1.0, 3.0
			uip, vip := 1.0, 3.0
			if i == 0 {
				uip = y[m+2]
				vip = y[n+2]
			} else if i == N-1 {
				uim = y[m-2]
				vim = y[n-2]
			} else {
				uim = y[m-2]
				vim = y[n-2]
				uip = y[m+2]
				vip = y[n+2]
			}
			w[m] = 1 - 4*ui + prod + gamma*(uim-2*ui+uip)
			w[n] = 0 + 3*ui - prod + gamma*(vim-2*vi+vip)
		}
		comm.AllReduceSum(f, w)
	}

	// Jacobian
	jac = func(dfdy *la.Triplet, dx, x float64, y la.Vector) {
		if dfdy.Max() == 0 {
			dfdy.Init(ndim, ndim, (endp1-start)*8)
		}
		dfdy.Start()
		for i := start; i < endp1; i++ {
			m := 2 * i
			n := m + 1
			dfdy.Put(m, m, -4+2*y[m]*y[m+1]-2*gamma)
			dfdy.Put(m, m+1, y[m]*y[m])
			dfdy.Put(n, n, -y[n-1]*y[n-1]-2*gamma)
			dfdy.Put(n, n-1, 3-2*y[n-1]*y[n])
			if i > 0 {
				dfdy.Put(m, m-2, gamma)
				dfdy.Put(n, n-2, gamma)
			}
			if i < N-1 {
				dfdy.Put(m, m+2, gamma)
				dfdy.Put(n, n+2, gamma)
			}
		}
	}

	// initial values
	for i := start; i < endp1; i++ {
		m := 2 * i
		n := m + 1
		x := float64(i+1) / float64(N+1)
		w[m] = 1 + math.Sin(2*math.Pi*x)
		w[n] = 3
	}
	yini = la.NewVector(ndim)
	comm.AllReduceSum(yini, w)
	return
}
