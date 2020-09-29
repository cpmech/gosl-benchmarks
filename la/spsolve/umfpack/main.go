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
	PrintMemUsage()

	io.Pf("reading matrix\n")
	startTime := time.Now()
	T := new(la.Triplet)
	T.ReadSmat("../data/inline_1.mtx", true)
	io.Pf("elapsed time = %v\n", time.Now().Sub(startTime))
	PrintMemUsage()

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
	PrintMemUsage()

	// Force GC to clear up, should see a memory drop
	runtime.GC()
	PrintMemUsage()
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
// From here: https://golangcode.com/print-the-current-memory-usage/
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

// The SI prefixes are for strict use with powers of 10, not powers of 2 as used by computers. To solve this the International Electrotechnical Commission (IEC) came up with a new prefix standard for powers of 2, known as binary prefixes. The solution was to add an “i” between the initial letter and the “B”. Full prefixes were also defined. As per the specification of the prefixes, the power of 2 prefix counterpart to Kilo- is Kibi-. The prefix Kilo- stands for 10^3 while Kibi- stands for 2^10. This means that a Kilobyte is 1000 bytes while a Kibibyte is 1024 bytes. This might not seem like a big difference but when it gets to Gigabytes versus Gibibytes the difference becomes much more noticeable. The full prefix chart is in the image below.
// from here: https://blog.digilentinc.com/mib-vs-mb-whats-the-difference/
