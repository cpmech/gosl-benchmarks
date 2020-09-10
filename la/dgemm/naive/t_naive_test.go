// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"gosl/chk"
	"gosl/io"
	"gosl/utl"
	"testing"
)

func TestMatMatMul(tst *testing.T) {
	io.Verbose = true
	chk.Verbose = true
	chk.PrintTitle("NaiveMatMatMul")
	a := [][]float64{ // 2 x 3
		{1.0, 2.00, 3.0},
		{0.5, 0.75, 1.5},
	}
	b := [][]float64{ // 3 x 4
		{0.1, 0.5, 0.5, 0.75},
		{0.2, 2.0, 2.0, 2.0},
		{0.3, 0.5, 0.5, 0.5},
	}
	cref := [][]float64{
		{2.80, 12.0, 12.0, 12.50},
		{1.30, +5.0, +5.0, +5.25},
	}
	c := utl.Alloc(2, 4)
	NaiveMatMatMul(c, 2, a, b)
	chk.Deep2(tst, "c := 2⋅a⋅b", 1e-15, c, cref)
}
