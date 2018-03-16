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
    Id int
    Address string
    PriKey string
    PubKey string
    Money int
//    State int
}

var userId int

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

func (t *SimpleChaincode) createUser(stub shim.ChaincodeStubInterface, args []byte) (User, error) {
    if len(args) != 1{
        return nil, errors.New("function createUser, args number expect 1")
    }
    
    var err error
    var money int

    money, err = strconv.Atoi(args[0])
    if err != nil {
        return nil, errors.New("function createUser, args is not a number")
    }

    address, priKey, pubKey := getEsdsa()

    user := User{Id:userId, Address:address, PriKey:priKey, PubKey:pubKey, Money:money} 
    err = writeUser(stub, user.address, user.money)
    if err != nil {
        return nil, errors.New("function createUser, write user error")
    }

    return user, nil
}

func writeUser(stub shim.ChaincodeStubInterface, address string, money int) error {
    err := stub.PutState(address, money)
    if err != nil {
        return errors.New("function writeUser, putState error")
    }

    return nil
}

func getEsdsa() (string, string, string){
    //just a test
    address := "address" + strconv.Itoa(userId)
    priKey := "priKey" + strconv.Itoa(userId)
    pubKey := "pubKey" + strconv.Itoa(userId)

    userId++

    return address, priKey, pubKey
}

func getAllUser(stub shim.ChaincodeStubInterface) error{
    if userId == 0 {
        fmt.Println("function getAllUser, there is no user login")
        return nil
    }
    
    var i = 0
    for i < userId {
        user, err:=getUserById(i)
        if err != nil {
            return errors.New("function getAllUser, userId is not found")
        }
        else {
            fmt.Println("userId:%d, userAddress:%s, userMonty:%s", i, user.Address, user.Money)
        }
    }
    return nil
}

func getUserById(stub shim.ChaincodeStubInterface, userId int) User {
    //getAddressById(userId)
    address :=  "address" + strconv.Itoa(userId)
    address := "address" + strconv.Itoa(userId)
    priKey := "priKey" + strconv.Itoa(userId)
   
    money, err := stub.GetState(address)
    if err != nil {
        return nil, errors.New("function getUserById, GetState Error")
    }

    user := User{Id:userId, Address:address, PriKey:priKey, PubKey:pubKey}

    return user, nil 
}
