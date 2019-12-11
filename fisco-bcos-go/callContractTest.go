package main

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"

	store "github.com/Chris-Ju/FISCO-BCOS-Demo/store"
	"github.com/FISCO-BCOS/go-sdk/accounts/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/common"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/FISCO-BCOS/go-sdk/crypto"
)

func main() {
	groupID := uint(1)
	client, err := client.Dial("http://localhost:8545", groupID)
	if err != nil {
		log.Fatal(err)
	}

	// load the contract
	address := common.HexToAddress("0xBcA77A5eF6dfDB888Fd6a01B73D54ef92505e167")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	opts := &bind.TransactOpts{
		From: common.HexToAddress("0x26EB22a9F43c0C90E2a7728f6c01eCD392F75087"),
		Signer: func(signer types.RawSigner, address common.Address, tx *types.RawTransaction) (*types.RawTransaction, error) {
			if address != common.HexToAddress("0x26EB22a9F43c0C90E2a7728f6c01eCD392F75087") {
				return nil, errors.New("not authorized to sign this account")
			}
			signature, err := crypto.Sign(signer.Hash(tx).Bytes(), &ecdsa.PrivateKey{PublicKey: ecdsa.PublicKey{}})
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(signer, signature)
		},
	}
	buf, err := instance.SelectCompany(opts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(buf)
}
