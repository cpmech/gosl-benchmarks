// Copyright 2020 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bmark

import (
	"testing"
	"time"

	"github.com/cpmech/gosl/chk"
)

func TestSolverResults01(tst *testing.T) {

	defer func() {
		ResetGlobalStopwatch()
	}()

	// verbose()
	chk.PrintTitle("SolverResults01")

	time.Sleep(333 * time.Millisecond)
	stepReadMatrix := MeasureTimeAndMemory(false)

	time.Sleep(444 * time.Millisecond)
	stepInitialize := MeasureTimeAndMemory(false)

	time.Sleep(555 * time.Millisecond)
	stepFactorize := MeasureTimeAndMemory(false)

	time.Sleep(666 * time.Millisecond)
	stepSolve := MeasureTimeAndMemory(false)

	r0 := &SolverResults{
		Symmetric:        true,
		NumberOfRows:     123,
		NumberOfNonZeros: 456,
		StepReadMatrix:   stepReadMatrix,
		StepInitialize:   stepInitialize,
		StepFactorize:    stepFactorize,
		StepSolve:        stepSolve,
	}
	r0.Save("/tmp/", "solver-res-01")

	r1 := ReadSolverResults("/tmp/solver-res-01.json")

	etReadMatrix, _ := time.ParseDuration(r1.StepReadMatrix.ElapsedTimeString)
	chk.Int64(tst, "etReadMatrix", etReadMatrix.Milliseconds(), 333)

	etInitialize, _ := time.ParseDuration(r1.StepInitialize.ElapsedTimeString)
	chk.Int64(tst, "etInitialize", etInitialize.Milliseconds(), 444)

	etFactorize, _ := time.ParseDuration(r1.StepFactorize.ElapsedTimeString)
	chk.Int64(tst, "etFactorize", etFactorize.Milliseconds(), 555)

	etSolve, _ := time.ParseDuration(r1.StepSolve.ElapsedTimeString)
	chk.Int64(tst, "etSolve", etSolve.Milliseconds(), 666)
}

func TestSolverResults02(tst *testing.T) {

	defer func() {
		ResetGlobalStopwatch()
	}()

	// verbose()
	chk.PrintTitle("SolverResults02")

	time.Sleep(333 * time.Millisecond)
	stepReadMatrix := MeasureTimeAndMemory(false)

	time.Sleep(400 * time.Millisecond)
	stepInitialize := MeasureTimeAndMemory(false)

	time.Sleep(400 * time.Millisecond)
	stepFactorize := MeasureTimeAndMemory(false)

	time.Sleep(200 * time.Millisecond)
	stepSolve := MeasureTimeAndMemory(false)

	r0 := &SolverResults{
		Symmetric:        true,
		NumberOfRows:     123,
		NumberOfNonZeros: 456,
		StepReadMatrix:   stepReadMatrix,
		StepInitialize:   stepInitialize,
		StepFactorize:    stepFactorize,
		StepSolve:        stepSolve,
	}

	factor := int64(1e6) // nano to milli
	r0.CalcSolverTime()
	chk.Int64(tst, "TimeSolver [s]", r0.TimeSolverNanoseconds/factor, 1000)
	chk.String(tst, r0.TimeSolverString[:3], "1.0")
}
