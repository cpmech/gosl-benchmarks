package main

import (
	"fmt"
	"gosl/chk"
	"gosl/io"
	"gosl/la"
	"runtime"
	"time"
)

func main() {
	// Print our starting memory usage (should be around 0mb)
	// PrintMemUsage()

	io.Pf("reading matrix\n")
	startTime := time.Now()
	T := new(la.Triplet)
	T.ReadSmat("../data/inline_1.mtx", true)
	io.Pf("elapsed time = %v\n", time.Now().Sub(startTime))
	// PrintMemUsage()

	m, n := T.Size()
	if m != n {
		chk.Panic("matrix must be square. m=%d, n=%d\n", m, n)
	}
	b := la.NewVector(m)
	PrintMemUsage()

	io.Pf("solving\n")
	startTime = time.Now()
	la.SpSolve(T, b)
	io.Pf("elapsed time = %v\n", time.Now().Sub(startTime))

	// Print our memory usage
	// PrintMemUsage()

	// Force GC to clear up, should see a memory drop
	// runtime.GC()
	// PrintMemUsage()
}

