// Copyright 2020 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bmark

import (
	"encoding/json"

	"github.com/cpmech/gosl/chk"
	"github.com/cpmech/gosl/io"
)

// SolverResults holds solver results
type SolverResults struct {
	Symmetric        bool
	NumberOfRows     int // equal to the number of columns
	NumberOfNonZeros int // same as "pattern entries"
	StepReadMatrix   *TimeAndMemory
	StepInitialize   *TimeAndMemory
	StepFactorize    *TimeAndMemory
	StepSolve        *TimeAndMemory
}

// Save saves json file with results
func (o SolverResults) Save(dirout, fnkey string) {
	b, err := json.Marshal(o)
	if err != nil {
		chk.Panic("%v\n", err)
	}
	io.WriteBytesToFileVD(dirout, fnkey+".json", b)
}

// ReadSolverResults reads json file with results
func ReadSolverResults(filename string) (o *SolverResults) {
	b := io.ReadFile(filename)
	o = new(SolverResults)
	err := json.Unmarshal(b, o)
	if err != nil {
		chk.Panic("%v\n", err)
	}
	return
}
