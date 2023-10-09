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

// declare the structure of the block present in blockchain
type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	Hash         string
}

// calculate the has of hte string which contain trasaction, nonce and previous hash
func CalculateHash(stringToHash string) string {
	hashGenerattor := sha256.New()
	hashGenerattor.Write([]byte(stringToHash))
	hashInBytes := hashGenerattor.Sum(nil)
	hashInStr := fmt.Sprintf("%x", hashInBytes)
	return hashInStr
}

// create a new block and returns a pointer of this block
func NewBlock(transaction string, nonce int, previousHash string) *Block {
	// create block
	block := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
	}
	// format the string to generate its hash
	formattedString := fmt.Sprintf("%s%d%s", block.Transaction, block.Nonce, block.PreviousHash)
	// calculate the hash of the block and assign hash to the attribute of block structure
	block.Hash = CalculateHash(formattedString)
	return block
}

// prints all blocks in the blockchain in well-structures format
func DisplayBlocks(Blockchain []*Block) {
	// iterate all blocks in a blockchain
	for i := 0; i < len(Blockchain); i++ {
		block := Blockchain[i]
		// print blocks in well structured format
		fmt.Printf("Transaction: %s\nNonce: %d\nPrevious Hash: %s\nCurrent Hash: %s\n\n", block.Transaction, block.Nonce, block.PreviousHash, block.Hash)
	}
}

// tamper a given block through the modified transection string
func ChangeBlock(block *Block, changedTransaction string) {
	// change the string of the block's trasaction
	block.Transaction = changedTransaction
	// format the string to generate its hash
	formattedString := fmt.Sprintf("%s%d%s", block.Transaction, block.Nonce, block.PreviousHash)
	// calculate the hash of the block and assign hash to the attribute of block structure
	block.Hash = CalculateHash(formattedString)
}

// verifie the the blockchain by checking the hash of block and validating its hash with the hash of previous block's hash
func VerifyChain(Blockchain []*Block) bool {
	// iterate all blocks in a blockchain
	for i := 1; i < len(Blockchain); i++ {
		// keep reference of the current block
		currentBlock := Blockchain[i]
		// keep reference of the previous block
		previousBlock := Blockchain[i-1]
		// format the string to generate its hash
		formattedString := fmt.Sprintf("%s%d%s", currentBlock.Transaction, currentBlock.Nonce, currentBlock.PreviousHash)
		currentHash := CalculateHash(formattedString)
		// compare the calculted hash with the previously stored hash in a block
		if currentBlock.Hash != currentHash {
			return false
		} else {
			// compare the hash of current block with the hash of previous block in a blockchain
			if currentBlock.PreviousHash != previousBlock.Hash {
				return false
			}
		}
	}
	return true
}
