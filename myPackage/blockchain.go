/*
	     	Name: Zain Javed
			Student-id: 20I-0522
			Assignment # 1
*/

package myPackage

import (
	"crypto/sha256"
	"fmt"
)

// Block represents an individual block in the blockchain.
type Block struct {
	transaction  string
	nonce        int
	previousHash string
	hash         string
}

// CalculateHash calculates the SHA-256 hash of a string.
func CalculateHash(stringToHash string) string {
	hasher := sha256.New()
	hasher.Write([]byte(stringToHash))
	hashBytes := hasher.Sum(nil)
	return fmt.Sprintf("%x", hashBytes)
}

// NewBlock creates a new block and returns a pointer to it.
func NewBlock(Transaction string, Nonce int, PreviousHash string) *Block {
	block := &Block{
		transaction:  Transaction,
		nonce:        Nonce,
		previousHash: PreviousHash,
	}
	formattedString := fmt.Sprintf("%s%d%s", block.transaction, block.nonce, block.previousHash)
	block.hash = CalculateHash(formattedString)
	return block
}

// DisplayBlocks prints all blocks in the blockchain.
func DisplayBlocks(Blockchain []*Block) {
	for i := 0; i < len(Blockchain); i++ {
		block := Blockchain[i]
		fmt.Printf("Transaction: %s\nNonce: %d\nPrevious Hash: %s\nCurrent Hash: %s\n\n",
			block.transaction, block.nonce, block.previousHash, block.hash)
	}
}

// ChangeBlock updates the transaction of a given block.
func ChangeBlock(block *Block, changedTransaction string) {
	block.transaction = changedTransaction
	formattedString := fmt.Sprintf("%s%d%s", block.transaction, block.nonce, block.previousHash)
	block.hash = CalculateHash(formattedString)
}

// VerifyChain verifies the integrity of the blockchain.
func VerifyChain(Blockchain []*Block) bool {
	for i := 1; i < len(Blockchain); i++ {
		currentBlock := Blockchain[i]
		previousBlock := Blockchain[i-1]
		formattedString := fmt.Sprintf("%s%d%s", currentBlock.transaction, currentBlock.nonce, currentBlock.previousHash)
		currentHash := CalculateHash(formattedString)
		if currentBlock.hash != currentHash {
			return false
		} else {
			if currentBlock.previousHash != previousBlock.hash {
				return false
			}
		}
	}
	return true
}
