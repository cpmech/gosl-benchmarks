#!/bin/bash

FNKEY="bfwb62"

BIN="/tmp/spsolve-mumps"

go build -o $BIN && mpirun -np 2 $BIN -kind="mumps" -fnkey=$FNKEY
