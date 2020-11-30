// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"_/lib/bmark"

	"github.com/cpmech/gosl/io"
	"github.com/cpmech/gosl/la"
)

func main() {
	fnkey := "tmt_unsym"

	sw := bmark.StartNewStopwatch()
	defer func() {
		bmark.MemoryUsage()
		sw.Stop()
	}()

	io.Pf("reading matrix\n")
	T := new(la.Triplet)
	T.ReadSmat(io.Sf("%s.mtx", fnkey), true, nil)

	io.Pf("saving matrix\n")
	T.WriteSmat("", fnkey, 1e-15, "", false, false)
}
