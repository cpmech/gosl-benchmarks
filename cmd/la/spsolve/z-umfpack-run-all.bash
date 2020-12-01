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

BIN="/tmp/spsolve-umfpack"

for fnkey in $FNKEYS; do
  echo
  echo "============================================================================"
  echo "running with fnkey = $fnkey"
  echo "============================================================================"
  go build -o $BIN && $BIN -kind="umfpack" -fnkey=$fnkey
done
