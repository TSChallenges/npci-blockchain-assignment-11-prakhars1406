#!/bin/bash

# Script to invoke chaincode functions

# Exit on first error
set -e

source ./envVar.sh
parsePeerConnectionParameters() {
  PEER_CONN_PARMS=""
  PEERS=""
  
  # Add all peer addresses
  PEER_CONN_PARMS="--peerAddresses peer0.tatamotors.loc.com:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/tatamotors.loc.com/peers/peer0.tatamotors.loc.com/tls/ca.crt"
  PEERS="peer0.tatamotors.loc.com"
  
  # Add other peers as needed (Tesla, ICICIBank, ChaseBank)
}

# Example function to request LOC
requestLOC() {
  parsePeerConnectionParameters
  set -x
  peer chaincode invoke -o orderer.loc.com:7050 --tls true --cafile $ORDERER_CA -C locchannel -n loc $PEER_CONN_PARMS -c '{"function":"RequestLOC","Args":["LOC123","TataMotors","Tesla","ICICIBank","ChaseBank","500000","USD","20231231","Car Batteries"]}'
  set +x
}

# Add other functions as needed (issueLOC, acceptLOC, etc.)

# Main execution
if [ "$1" = "request" ]; then
  requestLOC
elif [ "$1" = "issue" ]; then
  issueLOC
# Add other command cases
else
  echo "Usage: ./invoke.sh [request|issue|accept|ship|verify|release]"
fi
