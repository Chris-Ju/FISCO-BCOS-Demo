package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/FISCO-BCOS/go-sdk/common/hexutil"
	"github.com/FISCO-BCOS/go-sdk/crypto"
)

/*
private key:  546e55559497633390ef8c112cb9858a3e56451f31c6eb1be35788ee693884ee
publick key:  8a37ed99d5fe6045884210fdbc40b9dde10739549d7d10ec4d2d8302ba2d704b6578419b679bb2d413c726417282bd061e9884ac7f26ebea54bf424225fd9027
address:  0x26EB22a9F43c0C90E2a7728f6c01eCD392F75087
*/

func main() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("private key: ", hexutil.Encode(privateKeyBytes)[2:]) // privateKey in hex without "0x"

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("publick key: ", hexutil.Encode(publicKeyBytes)[4:]) // publicKey in hex without "0x"

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("address: ", address) // account address
}
