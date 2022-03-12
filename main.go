package main

import (
	"blockchain/block"
	"blockchain/wallet"
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	// myBlockchainAddress := "my_blockchain_address"
	// blockChain := NewBlockchain(myBlockchainAddress)
	// blockChain.Print()

	// blockChain.AddTransaction("A", "B", 1.0)
	// blockChain.Mining()
	// blockChain.Print()

	// // previousHash := blockChain.LastBlock().Hash() these lines and the following commented block are replaced by the Mining function
	// // nonce := blockChain.ProofOfWork()
	// // blockChain.CreateBlock(nonce, previousHash)

	// blockChain.AddTransaction("C", "D", 2.0)
	// blockChain.AddTransaction("X", "Y", 3.0)
	// blockChain.Mining()
	// blockChain.Print()
	// // previousHash = blockChain.LastBlock().Hash()
	// // nonce = blockChain.ProofOfWork()
	// // blockChain.CreateBlock(nonce, previousHash)

	// fmt.Printf("C %.1f\n", blockChain.CalculateTotalAmount("my_blockchain_address"))
	// fmt.Printf("C %.1f\n", blockChain.CalculateTotalAmount("C"))
	// fmt.Printf("C %.1f\n", blockChain.CalculateTotalAmount("D"))

	// w := wallet.NewWallet()
	// fmt.Println(w.PrivateKeyStr())
	// fmt.Println(w.PublicKeyStr())
	// fmt.Println(w.BlockchainAddress())
	walletM := wallet.NewWallet()
	walletA := wallet.NewWallet()
	walletB := wallet.NewWallet()

	//Wallet
	t := wallet.NewTransaction(walletA.PrivateKey(), walletA.PublicKey(), walletA.BlockchainAddress(),
		walletB.BlockchainAddress(), 1.0)

	//Blockchain
	blockchain := block.NewBlockchain(walletM.BlockchainAddress())
	isAdded := blockchain.AddTransaction(walletA.BlockchainAddress(), walletB.BlockchainAddress(), 1.0,
		walletA.PublicKey(), t.GenerateSignature())
	fmt.Println("Added? ", isAdded)
	// fmt.Printf("signature %s\n", t.GenerateSignature())

	blockchain.Mining()
	blockchain.Print()
	fmt.Printf("A %.1f\n", blockchain.CalculateTotalAmount(walletA.BlockchainAddress()))
	fmt.Printf("B %.1f\n", blockchain.CalculateTotalAmount(walletB.BlockchainAddress()))
	fmt.Printf("M %.1f\n", blockchain.CalculateTotalAmount(walletM.BlockchainAddress()))
}
