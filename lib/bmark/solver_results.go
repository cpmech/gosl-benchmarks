// Copyright 2020 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bmark

import (
	"encoding/json"
	"time"

	"github.com/cpmech/gosl/chk"
	"github.com/cpmech/gosl/io"
)

// SolverResults holds solver results
type SolverResults struct {
	Kind                  string
	Ordering              string
	MpiSize               int
	Symmetric             bool
	NumberOfRows          int
	NumberOfCols          int
	NumberOfNonZeros      int // same as "pattern entries"
	StepReadMatrix        *TimeAndMemory
	StepInitialize        *TimeAndMemory
	StepFactorize         *TimeAndMemory
	StepSolve             *TimeAndMemory
	TimeSolverNanoseconds int64  // total time Init+Fact+Solve (nanoseconds)
	TimeSolverString      string // total time Init+Fact+Solve (string representation)
}

// Save saves json file with results
func (o SolverResults) Save(dirout, fnkey string) {
	b, err := json.MarshalIndent(o, "", "  ")
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

// CalcSolverTime sums init+fact+solve time
func (o *SolverResults) CalcSolverTime() {
	if o.StepInitialize == nil || o.StepFactorize == nil || o.StepSolve == nil {
		chk.Panic("step time data must not be nil")
	}
	o.TimeSolverNanoseconds = o.StepInitialize.ElapsedTimeNanoseconds + o.StepFactorize.ElapsedTimeNanoseconds + o.StepSolve.ElapsedTimeNanoseconds
	o.TimeSolverString = time.Duration(o.TimeSolverNanoseconds).String()
}
