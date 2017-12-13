/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/mpc"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type MPCExampleChaincode struct {
}

func (t *MPCExampleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Init")

	return shim.Success(nil)
}

func (t *MPCExampleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Invoke")
	// Run function arg[0] as master/slave arg[1]
	// connecting to decorations["target"] on input
	// decorations["input"]

	args := stub.GetArgs()
	decorations := stub.GetDecorations()

	//function := string(args[0])
	master := true
	if args[1][0] == 0 {
		master = false
	}
	target := string(decorations["target"])
	input := decorations["input"]

	// Open channel
	channel := mpc.NewCommSCCChannel(stub)
	if master {
		// First send, then receive
		err := channel.Send(input, target)
		if err != nil {
			return shim.Error(err.Error())
		}

		res, err := channel.Receive(10)
		if err != nil {
			return shim.Error(err.Error())
		}
		fmt.Printf("got [%v] from [%s]", res, target)
	} else {
		// First receive, then send
		res, err := channel.Receive(10)
		if err != nil {
			return shim.Error(err.Error())
		}
		fmt.Printf("got [%v] from [%s]", res, target)

		err = channel.Send(input, target)
		if err != nil {
			return shim.Error(err.Error())
		}
	}

	return shim.Success(nil)
}

func main() {
	err := shim.Start(new(MPCExampleChaincode))
	if err != nil {
		fmt.Printf("Error starting schaincode: %s", err)
	}
}
