// Copyright 2020 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bmark

import (
	"encoding/json"

	"github.com/cpmech/gosl/chk"
	"github.com/cpmech/gosl/io"
)

// TimeAndMemory holds time and memory results
type TimeAndMemory struct {
	ElapsedTimeNanoseconds int64
	ElapsedTimeString      string
	MemoryCurrentBytes     uint64
	MemoryTotalBytes       uint64
	MemoryCurrentMiB       uint64
	MemoryTotalMiB         uint64
}

// Save saves json file with results
func (o TimeAndMemory) Save(dirout, fnkey string) {
	b, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		chk.Panic("%v\n", err)
	}
	io.WriteBytesToFileVD(dirout, fnkey+".json", b)
}

// ReadTimeAndMemory reads json file with results
func ReadTimeAndMemory(filename string) (o *TimeAndMemory) {
	b := io.ReadFile(filename)
	o = new(TimeAndMemory)
	err := json.Unmarshal(b, o)
	if err != nil {
		chk.Panic("%v\n", err)
	}
	return
}
