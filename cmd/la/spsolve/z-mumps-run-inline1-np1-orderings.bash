#!/bin/bash

FNKEY="inline_1"

ORDERINGS="\
  amd \
  amf \
  pord \
  qamd \
"

BIN="/tmp/spsolve-mumps"

for ord in $ORDERINGS; do
  echo
  echo "============================================================================"
  echo "running with ordering = $ord"
  echo "============================================================================"
  go build -o $BIN && $BIN -kind="mumps" -fnkey=$FNKEY -ordering=$ord
done
