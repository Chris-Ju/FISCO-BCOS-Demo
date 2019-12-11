package main

import (
	"fmt"
	store "github.com/Chris-Ju/FISCO-BCOS-Demo/store"
	"github.com/FISCO-BCOS/go-sdk/accounts/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/crypto"
	"log"
)

/*
contract address:  0xBcA77A5eF6dfDB888Fd6a01B73D54ef92505e167
transaction hash:  0x538dd0ae9799ee9066d728f3104abeeb22d2ec4fccfa1fbb901221f79657784f
*/

func main() {
	groupID := uint(1)
	client, err := client.Dial("http://localhost:8545", groupID)
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := crypto.HexToECDSA("546e55559497633390ef8c112cb9858a3e56451f31c6eb1be35788ee693884ee")
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	address, tx, instance, err := store.DeployStore(auth, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("contract address: ", address.Hex()) // the address should be saved
	fmt.Println("transaction hash: ", tx.Hash().Hex())
	_ = instance
}
