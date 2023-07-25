package main

import (
	"fmt"
	"log" //shows Date()
	"time"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

type Block struct{
    nonce int;
    previousHash string;
    timestamp int64;
    transactions []string;
}

func NewBlock(nonce int, previousHash string) *Block{
    b := new(Block)
    b.timestamp = time.Now().UnixNano()
    b.nonce = nonce 
    b.previousHash = previousHash
    return b
}

//this is the way we want our block function to be printed:
func(b *Block) Print(){
    fmt.Println("Timestamp:     \n", b.timestamp)
    fmt.Println("Nonce:         \n", b.nonce)
    fmt.Println("Previous Hash: \n", b.previousHash)
    fmt.Println("Transactions:  \n", b.transactions)
}

func main() {
	b := NewBlock(123, "123abc")

    // fmt.Print(b)
    b.Print()
    
	log.Println("test log")
}
