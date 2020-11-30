// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"math/rand"
	"time"

	"github.com/cpmech/gosl/io"
	"github.com/cpmech/gosl/utl"
)

// NaiveMatMatMul is the naive version of la.MatMatMul
func NaiveMatMatMul(c [][]float64, α float64, a, b [][]float64) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b[0]); j++ {
			c[i][j] = 0.0
			for k := 0; k < len(a[0]); k++ {
				c[i][j] += α * a[i][k] * b[k][j]
			}
		}
	}
}

func main() {

	// run small values
	mValues := utl.IntRange3(2, 66, 2)
	nSamples := 1000
	bench("naive-dgemm-small", nSamples, mValues)

	io.Pl()

	// run larger values
	mValues = utl.IntRange3(16, 1424, 64)
	nSamples = 100
	bench("naive-dgemm-large", nSamples, mValues)
}

func bench(fnkey string, nSamples int, mValues []int) {

	// constants
	α := 1.0

	// export results
	buf := new(bytes.Buffer)
	io.Ff(buf, "%4s %4s %23s %23s\n", "m", "n", "Gflops", "DtMicros")

	// header
	io.Pf("   size   |     Naive dgemm        (Dt) \n")
	io.Pf("----------|-----------------------------\n")

	// run all sizes
	for _, m := range mValues {

		// Naive: allocate matrices
		a := make([][]float64, m)
		b := make([][]float64, m)
		c := make([][]float64, m)

		// Naive: generate random matrices
		for i := 0; i < m; i++ {
			a[i] = make([]float64, m)
			b[i] = make([]float64, m)
			c[i] = make([]float64, m)
			for j := 0; j < m; j++ {
				a[i][j] = rand.Float64() - 0.5
				b[i][j] = rand.Float64() - 0.5
				c[i][j] = rand.Float64() - 0.5
			}
		}

		// Naive: run benchmark
		t0 := time.Now()
		for l := 0; l < nSamples; l++ {
			NaiveMatMatMul(c, α, a, b)
		}

		// Naive: compute MFlops
		dt := time.Now().Sub(t0) / time.Duration(nSamples)
		dtMicros := float64(dt.Nanoseconds()) * 1e-3
		mflops := 2.0 * float64(m) * float64(m) * float64(m) / dtMicros
		gflops := mflops * 1e-3

		// print message
		io.Pf("%4d×%4d | %5.2f GFlops (%12v)\n", m, m, gflops, dt)

		// save buffer
		io.Ff(buf, "%4d %4d %23.15e %23.15e\n", m, m, gflops, dtMicros)
	}

	// save file
	io.WriteFileVD("./results", io.Sf("%s-%dsamples.res", fnkey, nSamples), buf)
}
