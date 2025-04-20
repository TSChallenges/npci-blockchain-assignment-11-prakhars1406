package main

import (
	"encoding/json"
	"fmt"
	
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// LetterOfCredit defines the structure for the Letter of Credit
type LetterOfCredit struct {
	LOCID            string   `json:"locId"`
	Buyer            string   `json:"buyer"`
	Seller           string   `json:"seller"`
	IssuingBank      string   `json:"issuingBank"`
	AdvisingBank     string   `json:"advisingBank"`
	Amount           string   `json:"amount"`
	Currency         string   `json:"currency"`
	ExpiryDate       string   `json:"expiryDate"`
	GoodsDescription string   `json:"goodsDescription"`
	Status           string   `json:"status"`
	DocumentHashes   []string `json:"documentHashes"`
	History          []string `json:"history"`
}

// SmartContract provides functions for managing the Letter of Credit
type SmartContract struct {
	contractapi.Contract
}

// InitLedger initializes the chaincode (optional)
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	// TODO: Initialization code if needed
	return nil
}

// RequestLOC creates a new LoC request
func (s *SmartContract) RequestLOC(ctx contractapi.TransactionContextInterface, locID string, buyer string, seller string, 
	issuingBank string, advisingBank string, amount string, currency string, expiryDate string, goodsDescription string) error {
	mspId,err:=ctx.GetClientIdentity().GetMSPID()
	if err!=nil{
		fmt.Printf("error in getting msp id %v\n",err)
		return err
	}
	if mspId!="TataMotors"{
		fmt.Println("only TataMotors can request LOC")
		return fmt.Errorf("only TataMotors can request LOC")
	}
	
	existingData,err:=ctx.GetStub().GetState(locID)
	if err!=nil{
		fmt.Printf("error in getting world state %v\n",err)
		return err
	}
	if existingData!=nil{
		fmt.Printf("locId %s already exits\n",locID)
		return fmt.Errorf("locId %s already exits",locID)
	}


	var letterOfCredit LetterOfCredit
	letterOfCredit.LOCID=locID
	letterOfCredit.Buyer=buyer
	letterOfCredit.Seller=seller
	letterOfCredit.IssuingBank=issuingBank
	letterOfCredit.AdvisingBank=advisingBank
	letterOfCredit.Amount=amount
	letterOfCredit.Currency=currency
	letterOfCredit.ExpiryDate=expiryDate
	letterOfCredit.GoodsDescription=goodsDescription
	letterOfCredit.Status="Requested"
	var history=make([]string,1)
	history=append(history,"TataMotors|LOC_Requested")
	letterOfCredit.History=history
	letterOfCreditInByes,err:=json.Marshal(letterOfCredit)
	if err!=nil{
		fmt.Printf("error in marshalling %v\n",err)
		return err
	}

	err=ctx.GetStub().PutState(locID,letterOfCreditInByes)
	if err!=nil{
		fmt.Printf("error in adding to world state %v\n",err)
		return err
	}

	ctx.GetStub().SetEvent("LOC_Requested",letterOfCreditInByes)
	return nil
}

func (s *SmartContract) IssueLOC(ctx contractapi.TransactionContextInterface, locID string) error{
	mspId,err:=ctx.GetClientIdentity().GetMSPID()
	if err!=nil{
		fmt.Printf("error in getting msp id %v\n",err)
		return err
	}
	if mspId!="ICICIBank"{
		fmt.Println("only ICICIBank can issue LOC")
		return fmt.Errorf("only ICICIBank can issue LOC")
	}

	existingData,err:=ctx.GetStub().GetState(locID)
	if err!=nil{
		fmt.Printf("error in getting world state %v\n",err)
		return err
	}
	if existingData==nil{
		fmt.Printf("locId %s does not exits\n",locID)
		return fmt.Errorf("locId %s does not exits",locID)
	}
	var letterOfCredit LetterOfCredit
	err=json.Unmarshal(existingData,&letterOfCredit)
	if err!=nil{
		fmt.Printf("error in unmarshalling %v\n",err)
		return err
	}

	letterOfCredit.Status="Issued"
	letterOfCredit.History=append(letterOfCredit.History,"TataMotors|LOC_ISSUED")
	letterOfCreditInByes,err:=json.Marshal(letterOfCredit)
	if err!=nil{
		fmt.Printf("error in marshalling %v\n",err)
		return err
	}

	err=ctx.GetStub().PutState(locID,letterOfCreditInByes)
	if err!=nil{
		fmt.Printf("error in adding to world state %v\n",err)
		return err
	}

	ctx.GetStub().SetEvent("LOC_ISSUED",letterOfCreditInByes)
	return nil
}



