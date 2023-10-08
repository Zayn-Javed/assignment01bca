/*
	     	Name: Zain Javed
			Student-id: 20I-0522
			Assignment # 1
*/

package assignment01bca

import (
	"crypto/sha256"
	"fmt"
)

// Block represents an individual block in the blockchain.
type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	Hash         string
}

// CalculateHash calculates the SHA-256 hash of a string.
func CalculateHash(stringToHash string) string {
	hasher := sha256.New()
	hasher.Write([]byte(stringToHash))
	hashBytes := hasher.Sum(nil)
	return fmt.Sprintf("%x", hashBytes)
}

// NewBlock creates a new block and returns a pointer to it.
func NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
	}
	formattedString := fmt.Sprintf("%s%d%s", block.Transaction, block.Nonce, block.PreviousHash)
	block.Hash = CalculateHash(formattedString)
	return block
}

// DisplayBlocks prints all blocks in the blockchain.
func DisplayBlocks(Blockchain []*Block) {
	for i := 0; i < len(Blockchain); i++ {
		block := Blockchain[i]
		fmt.Printf("Transaction: %s\nNonce: %d\nPrevious Hash: %s\nCurrent Hash: %s\n\n",
			block.Transaction, block.Nonce, block.PreviousHash, block.Hash)
	}
}

// ChangeBlock updates the transaction of a given block.
func ChangeBlock(block *Block, changedTransaction string) {
	block.Transaction = changedTransaction
	formattedString := fmt.Sprintf("%s%d%s", block.Transaction, block.Nonce, block.PreviousHash)
	block.Hash = CalculateHash(formattedString)
}

// VerifyChain verifies the integrity of the blockchain.
func VerifyChain(Blockchain []*Block) bool {
	for i := 1; i < len(Blockchain); i++ {
		currentBlock := Blockchain[i]
		previousBlock := Blockchain[i-1]
		formattedString := fmt.Sprintf("%s%d%s", currentBlock.Transaction, currentBlock.Nonce, currentBlock.PreviousHash)
		currentHash := CalculateHash(formattedString)
		if currentBlock.Hash != currentHash {
			return false
		} else {
			if currentBlock.PreviousHash != previousBlock.Hash {
				return false
			}
		}
	}
	return true
}
