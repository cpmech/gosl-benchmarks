// Copyright 2020 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bmark

import (
	"testing"
	"time"

	"github.com/cpmech/gosl/chk"
)

func TestTimeAndMemory01(tst *testing.T) {

	defer func() {
		ResetGlobalStopwatch()
	}()

	// verbose()
	chk.PrintTitle("TimeAndMemory01")

	time.Sleep(333 * time.Millisecond)
	r := MeasureTimeAndMemory(chk.Verbose)
	et, _ := time.ParseDuration(r.ElapsedTimeString)
	chk.Int64(tst, "duration", et.Milliseconds(), 333)
}
