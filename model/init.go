package model

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"

	"github.com/Chris-Ju/FISCO-BCOS-Demo/store"
	"github.com/FISCO-BCOS/go-sdk/accounts/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/common"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/FISCO-BCOS/go-sdk/crypto"
	"github.com/boltdb/bolt"
)

func init() {
	db, err := bolt.Open(DATAPATH, 0666, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Users"))
		if err != nil {
			return fmt.Errorf("create bucket: %v", err)
		}
		return nil
	})
}

func connect() (*store.Store, *bind.TransactOpts) {
	client, err := client.Dial(URL, groupID)
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
	return instance, opts
}