func (s *SmartContract) AcceptLOC(ctx contractapi.TransactionContextInterface, locID string) error{
	mspId,err:=ctx.GetClientIdentity().GetMSPID()
	if err!=nil{
		fmt.Printf("error in getting msp id %v\n",err)
		return err
	}
	if mspId!="Tesla"{
		fmt.Println("only Tesla can accept LOC")
		return fmt.Errorf("only Tesla can accept LOC")
	}

	existingData,err:=ctx.GetStub().GetState(locID)
	if err!=nil{
		fmt.Printf("error in getting world state %v\n",err)
		return err
	}
	if existingData==nil{
		fmt.Printf("locId %s does not exits\n",locID)
		return fmt.Errorf("locId %s does not exits",locID)
	}
	var letterOfCredit LetterOfCredit
	err=json.Unmarshal(existingData,&letterOfCredit)
	if err!=nil{
		fmt.Printf("error in unmarshalling %v\n",err)
		return err
	}

	letterOfCredit.Status="Accepted"
	letterOfCredit.History=append(letterOfCredit.History,"Tesla|LOC_ACCEPTED")
	letterOfCreditInByes,err:=json.Marshal(letterOfCredit)
	if err!=nil{
		fmt.Printf("error in marshalling %v\n",err)
		return err
	}

	err=ctx.GetStub().PutState(locID,letterOfCreditInByes)
	if err!=nil{
		fmt.Printf("error in adding to world state %v\n",err)
		return err
	}

	ctx.GetStub().SetEvent("LOC_ACCEPTED",letterOfCreditInByes)
	return nil
}


func (s *SmartContract) RejectLOC(ctx contractapi.TransactionContextInterface, locID string) error{
	mspId,err:=ctx.GetClientIdentity().GetMSPID()
	if err!=nil{
		fmt.Printf("error in getting msp id %v\n",err)
		return err
	}
	if mspId!="Tesla"{
		fmt.Println("only Tesla can reject LOC")
		return fmt.Errorf("only Tesla can reject LOC")
	}

	existingData,err:=ctx.GetStub().GetState(locID)
	if err!=nil{
		fmt.Printf("error in getting world state %v\n",err)
		return err
	}
	if existingData==nil{
		fmt.Printf("locId %s does not exits\n",locID)
		return fmt.Errorf("locId %s does not exits",locID)
	}
	var letterOfCredit LetterOfCredit
	err=json.Unmarshal(existingData,&letterOfCredit)
	if err!=nil{
		fmt.Printf("error in unmarshalling %v\n",err)
		return err
	}

	letterOfCredit.Status="Rejected"
	letterOfCredit.History=append(letterOfCredit.History,"Tesla|LOC_REJECTED")
	letterOfCreditInByes,err:=json.Marshal(letterOfCredit)
	if err!=nil{
		fmt.Printf("error in marshalling %v\n",err)
		return err
	}

	err=ctx.GetStub().PutState(locID,letterOfCreditInByes)
	if err!=nil{
		fmt.Printf("error in adding to world state %v\n",err)
		return err
	}

	ctx.GetStub().SetEvent("LOC_REJECTED",letterOfCreditInByes)
	return nil
}



func (s *SmartContract) ShipGoods(ctx contractapi.TransactionContextInterface, locID ,documentHash string) error{
	mspId,err:=ctx.GetClientIdentity().GetMSPID()
	if err!=nil{
		fmt.Printf("error in getting msp id %v\n",err)
		return err
	}
	if mspId!="Tesla"{
		fmt.Println("only Tesla can ship LOC")
		return fmt.Errorf("only Tesla can ship LOC")
	}

	existingData,err:=ctx.GetStub().GetState(locID)
	if err!=nil{
		fmt.Printf("error in getting world state %v\n",err)
		return err
	}
	if existingData==nil{
		fmt.Printf("locId %s does not exits\n",locID)
		return fmt.Errorf("locId %s does not exits",locID)
	}
	var letterOfCredit LetterOfCredit
	err=json.Unmarshal(existingData,&letterOfCredit)
	if err!=nil{
		fmt.Printf("error in unmarshalling %v\n",err)
		return err
	}

	var docHash = make([]string,1)
	docHash=append(docHash,documentHash)
	letterOfCredit.DocumentHashes=docHash
	letterOfCredit.Status="Shipped"
	letterOfCredit.History=append(letterOfCredit.History,"Tesla|GOODS_SHIPPED")
	letterOfCreditInByes,err:=json.Marshal(letterOfCredit)
	if err!=nil{
		fmt.Printf("error in marshalling %v\n",err)
		return err
	}

	err=ctx.GetStub().PutState(locID,letterOfCreditInByes)
	if err!=nil{
		fmt.Printf("error in adding to world state %v\n",err)
		return err
	}

	ctx.GetStub().SetEvent("GOODS_SHIPPED",letterOfCreditInByes)
	return nil
}



