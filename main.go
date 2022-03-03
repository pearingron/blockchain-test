package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

func NewBlock(nonce int, previousHash string) *Block { //generates a new block
	b := new(Block) //stores a pointer to this Block struct in b
	b.nonce = nonce
	b.previousHash = previousHash
	b.timestamp = time.Now().UnixNano()
	return b
	// alternative method
	// return &Block{
	// 	timestamp: time.Now().UnixNano(),
	// }
}

func (b *Block) Print() {
	fmt.Printf("nonce	%d\n", b.nonce)
	fmt.Printf("previous_hash	%s\n", b.previousHash)
	fmt.Printf("timestamp	%d\n", b.timestamp)
	fmt.Printf("transactions	%s\n", b.transactions)
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	fmt.Println(string(m))
	return sha256.Sum256([]byte(m))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64
		Nonce        int
		PreviousHash string
		Transactions []string
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "Init hash")
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash) //creating a pointer to a new block
	bc.chain = append(bc.chain, b)     //adds pointer to the generated block to the chain
	return b
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain { //loop through every block within the blockchain
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	block := &Block{nonce: 1}
	fmt.Printf("%x\n", block.Hash())
	// blockChain := NewBlockchain()
	// blockChain.Print()
	// blockChain.CreateBlock(5, "hash 1")
	// blockChain.Print()
	// blockChain.CreateBlock(2, "hash 2")
	// blockChain.Print()
}
