package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var infura = "Your infura url"

func main() {
	client, err := ethclient.DialContext(context.Background(), infura) //DialContext connects a client to the given URL with context.
	if err != nil {
		fmt.Println("Error to create a ether client", err)

	}
	defer client.Close()
	block, err := client.BlockByNumber(context.Background(), nil) //Returns a Current Block from the Canonical Chain.
	if err != nil {
		log.Fatalf("Error to get a block:%v", err)

	}
	fmt.Println("The Block Number is:", block.Number())
	account := common.HexToAddress("your address")                                  //Returns your eth address
	balance, err := client.BalanceAt(context.Background(), account, block.Number()) //Returns the Balance from your account
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The Balance in your account is:", balance)

	privatekey, err := crypto.GenerateKey() //Returns the private key

	if err != nil {
		log.Fatal(err)
	}
	privatekey_bytes := crypto.FromECDSA(privatekey) //Private key is encrypted using Elliptic Curve Cryptography and it is returned

	fmt.Println("Your Private Key is:", hexutil.Encode(privatekey_bytes))
	publickey := privatekey.Public()

	public_encrypted, ok := publickey.(*ecdsa.PublicKey) //public key as well.

	if !ok {
		fmt.Println("Public key is not of *&ecdsa.PublicKey")
	}
	publickey_bytes := crypto.FromECDSA(public_encrypted)
	fmt.Println("Public key is:", hexutil.Encode(publickey_bytes))
	address := crypto.PubkeyToAddress(*publickey_bytes).Hex()
	fmt.Println("address:", address)

	blockno := big.NewInt(123333)

	Block, err := client.BlockByNumber(context.Background(), blockno) //Returns the Entire current Block.
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The nonce of the block:", Block.Number().Uint64())
	fmt.Println("The time when block is minned:", Block.Time())
	fmt.Println("The total transactions occured:", len(Block.Transactions()))

}
