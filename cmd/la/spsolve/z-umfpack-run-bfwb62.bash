#!/bin/bash

FNKEY="bfwb62"

BIN="/tmp/spsolve-umfpack"

go build -o $BIN && $BIN -kind="umfpack" -fnkey=$FNKEY
