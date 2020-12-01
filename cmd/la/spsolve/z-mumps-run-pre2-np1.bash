#!/bin/bash

FNKEY="pre2"

BIN="/tmp/spsolve-mumps"

go build -o $BIN && $BIN -kind="mumps" -fnkey=$FNKEY
