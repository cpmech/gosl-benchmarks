// Copyright 2020 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package bmark assists on benchmarks
package bmark

import (
	"gosl/io"
	"runtime"
	"time"
)

// Stopwatch assists measuring computational time
type Stopwatch struct {
	initialTime time.Time
}

// StartNewStopwatch returns a new Stopwatch already started
func StartNewStopwatch() (o *Stopwatch) {
	return &Stopwatch{initialTime: time.Now()}
}

// Stop stops stopwatch and print the elapsed time
func (o *Stopwatch) Stop(messagePrefix string) {
	io.Pf("%selapsed time = %v\n", messagePrefix, time.Now().Sub(o.initialTime))
}

// Reset resets stopwatch
func (o *Stopwatch) Reset() {
	o.initialTime = time.Now()
}

// MemoryUsage prints the current memory usage
// See: https://golang.org/pkg/runtime/#MemStats
//
//  Note:
//   The SI prefixes are for strict use with powers of 10, not powers of 2 as used by computers.
//   To solve this the International Electrotechnical Commission (IEC) came up with a new prefix
//   standard for powers of 2, known as binary prefixes. The solution was to add an "i" between
//   the initial letter and the "B". Full prefixes were also defined. As per the specification
//   of the prefixes, the power of 2 prefix counterpart to Kilo- is Kibi-. The prefix Kilo-
//   stands for 10^3 while Kibi- stands for 2^10. This means that a Kilobyte is 1000 bytes while
//   a Kibibyte is 1024 bytes. This might not seem like a big difference but when it gets to
//   Gigabytes versus Gibibytes the difference becomes much more noticeable
//   REFERENCE: https://blog.digilentinc.com/mib-vs-mb-whats-the-difference/
//
func MemoryUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	io.Pf("Current Alloc = %v MiB", getMiB(m.Alloc))
	io.Pf("\tTotalAlloc = %v MiB\n", getMiB(m.TotalAlloc))
}

func getMiB(b uint64) uint64 {
	return b / 1024 / 1024
}