func (s *SmartContract) VerifyDocuments(ctx contractapi.TransactionContextInterface, locID string) error{
	mspId,err:=ctx.GetClientIdentity().GetMSPID()
	if err!=nil{
		fmt.Printf("error in getting msp id %v\n",err)
		return err
	}
	if mspId!="ChaseBank"{
		fmt.Println("only ChaseBank can verify doc")
		return fmt.Errorf("only ChaseBank can verify doc")
	}

	existingData,err:=ctx.GetStub().GetState(locID)
	if err!=nil{
		fmt.Printf("error in getting world state %v\n",err)
		return err
	}
	if existingData==nil{
		fmt.Printf("locId %s does not exits\n",locID)
		return fmt.Errorf("locId %s does not exits",locID)
	}
	var letterOfCredit LetterOfCredit
	err=json.Unmarshal(existingData,&letterOfCredit)
	if err!=nil{
		fmt.Printf("error in unmarshalling %v\n",err)
		return err
	}

	letterOfCredit.Status="Verified"
	letterOfCredit.History=append(letterOfCredit.History,"ChaseBank|DOCUMENTS_VERIFIED")
	letterOfCreditInByes,err:=json.Marshal(letterOfCredit)
	if err!=nil{
		fmt.Printf("error in marshalling %v\n",err)
		return err
	}

	err=ctx.GetStub().PutState(locID,letterOfCreditInByes)
	if err!=nil{
		fmt.Printf("error in adding to world state %v\n",err)
		return err
	}

	ctx.GetStub().SetEvent("DOCUMENTS_VERIFIED",letterOfCreditInByes)
	return nil
}


func (s *SmartContract) ReleasePayment(ctx contractapi.TransactionContextInterface, locID string) error{
	mspId,err:=ctx.GetClientIdentity().GetMSPID()
	if err!=nil{
		fmt.Printf("error in getting msp id %v\n",err)
		return err
	}
	if mspId!="ICICIBank"{
		fmt.Println("only ICICIBank can release payment")
		return fmt.Errorf("only ICICIBank can release payment")
	}

	existingData,err:=ctx.GetStub().GetState(locID)
	if err!=nil{
		fmt.Printf("error in getting world state %v\n",err)
		return err
	}
	if existingData==nil{
		fmt.Printf("locId %s does not exits\n",locID)
		return fmt.Errorf("locId %s does not exits",locID)
	}
	var letterOfCredit LetterOfCredit
	err=json.Unmarshal(existingData,&letterOfCredit)
	if err!=nil{
		fmt.Printf("error in unmarshalling %v\n",err)
		return err
	}

	letterOfCredit.Status="Paid"
	letterOfCredit.History=append(letterOfCredit.History,"ICICIBank|PAYMENT_RELEASED")
	letterOfCreditInByes,err:=json.Marshal(letterOfCredit)
	if err!=nil{
		fmt.Printf("error in marshalling %v\n",err)
		return err
	}

	err=ctx.GetStub().PutState(locID,letterOfCreditInByes)
	if err!=nil{
		fmt.Printf("error in adding to world state %v\n",err)
		return err
	}

	ctx.GetStub().SetEvent("PAYMENT_RELEASED",letterOfCreditInByes)
	return nil
}

func (s *SmartContract) GetLOCHistory(ctx contractapi.TransactionContextInterface, locID string) ([]string, error) {

	var loc = LetterOfCredit{}
	var err error

	locData, err := ctx.GetStub().GetState(locID)
	if err != nil || locData == nil {
		return loc.History, err
	}

	err = json.Unmarshal(locData, &loc)
	if err != nil {
		return loc.History, err
	}

	return loc.History, nil
}

func (s *SmartContract) GetLOCStatus(ctx contractapi.TransactionContextInterface, locID string) (string, error) {

	var loc = LetterOfCredit{}
	var err error

	locData, err := ctx.GetStub().GetState(locID)
	if err != nil || locData == nil {
		return loc.Status, err
	}

	err = json.Unmarshal(locData, &loc)
	if err != nil {
		return loc.Status, err
	}

	return loc.Status, nil
}

// TODO: Implement other functions here (IssueLOC, AcceptLOC, etc.)

func main() {
	chaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		fmt.Printf("Error creating loc chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting loc chaincode: %s", err.Error())
	}
}
