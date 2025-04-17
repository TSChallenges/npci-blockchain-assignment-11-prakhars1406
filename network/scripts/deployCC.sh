#!/bin/bash

# This script deploys the loc chaincode to the Fabric network

# Exit on first error
set -e

# Set environment variables for Org1
source ./envVar.sh

# Package the chaincode
echo "Packaging chaincode..."
peer lifecycle chaincode package loc.tar.gz --path ../chaincode --lang golang --label loc_1.0

# Install the chaincode
echo "Installing chaincode on peer0.org1..."
peer lifecycle chaincode install loc.tar.gz

# TODO: Approve chaincode definitions for all organizations
# TODO: Commit chaincode definition to the channel

echo "Chaincode deployment completed successfully."
