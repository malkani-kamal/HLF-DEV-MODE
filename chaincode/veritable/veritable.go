package main

import (
	"encoding/json"
	"fmt"

	//	"response"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

//VeritableContract Smart Contract for Veritable Organization
type VeritableContract struct {
}

const (
	foramtDate = "2006-01-02"
)

//https://stackoverflow.com/questions/24809235/how-to-initialize-a-nested-struct

//User ... data structure will be passed by Java server
type User struct {
	UserVeritableId      string      `json:"userVeritableId"`
	UserEmailId          string      `json:"userEmailId"`
	UserType             int         `json:"userType"`
	UserBusinessName     string      `json:"userBusinessName"`
	UserFirstName        string      `json:"userFirstName"`
	UserLastName         string      `json:"userLastName"`
	UserName             string      `json:"userName"`
	UserAddress1         string      `json:"userAddress1"`
	UserAddress2         string      `json:"userAddress2"`
	UserPhone            string      `json:"userPhone"`
	TimeStamp            string      `json:"timeStamp"`
	City                 string      `json:"City"`
	State                string      `json:"State"`
	Zip                  string      `json:"Zip"`
	Country              string      `json:"Country"`
	ScanReference        string      `json:"scanReference"`
	NotaryPledge         int         `json:"notaryPledge"`
	NotaryPledgeTextHash string      `json:"notaryPledgeTextHash"`
	UserPhotoName        string      `json:"userPhotoName"`
	UserPhotoURL         string      `json:"userPhotoURL"`
	UserPhotoHash        string      `json:"userPhotoHash"`
	CreatedDate          string      `json:"createDate"`
	UpdatedDate          string      `json:"updatedDate"`
	DeletedDate          string      `json:"deletedDate"`
	Documents            []Documents `json:"documents"`
	UserStamps           Stamp       `json:"userStamps"`
}

//Documents struct to handle user documents

type Documents struct {
	UserId             string `json:"userId"`
	DocType            int    `json:"docType"`
	DocUniqueNum       string `json:"docUniqueNum"`
	DocName            string `json:"docName"`
	DocIssuedStateCode string `json:"docIssuedStateCode"`
	DocExpireDate      string `json:"docExpireDate"`
	DocAddressLine1    string `json:"docAddressLine1"`
	DocAddressLine2    string `json:"docAddressLine2"`
	DocTownName        string `json:"docTownName"`
	DocPostalCode      string `json:"docPostalCode"`
	DocPhone           string `json:"docPhone"`
	DocStatus          int    `json:"docStatus"`
	DocCounty          string `json:"docCounty"`
	DocImageName       string `json:"docImageName"`
	DocImageURL        string `json:"docImageURL"`
	DocImageHash       string `json:"docImageHash"`
	CreatedDate        string `json:"createDate"`
	UpdatedDate        string `json:"updatedDate"`
	DeletedDate        string `json:"deletedDate"`
}

//Stamp struct to handle stamp details
type Stamp struct {
	UserID        string `json:"userID"`
	DocID         string `json:"docID"`
	NotaryStampId string `json:"notaryStampId"`
	StampFileName string `json:"stampFileName"`
	StampName     string `json:"stampName"`
	StampURL      string `json:"stampURL"`
	StampFileHash string `json:"stampFileHash"`
	CreatedDate   string `json:"createDate"`
	UpdatedDate   string `json:"updatedDate"`
	DeletedDate   string `json:"deletedDate"`
}

//Init ... this function is used to intialize the smart contract
func (s *VeritableContract) Init(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("Init executed")
	return shim.Success([]byte("Chaincode Successfully initialized"))
}

//CreateUser ... this function is used to create Notary
func (s *VeritableContract) CreateUser(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	fmt.Println(args[0])

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 2 arguments 1. function name ")
	}
	fmt.Println(args[0])

	user := User{}
	userparseerr := json.Unmarshal([]byte(args[0]), &user)
	if userparseerr != nil {
		fmt.Println("Error occuring here 147")
		return shim.Error(userparseerr.Error())
	}

	resultsIterator, err1 := stub.GetQueryResult("{\"selector\":{\"doc_type\":\"User\",\"_id\":\"" + user.UserVeritableId + "\"}}")
	if err1 != nil {
		return shim.Error(err1.Error())
	}
	if resultsIterator.HasNext() {
		return shim.Error("Duplicate user with same veritable ID." + user.UserVeritableId)
	}

	userbytes, usermarserr := json.Marshal(user)
	if usermarserr != nil {

		fmt.Println("Error occuring here 153")
		return shim.Error(usermarserr.Error())
	}
	fmt.Println(user)
	err0 := stub.PutState(stub.GetTxID(), userbytes)
	if err0 != nil {
		return shim.Error(err0.Error())
	}

	//	return shim.Success(response.CreateSuccessResponse(0,stub.GetTxID(),		"User created successfully"))
	return shim.Success([]byte("Notary profile created successfully." + stub.GetTxID()))
}

