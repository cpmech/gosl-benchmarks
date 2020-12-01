// Copyright 2020 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bmark

import (
	"time"

	"github.com/cpmech/gosl/mpi"
)

var globalSW *Stopwatch

func init() {
	globalSW = StartNewStopwatch()
}

// ResetGlobalStopwatch resets global stopwatch
func ResetGlobalStopwatch() {
	globalSW.Reset()
}

// MeasureTimeAndMemory shows elapsed time and memory usage and resets global StopWatch
// Input:
//   comm -- optional: communicator so only Rank#0 will print
func MeasureTimeAndMemory(doPrint bool, comm ...*mpi.Communicator) (o *TimeAndMemory) {
	if len(comm) > 0 {
		if comm[0] != nil {
			if comm[0].Rank() != 0 {
				doPrint = false
			}
		}
	}
	var elapsedTime time.Duration
	var currentBytes, totalBytes uint64
	if doPrint {
		elapsedTime = globalSW.Stop("... ")
		currentBytes, totalBytes = MemoryUsage("... ")
	} else {
		elapsedTime = globalSW.Stop()
		currentBytes, totalBytes = MemoryUsage()
	}
	globalSW.Reset()
	return &TimeAndMemory{
		ElapsedTimeNanoseconds: elapsedTime.Nanoseconds(),
		ElapsedTimeString:      elapsedTime.String(),
		MemoryCurrentBytes:     currentBytes,
		MemoryTotalBytes:       totalBytes,
		MemoryCurrentMiB:       Bytes2MiB(currentBytes),
		MemoryTotalMiB:         Bytes2MiB(totalBytes),
	}
}
