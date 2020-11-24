#!/bin/bash

# fnkey = "atmosmodl" // fail
# fnkey = "Hamrle3" // fail

FNKEYS="\
  bfwb62 \
  inline_1 \
  audikw_1 \
  Flan_1565 \
  tmt_unsym \
  pre2 \
"

BIN="/tmp/spsolve-mumps"

for fnkey in $FNKEYS; do
  echo "... running $fnkey"
  go build -o $BIN
  $BIN -fnkey=$fnkey > ../results/"mumps_$fnkey.txt"
done
