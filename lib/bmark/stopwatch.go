// Copyright 2020 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bmark

import (
	"time"

	"github.com/cpmech/gosl/io"
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
// Input:
//   messagePrefix -- if given, will print message, otherwise nothing is printed
func (o *Stopwatch) Stop(messagePrefix ...string) (elapsedTime time.Duration) {
	elapsedTime = time.Now().Sub(o.initialTime)
	if len(messagePrefix) > 0 {
		prefix := messagePrefix[0]
		io.Pf("%selapsed time = %v\n", prefix, elapsedTime)
	}
	return
}

// Reset resets stopwatch
func (o *Stopwatch) Reset() {
	o.initialTime = time.Now()
}
