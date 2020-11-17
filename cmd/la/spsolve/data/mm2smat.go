package main

import (
	"_/lib/bmark"
	"gosl/io"
	"gosl/la"
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
	T.ReadSmat(io.Sf("%s.mtx", fnkey))

	io.Pf("saving matrix\n")
	T.WriteSmat("", fnkey, 1e-15, "", false, false)
}
