/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Writing Your First Blockchain Application
 */

 package main

 /* Imports
  * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
  * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
  */
 import (
	 "bytes"
	 "encoding/json"
	 "fmt"
   "time"
	 "strconv"

	 "github.com/hyperledger/fabric/core/chaincode/shim"
	 // "github.com/hyperledger/fabric/core/chaincode/lib/cid"
	 sc "github.com/hyperledger/fabric/protos/peer"
 )

 // Define the Smart Contract structure
 type SmartContract struct {
 }

 // Define the car structure, with 4 properties.  Structure tags are used by encoding/json library
 type Invoice struct {
	 InvoiceNumber   string `json:"invoiceNumber"`
	 BilledTo        string `json:"billedTo"`
	 InvoiceDate     string `json:"invoiceDate"`
	 InvoiceAmount   float64 `json:"invoiceAmount"`
	 ItemDescription string `json:"itemDescription"`
	 GR              string `json:"gr"`
	 IsPaid          bool `json:"isPaid"`
	 PaidAmount      float64 `json:"paidAmount"`
	 Repaid          bool `json:"repaid"`
	 RepaymentAmount float64 `json:"repaymentAmount"`
 }

 /*
  * The Init method is called when the Smart Contract "fabcar" is instantiated by the blockchain network
  * Best practice is to have any Ledger initialization in separate function -- see initLedger()
  */
 func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	 return shim.Success(nil)
 }

 /*
  * The Invoke method is called as a result of an application request to run the Smart Contract "fabcar"
  * The calling application program has also specified the particular smart contract function to be called, with arguments
  */
 func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	 // Retrieve the requested Smart Contract function and arguments
	 function, args := APIstub.GetFunctionAndParameters()
	 // Route to the appropriate handler function to interact with the ledger appropriately
	 if function == "initLedger" {
		 return s.initLedger(APIstub)
	 } else if function == "newInvoice" {
		 return s.newInvoice(APIstub, args)
	 } else if function == "queryAllInvoices" {
		 return s.queryAllinvoices(APIstub)
	 } else if function == "getUser" {
		 return s.getUser(APIstub, args)
	 } else if function == "createInvoiceWithJsonInput" {
		 return s.createInvoiceWithJsonInput(APIstub, args)
	 } else if function == "isGoodsReceived" {
		 return s.isGoodsReceived(APIstub, args)
	 } else if function == "isPaidToSupplier" {
		 return s.isPaidToSupplier(APIstub, args)
	 } else if function == "isPaidToBank" {
		 return s.isPaidToBank(APIstub, args)
	 } else if function == "getInvoiceAuditHistory" {
		 return s.getInvoiceAuditHistory(APIstub, args)
	 }

	 return shim.Error("Invalid Smart Contract function name.")
 }

 func (s *SmartContract) queryCarsByOwner(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	 //TODO Write approriate code here
	 if len(args) < 1 {
		 return shim.Error("Incorrect number of arguments. Expecting 1")
	 }

	 //assign value of owner
	 owner := args[0]

	 //get and display value of owner
	 queryString := fmt.Sprintf("{\"selector\":{\"owner\":\"%s\"}}", owner)

	 //display error message if the query result is invalid
	 queryResults, err := getQueryResultForQueryString(APIstub, queryString)
	 if err != nil {
		 return shim.Error(err.Error())
	 }
	 return shim.Success(queryResults)

 }

 func getQueryResultForQueryString(APIstub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	 resultsIterator, err := APIstub.GetQueryResult(queryString)
	 if err != nil {
		 return nil, err
	 }
	 defer resultsIterator.Close()

	 // buffer is a JSON array containing QueryRecords
	 var buffer bytes.Buffer
	 buffer.WriteString("[")

	 bArrayMemberAlreadyWritten := false
	 for resultsIterator.HasNext() {
		 queryResponse, err := resultsIterator.Next()
		 if err != nil {
			 return nil, err
		 }
		 // Add a comma before array members, suppress it for the first array member
		 if bArrayMemberAlreadyWritten == true {
			 buffer.WriteString(",")
		 }
		 buffer.WriteString("{\"Key\":")
		 buffer.WriteString("\"")
		 buffer.WriteString(queryResponse.Key)
		 buffer.WriteString("\"")

		 buffer.WriteString(", \"Record\":")
		 buffer.WriteString(string(queryResponse.Value))
		 buffer.WriteString("}")
		 bArrayMemberAlreadyWritten = true
	 }
	 buffer.WriteString("]")

	 return buffer.Bytes(), nil
 }

 // func (s *SmartContract) queryCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 //
 // 	if len(args) != 1 {
 // 		return shim.Error("Incorrect number of arguments. Expecting 1")
 // 	}
 //
 // 	carAsBytes, _ := APIstub.GetState(args[0])
 // 	return shim.Success(carAsBytes)
 // }

 func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	 invoice := []Invoice{
		 Invoice{
       InvoiceNumber: "0001",
       BilledTo: "Samsung",
       InvoiceDate: "10-30-2014",
       InvoiceAmount: 10000.00,
       ItemDescription: "Screen",
       GR: "false",
       IsPaid: false,
       PaidAmount: 0.00,
       Repaid: false,
       RepaymentAmount: 0},
	 }

	 i := 0
	 for i < len(invoice) {
		 fmt.Println("i is ", i)
		 invoiceAsBytes, _ := json.Marshal(invoice[i])
		 APIstub.PutState("INV"+strconv.Itoa(i), invoiceAsBytes)
		 fmt.Println("Added", invoice[i])
		 i = i + 1
	 }

	 return shim.Success(nil)
 }

 func (s *SmartContract) newInvoice(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	 if len(args) != 6 {
		 return shim.Error("Incorrect number of arguments. Expecting 6")
	 }

   iAmount, _ := strconv.ParseFloat(args[4], 64)
  // pAmount, _ := strconv.ParseFloat(args[6], 64)
  // rpAmount, _ := strconv.ParseFloat(args[7], 64)

	 var invoice = Invoice{InvoiceNumber: args[1],
     BilledTo: args[2],
     InvoiceDate: args[3],
     InvoiceAmount: iAmount,
     ItemDescription: args[5],
     GR: "false",
     IsPaid: false,
     PaidAmount: 0,
      Repaid: false,
      RepaymentAmount: 0}

	 invoiceAsBytes, _ := json.Marshal(invoice)
	 APIstub.PutState(args[0], invoiceAsBytes)

	 return shim.Success(nil)
 }

 func (s *SmartContract) createInvoiceWithJsonInput(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	 if len(args) != 1 {
		 return shim.Error("Incorrect number of arguments. Expecting 5")
	 }
	 fmt.Println("args[1] > ", args[1])
	 invoiceAsBytes := []byte(args[1])
	 invoice := Invoice{}
	 err := json.Unmarshal(invoiceAsBytes, &invoice)

	 if err != nil {
		 return shim.Error("Error During Invoice Unmarshall")
	 }
	 APIstub.PutState(args[0], invoiceAsBytes)
	 return shim.Success(nil)
 }

 func (s *SmartContract) queryAllinvoices(APIstub shim.ChaincodeStubInterface) sc.Response {

	 startKey := "INV0"
	 endKey := "INV999"

	 resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	 if err != nil {
		 return shim.Error(err.Error())
	 }
	 defer resultsIterator.Close()

	 // buffer is a JSON array containing QueryResults
	 var buffer bytes.Buffer
	 buffer.WriteString("[")

	 bArrayMemberAlreadyWritten := false
	 for resultsIterator.HasNext() {
		 queryResponse, err := resultsIterator.Next()
		 if err != nil {
			 return shim.Error(err.Error())
		 }
		 // Add a comma before array members, suppress it for the first array member
		 if bArrayMemberAlreadyWritten == true {
			 buffer.WriteString(",")
		 }
		 buffer.WriteString("{\"Invoice\":")
		 buffer.WriteString("\"")
		 buffer.WriteString(queryResponse.Key)
		 buffer.WriteString("\"")

		 buffer.WriteString(", \"Record\":")
		 // Record is a JSON object, so we write as-is
		 buffer.WriteString(string(queryResponse.Value))
		 buffer.WriteString("}")
		 bArrayMemberAlreadyWritten = true
	 }
	 buffer.WriteString("]")

	 fmt.Printf("- queryAllInvoices:\n%s\n", buffer.String())

	 return shim.Success(buffer.Bytes())
 }

 func (s *SmartContract) isGoodsReceived(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	 if len(args) != 2 {
		 return shim.Error("Incorrect number of arguments. Expecting 2")
	 }

	 invoiceAsBytes, _ := APIstub.GetState(args[0])
	 invoice := Invoice{}

	 json.Unmarshal(invoiceAsBytes, &invoice)
	 invoice.GR = args[1]

	 invoiceAsBytes, _ = json.Marshal(invoice)
	 APIstub.PutState(args[0], invoiceAsBytes)

	 return shim.Success(nil)
 }

 func (s *SmartContract) isPaidToSupplier(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	 if len(args) != 2 {
		 return shim.Error("Incorrect number of arguments. Expecting 2")
	 }

	 invoiceAsBytes, _ := APIstub.GetState(args[0])
	 invoice := Invoice{}

   pAmount, _ := strconv.ParseFloat(args[1], 64)
   json.Unmarshal(invoiceAsBytes, &invoice)
   if pAmount < invoice.InvoiceAmount {
     invoice.PaidAmount = pAmount
     invoice.IsPaid = true
   } else {
     return shim.Error("Paid Amount should be less than Invoice Amount")
   }

   invoiceAsBytes, _ = json.Marshal(invoice)
   APIstub.PutState(args[0], invoiceAsBytes)
   return shim.Success(nil)
 }

 func (s *SmartContract) isPaidToBank(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	 if len(args) != 2 {
		 return shim.Error("Incorrect number of arguments. Expecting 2")
	 }

   invoiceAsBytes, _ := APIstub.GetState(args[0])
	 invoice := Invoice{}

   rpAmount, _ := strconv.ParseFloat(args[1], 64)
   json.Unmarshal(invoiceAsBytes, &invoice)
   if rpAmount > invoice.InvoiceAmount {
     invoice.RepaymentAmount = rpAmount
     invoice.Repaid = true
   } else {
     return shim.Error("Repayment Amount should be greater than Paid Amount")
   }

   invoiceAsBytes, _ = json.Marshal(invoice)
   APIstub.PutState(args[0], invoiceAsBytes)
   return shim.Success(nil)
 }


 func (s *SmartContract) getUser(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	 return shim.Success(nil)

 }

 func (s *SmartContract) getInvoiceAuditHistory(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

 	if len(args) < 1 {
 		return shim.Error("Incorrect number of arguments. Expecting 1")
 	}

 	invoice := args[0]

 	resultsIterator, err := APIstub.GetHistoryForKey(invoice)
 	if err != nil {
 		return shim.Error(err.Error())
 	}
 	defer resultsIterator.Close()

 	// buffer is a JSON array containing historic values for the car
 	var buffer bytes.Buffer
 	buffer.WriteString("[")

 	bArrayMemberAlreadyWritten := false
 	for resultsIterator.HasNext() {
 		response, err := resultsIterator.Next()
 		if err != nil {
 			return shim.Error(err.Error())
 		}
 		// Add a comma before array members, suppress it for the first array member
 		if bArrayMemberAlreadyWritten == true {
 			buffer.WriteString(",")
 		}
 		buffer.WriteString("{\"TxId\":")
 		buffer.WriteString("\"")
 		buffer.WriteString(response.TxId)
 		buffer.WriteString("\"")

 		buffer.WriteString(", \"Value\":")
 		buffer.WriteString(string(response.Value))

 		buffer.WriteString(", \"Timestamp\":")
 		buffer.WriteString("\"")
 		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
 		buffer.WriteString("\"")

 		buffer.WriteString("}")
 		bArrayMemberAlreadyWritten = true
 	}
 	buffer.WriteString("]")

 	return shim.Success(buffer.Bytes())
 }

 // The main function is only relevant in unit test mode. Only included here for completeness.
 func main() {

	 // Create a new Smart Contract
	 err := shim.Start(new(SmartContract))
	 if err != nil {
		 fmt.Printf("Error creating new Smart Contract: %s", err)
	 }
 }
