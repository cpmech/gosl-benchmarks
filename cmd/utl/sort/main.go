// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"sort"
	"time"

	"github.com/cpmech/gosl/io"
	"github.com/cpmech/gosl/rnd"
	"github.com/cpmech/gosl/utl"
)

func main() {
	rnd.Init(0)

	// generate array sizes
	nonlinearSizes := utl.NonlinSpace(1e3, 1e6, 10, 5, false)
	arraySizes := make([]int, len(nonlinearSizes))
	for i, n := range nonlinearSizes {
		arraySizes[i] = int(n)
	}

	// using gosl utl.Qsort
	bench("utl_qsort", arraySizes, func(A []float64) { utl.Qsort(A) })

	// using go's sort function
	bench("go_sort", arraySizes, func(A []float64) { sort.Float64s(A) })
}

func bench(fnkey string, arraySizes []int, sorter func(A []float64)) {

	// buffer with results
	buf := new(bytes.Buffer)
	io.Ff(buf, "%8s %23s %23s %23s\n", "size", "tfwd", "tbwd", "trnd")

	for _, n := range arraySizes {

		// forward series
		Afwd := utl.LinSpace(0, float64(n-1), n)
		t0 := time.Now()
		sorter(Afwd)
		tfwd := float64(time.Now().Sub(t0).Nanoseconds())

		// backward series
		Abwd := utl.LinSpace(float64(n-1), 0, n)
		t0 = time.Now()
		sorter(Abwd)
		tbwd := float64(time.Now().Sub(t0).Nanoseconds())

		// random series
		ntrials := 10
		sum := 0.0
		for j := 0; j < ntrials; j++ {
			Arnd := utl.GetCopy(Afwd)
			rnd.Shuffle(Arnd)
			t0 = time.Now()
			sorter(Arnd)
			del := float64(time.Now().Sub(t0).Nanoseconds())
			sum += del
		}
		trnd := sum / float64(ntrials)

		// buffer results
		io.Ff(buf, "%8d %23.15e %23.15e %23.15e\n", n, tfwd, tbwd, trnd)
	}

	// save file
	io.WriteFileVD("./results", io.Sf("%s.res", fnkey), buf)
}
