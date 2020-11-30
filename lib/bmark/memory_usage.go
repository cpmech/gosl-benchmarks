// Copyright 2020 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bmark

import (
	"runtime"

	"github.com/cpmech/gosl/io"
)

// Bytes2MiB converts bytes to MiB
func Bytes2MiB(b uint64) uint64 {
	return b / 1024 / 1024
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
// Input:
//   messagePrefix -- if given, will print message, otherwise nothing is printed
func MemoryUsage(messagePrefix ...string) (currentBytes, totalBytes uint64) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	currentBytes = m.Alloc
	totalBytes = m.TotalAlloc
	if len(messagePrefix) > 0 {
		prefix := messagePrefix[0]
		io.Pf("%smemory: current = %v Bytes = %v MiB", prefix, currentBytes, Bytes2MiB(currentBytes))
		io.Pf(", total = %v Bytes = %v MiB\n", totalBytes, Bytes2MiB(totalBytes))
	}
	return
}
