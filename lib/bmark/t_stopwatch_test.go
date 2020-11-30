// Copyright 2020 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bmark

import (
	"testing"
	"time"

	"github.com/cpmech/gosl/chk"
)

func TestStopwatch01(tst *testing.T) {

	defer func() {
		ResetGlobalStopwatch()
	}()

	// verbose()
	chk.PrintTitle("Stopwatch01")

	sw := StartNewStopwatch()
	time.Sleep(333 * time.Millisecond)
	sw.Stop("First step: ")

	dur := time.Now().Sub(sw.initialTime)
	chk.Int64(tst, "duration", dur.Milliseconds(), 333)

	sw.Reset()
	time.Sleep(222 * time.Millisecond)

	dur = time.Now().Sub(sw.initialTime)
	chk.Int64(tst, "duration", dur.Milliseconds(), 222)
}
