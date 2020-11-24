#!/bin/bash

BIN="/tmp/spsolve-mumps"

go build -o $BIN && mpirun -np 2 $BIN

