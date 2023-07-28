package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log" //print manipulation
	"strings"
	"time"
)

// Block Struct: What a block contains
type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

func NewBlock(nonce int, previousHash string) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	//tx
	return b
}

// FOR EASE OF VIEW: this is the way we want our block function to be printed:
func (b *Block) PrintBlock() {
	fmt.Println("Timestamp:     \n", b.timestamp)
	fmt.Println("Nonce:         \n", b.nonce)
	fmt.Println("Previous Hash: \n", b.previousHash)
	fmt.Println("Transactions:  \n", b.transactions)
}

func (b *Block) Hash() [32]byte{
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}

// Blockchain Struct: what a blockchain contains
type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "Init Hash") //we are initializing very first(or 0th) block here. Rest of the blocks will be added after this.
	fmt.Println(bc)
	return bc
}

// '(bc *Blockchain)' means this function is defined as a method of the 'Blockchain' struct. '(bc *Blockchain)' is called the reciever of the function 'CreateBlock'. --> Go specific type
func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	// fmt.Println(bc.chain)
	bc.chain = append(bc.chain, b)
	// fmt.Println(bc.chain)

	return b
}

func init() {
	log.SetPrefix("Blockchain: ")
}

// EASE OF VIEW:
func (bc *Blockchain) PrintBlockchain() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain id: %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.PrintBlock()
	}
	//at the end of the log
	fmt.Printf("%s\n", strings.Repeat("*", 70))
}

func main() {
	// blockchain := NewBlockchain()
	// blockchain.CreateBlock(1, "hash 1")
	// blockchain.CreateBlock(2, "hash 2")
	// blockchain.CreateBlock(3, "hash 3")
	// blockchain.PrintBlockchain()
	//   NewBlockchain().CreateBlock(12,"abc")
}

