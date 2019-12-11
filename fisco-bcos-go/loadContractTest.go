package main

import (
	"fmt"
	store "github.com/Chris-Ju/FISCO-BCOS-Demo/store"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/common"
	"log"
)

func main() {
	groupID := uint(1)
	client, err := client.Dial("http://localhost:8545", groupID)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0xBcA77A5eF6dfDB888Fd6a01B73D54ef92505e167")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("contract is loaded")
	_ = instance
}
