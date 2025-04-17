#!/bin/bash

# Environment variables for Org1 (TataMotors)

export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="TataMotorsMSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/tatamotors.loc.com/peers/peer0.tatamotors.loc.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/tatamotors.loc.com/users/Admin@tatamotors.loc.com/msp
export CORE_PEER_ADDRESS=localhost:7051
export ORDERER_CA=${PWD}/organizations/ordererOrganizations/loc.com/orderers/orderer.loc.com/msp/tlscacerts/tlsca.loc.com-cert.pem

# Path to binary
export PATH=${PWD}/../bin:$PATH

# FABRIC_CFG_PATH
export FABRIC_CFG_PATH=${PWD}/configtx
