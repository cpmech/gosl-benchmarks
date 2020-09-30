package main

import (
	"_/lib/bmark"
	"gosl/io"
)

func main() {
	io.Pf("hello\n")
	sw := bmark.StartNewStopwatch()
	sw.Stop("")
}
