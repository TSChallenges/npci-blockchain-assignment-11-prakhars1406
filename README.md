# NPCI-Blockchain-Assignment-11 :  Letter of Credit System [10 Marks]

## Assignment Objective

Implement a blockchain-based Letter of Credit (LoC) management system using Hyperledger Fabric to ensure transparency, authenticity, and efficiency in international trade financing. Develop a Hyperledger Fabric network and smart contract ( `locChaincode.go` ) to manage the lifecycle of Letters of Credit between real-world trade participants.

## You are required to

* Design a Fabric network with real trade participants (importers, exporters,banks).
* Implement **chaincode** to handle LoC issuance, acceptance, shipment, document submission, and payment release.
* Enable **role-based access control** for each participant.
* Simulate the full lifecycle of a trade transaction involving an LoC.

## Prerequisites

1. GitHub Account (for code submission).
2. Docker, Docker Compose installed.
3. Go programming language knowledge.
4. Familiarity with Hyperledger Fabric concepts (MSPs, Peers, Orderers, Chaincode, Channels).
5. Clone the `fabric-samples` repo: [https://github.com/hyperledger/fabric-samples](https://github.com/hyperledger/fabric-samples)

##  1. Setup Hyperledger Fabric Network

Use `fabric-samples/test-network` as a base and customize it with the following organizations:

### Organizations

| Organization | Role              | Description                                 |
|--------------|-------------------|---------------------------------------------|
| TataMotors   | Importer          | Indian automobile company (buyer)           |
| Tesla        | Exporter          | US-based electric car parts supplier        |
| ICICIBank    | Importer’s Bank   | Issues the LoC on behalf of TataMotors      |
| ChaseBank    | Exporter’s Bank   | Advises the LoC to Tesla                    |
| OrdererOrg   | Orderer           | One orderer node for transaction ordering   |

### Channel

- Create a single Fabric channel: `locchannel`

### Each Org Must Have:

- One peer node
- Certificate Authorities (optional: use crypto-generator or Fabric CA)
- MSP and Admin identity

## 2. Implement Chaincode: `locChaincode.go`

Develop a Go chaincode supporting the LoC lifecycle:

### Letter of Credit Lifecycle Operations

| Function         | Description                                                | Who Can Invoke   |
|------------------|------------------------------------------------------------|------------------|
| `RequestLOC`     | TataMotors requests LoC from ICICIBank                     | TataMotors       |
| `IssueLOC`       | ICICIBank issues LoC and sends it to ChaseBank             | ICICIBank        |
| `AcceptLOC`      | Tesla reviews and accepts the LoC                          | Tesla            |
| `ShipGoods`      | Tesla updates shipment status and submits shipping docs    | Tesla            |
| `VerifyDocuments`| ChaseBank verifies submitted documents                     | ChaseBank        |
| `ReleasePayment` | ICICIBank releases payment after document verification     | ICICIBank        |
| `GetLOCHistory`  | Trace full history of LoC (immutable ledger)               | All              |
| `GetLOCStatus`   | View current status and owner of the LoC                   | All              |
| `RejectLOC` *(optional)* | Tesla can reject the LoC due to terms mismatch    | Tesla            |

### Sample LoC Object Model

```go
type LetterOfCredit struct {
    LOCID            string   `json:"locId"`
    Buyer            string   `json:"buyer"`        // TataMotors
    Seller           string   `json:"seller"`       // Tesla
    IssuingBank      string   `json:"issuingBank"`  // ICICIBank
    AdvisingBank     string   `json:"advisingBank"` // ChaseBank
    Amount           string   `json:"amount"`
    Currency         string   `json:"currency"`
    ExpiryDate       string   `json:"expiryDate"`
    GoodsDescription string   `json:"goodsDescription"`
    Status           string   `json:"status"`       // Requested, Issued, Accepted, Shipped, Verified, Paid, Rejected
    DocumentHashes   []string `json:"documentHashes"`
    History          []string `json:"history"`      // Transaction history logs
}
```

## 3. Deploy and Test the Scenario

### Scenario Flow:
1. TataMotors requests a LoC for $500,000 USD to import car batteries.
2. ICICIBank issues the LoC to ChaseBank.
Blockchain Assignment: Letter of Credit System 3
3. Tesla accepts the LoC terms and ships the goods.
4. Tesla submits shipping documents (as document hashes).
5. ChaseBank verifies documents and notifies ICICIBank.
6. ICICIBank releases payment to Tesla.
7. Query the final status and transaction history.

## 4. Test Scripts

Create the following helper scripts:

* `invoke.sh` – Script to invoke chaincode functions for testing lifecycle
* `query.sh` – Script to query LoC status or history
* `deployCC.sh` – Deploy `locChaincode.go` to all peers
* `envVar.sh` – Helper for setting environment variables for each org

## 5. Access Control & Validation

Ensure the following:

* Only **TataMotors** can initiate `RequestLOC`
* Only **ICICIBank** can invoke `IssueLOC` and `ReleasePayment`
* Only **Tesla** can `AcceptLOC` and `ShipGoods`
* Only **ChaseBank** can `VerifyDocuments`
* All participants can run queries for `GetLOCHistory` or `GetLOCStatus`

## 6. Sample Chaincode Events (optional)

Emit events like:

* `LOC_ISSUED`
* `GOODS_SHIPPED`
* `DOCUMENTS_VERIFIED`
* `PAYMENT_RELEASED`

This can help in integrating with external notification systems

## Deliverables

Your GitHub repo must include:

* Customized network setup with org folders
* `locChaincode.go` with documented code
* Deployment and test scripts ( `invoke.sh` , `query.sh` , `deployCC.sh` )
* Sample command logs or screenshots for:
  
    * LoC lifecycle steps
    * Chaincode queries
    * Role-based access control working as
* Submit screenshots to SubmissionImages folder

