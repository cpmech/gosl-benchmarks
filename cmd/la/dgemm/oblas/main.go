// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"gosl/io"
	"gosl/la/oblas"
	"gosl/utl"
	"math/rand"
	"time"
)

func main() {

	// set number of threads
	oblas.SetNumThreads(1)

	// run small values
	mValues := utl.IntRange3(2, 66, 2)
	nSamples := 1000
	bench("oblas-dgemm-small", nSamples, mValues)

	io.Pl()

	// run larger values
	mValues = utl.IntRange3(16, 1424, 64)
	nSamples = 100
	bench("oblas-dgemm-large", nSamples, mValues)
}

func bench(fnkey string, nSamples int, mValues []int) {

	// constants
	α, β := 1.0, 0.0
	_, mMax := utl.IntMinMax(mValues)

	// Dgemm: allocate matrices
	a := make([]float64, mMax*mMax)
	b := make([]float64, mMax*mMax)
	c := make([]float64, mMax*mMax)

	// Dgemm: generate random matrices
	for j := 0; j < mMax; j++ {
		for i := 0; i < mMax; i++ {
			a[i+j*mMax] = rand.Float64() - 0.5
			b[i+j*mMax] = rand.Float64() - 0.5
			c[i+j*mMax] = rand.Float64() - 0.5
		}
	}

	// Dgemm: run first to "warm-up"
	oblas.Dgemm(false, false, 2, 2, 2, α, a, 2, b, 2, β, c, 2)
	oblas.Dgemm(false, false, 4, 4, 4, α, a, 4, b, 4, β, c, 4)
	oblas.Dgemm(false, false, 8, 8, 8, α, a, 8, b, 8, β, c, 8)

	// export results
	buf := new(bytes.Buffer)
	io.Ff(buf, "%4s %4s %23s %23s\n", "m", "n", "Gflops", "DtMicros")

	// header
	io.Pf("   size   |     Naive dgemm        (Dt) \n")
	io.Pf("----------|-----------------------------\n")

	// run all sizes
	for _, m := range mValues {

		// Dgemm: run benchmark
		t0 := time.Now()
		for l := 0; l < nSamples; l++ {
			oblas.Dgemm(false, false, m, m, m, α, a, m, b, m, β, c, m)
		}

		// Dgemm: compute MFlops
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