//GetUserByKey get User details
func (s *VeritableContract) GetUserByKey(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	userAsBytes, _ := stub.GetState(args[0])

	fmt.Println(userAsBytes)

	return shim.Success(userAsBytes)
}

func (s *VeritableContract) addDocument(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	fmt.Println(args[0])
	fmt.Println(args[1])

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	userAsBytes, _ := stub.GetState(args[0])
	user := User{}
	userUnmarserr := json.Unmarshal(userAsBytes, &user)

	if userUnmarserr != nil {
		return shim.Error(userUnmarserr.Error())
	}

	doc := Documents{}
	docErr := json.Unmarshal([]byte(args[1]), &doc)

	if docErr != nil {
		return shim.Error(docErr.Error())
	}

	user.Documents = doc

	fmt.Println("Full record print before document update")

	fmt.Println(user)

	userbytes, usermarserr := json.Marshal(user)

	if usermarserr != nil {
		return shim.Error(usermarserr.Error())
	}

	err0 := stub.PutState(args[0], userbytes)
	if err0 != nil {
		return shim.Error(err0.Error())
	}

	/*return shim.Success(response.CreateSuccessResponse(0,
	stub.GetTxID(),
	"User created successfully"))*/
	//return shim.Success([]byte("Notary profile created successfully." + stub.GetTxID()))

	return shim.Success(userAsBytes)
}

//addStamp to handle stamp
func (s *VeritableContract) addStamp(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	fmt.Println(args[0])
	fmt.Println(args[1])

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	userAsBytes, _ := stub.GetState(args[0])
	user := User{}
	userUnmarserr := json.Unmarshal(userAsBytes, &user)

	if userUnmarserr != nil {
		return shim.Error(userUnmarserr.Error())
	}

	stamp := Stamp{}
	stampErr := json.Unmarshal([]byte(args[1]), &stamp)

	if stampErr != nil {
		return shim.Error(stampErr.Error())
	}

	user.UserStamps = stamp

	fmt.Println("Full record print before stamp update")

	fmt.Println(user)

	userbytes, usermarserr := json.Marshal(user)

	if usermarserr != nil {
		return shim.Error(usermarserr.Error())
	}

	err0 := stub.PutState(args[0], userbytes)
	if err0 != nil {
		return shim.Error(err0.Error())
	}

	/*return shim.Success(response.CreateSuccessResponse(0,
	stub.GetTxID(),
	"User created successfully"))*/
	//return shim.Success([]byte("Notary profile created successfully." + stub.GetTxID()))

	return shim.Success(userAsBytes)
}

//Invoke function
func (s *VeritableContract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fun, args := stub.GetFunctionAndParameters()
	if fun == "CreateUser" {
		fmt.Println("Error occured ==> here in invoke")
		//logger.Info("########### create docs ###########")
		return s.CreateUser(stub, args)
	} else if fun == "GetUserByKey" {
		fmt.Println("Calling get  ==> ")
		//logger.Info("########### create docs ###########")
		return s.GetUserByKey(stub, args)
	} else if fun == "addDocument" {
		fmt.Println("Calling addDocument ==> ")
		//logger.Info("########### create docs ###########")
		return s.addDocument(stub, args)
	} else if fun == "addStamp" {
		fmt.Println("Calling addStamp  ==> ")
		//logger.Info("########### create docs ###########")
		return s.addStamp(stub, args)
	}

	return shim.Error(fmt.Sprintf("Unknown action, check the first argument, must be one of 'delete', 'query', or 'move'. But got: %v", fun))

}

func main() {
	err := shim.Start(new(VeritableContract))

	if err != nil {
		fmt.Print(err)
	}
}
