package main

import (
	"fmt"
	"log" //print manipulation
	"time"
  "strings"
)

//Block Struct: What a block contains
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
	return b
}

// FOR EASE OF VIEW: this is the way we want our block function to be printed:
func (b *Block) PrintBlock() {
	fmt.Println("Timestamp:     \n", b.timestamp)
	fmt.Println("Nonce:         \n", b.nonce)
	fmt.Println("Previous Hash: \n", b.previousHash)
	fmt.Println("Transactions:  \n", b.transactions)
}

//Blockchain Struct: what a blockchain contains
type Blockchain struct{
  transactionPool []string
  chain []*Block
}

func NewBlockchain() *Blockchain{
  bc:= new(Blockchain)
  bc.CreateBlock(0, "Init Hash")
  return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block{
  b := NewBlock(nonce, previousHash)
  bc.chain = append(bc.chain,b)
  return b
}


func init() {
	log.SetPrefix("Blockchain: ")
}

//EASE OF VIEW: 
func (bc *Blockchain) PrintBlockchain(){
  for i, block := range bc.chain{
    fmt.Printf("%s Chain id: %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
    block.PrintBlock()
  }
  //at the end of the log
  fmt.Printf("%s\n",strings.Repeat("*",70))
}
func main() {
	// b := NewBlock(123, "123abc")
	// b.Print()

  blockchain := NewBlockchain()
  // blockchain.Print()
  blockchain.CreateBlock(5, "hash 1")
  // blockchain.Print()
  blockchain.CreateBlock(2, "hash 2")
  blockchain.PrintBlockchain()
}
