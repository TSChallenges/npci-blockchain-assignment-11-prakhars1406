#!/bin/bash

# Script to query chaincode

# Exit on first error
set -e

source ./envVar.sh

# Query LOC status
getLOCStatus() {
  set -x
  peer chaincode query -C locchannel -n loc -c '{"function":"GetLOCStatus","Args":["LOC123"]}'
  set +x
}

# Query LOC history
getLOCHistory() {
  set -x
  peer chaincode query -C locchannel -n loc -c '{"function":"GetLOCHistory","Args":["LOC123"]}'
  set +x
}

# Main execution
if [ "$1" = "status" ]; then
  getLOCStatus
elif [ "$1" = "history" ]; then
  getLOCHistory
else
  echo "Usage: ./query.sh [status|history]"
fi
