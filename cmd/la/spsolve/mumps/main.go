package main

import (
	"_/lib/bmark"
	"gosl/io"
)

func main() {
	io.Pf("hello\n")
	// bmark.StartNewStopwatch()
	// bmark.StartNewStopwatch()
	sw := bmark.StartNewStopwatch()
	sw.Stop("")
}
