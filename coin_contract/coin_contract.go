package main

import (
    "fmt"
    "error"
    "strconv"
    "github.com/hyperledger/fabric/core/chaincode/shim" 
 )

type SimpleChainCode struct{
}

type User struct {
    Name string
    Address string
    PriKey string
    PubKey string
    Money int
//    State int
}

type Transaction struct {
    ToPubKey string
    FromSign string
    Money int
    IsGenesis bool
    IsUsed bool
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string )([]byte, error){
    //if len(args) != 4 {
    //    return nil, errors.New("it need at least 4 args! it's %d", len(args)
    //}

   // var A, B string
   // var Aval, Bval int
   // var err error

   // A = args[0]
   // Aval, err = strconv.Atoi(args[1])
   // if err != nil {
   //     return nil,  errors.New("expect args[1] is a number but novalid")
   // }

   // B = args[2]
   // Aval, err = strconv.Atoi(args[3])
   // if err != nil {
   //     return nil,  errors.New("expect args[3] is a number but novalid")
   // } 

   // var addressA, priKeyA, pubKeyA string
   // var addressB, priKeyB, pubKeyB string

   // addressA, priKeyA, pubKeyA = getEscda()
   // addressB, priKeyB, pubKeyB = getEscda()
    if function == "createUser" {
        return t.createUser(stub, args)
    }

    return nil, nil
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string )([]byte, error){
     if function == {
        return 
     }      

     return nil, nil
}

func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string )([]byte, error){
     if function == {
        return 
     }      

     return nil, nil
}

