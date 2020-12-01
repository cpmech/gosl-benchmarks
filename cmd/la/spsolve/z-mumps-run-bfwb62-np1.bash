#!/bin/bash

FNKEY="bfwb62"

BIN="/tmp/spsolve-mumps"

go build -o $BIN && $BIN -kind="mumps" -fnkey=$FNKEY
