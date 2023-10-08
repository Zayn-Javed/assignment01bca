/*
	     	Name: Zain Javed
			Student-id: 20I-0522
			Assignment # 1
*/

package main

import (
	"fmt"

	"a1/myPackage"
)

func main() {
	// Initialize the blockchain with a genesis block.
	genesisBlock := myPackage.NewBlock("Alice to Bob", 0, "")
	Blockchain := []*myPackage.Block{genesisBlock}

	// Add more blocks to the blockchain.
	Blockchain = append(Blockchain, myPackage.NewBlock("Alexa to Micheal", 123, Blockchain[len(Blockchain)-1].Hash))
	Blockchain = append(Blockchain, myPackage.NewBlock("John to Mellus", 456, Blockchain[len(Blockchain)-1].Hash))

	// Display all blocks in the blockchain.
	fmt.Print("\nBlockchain:\n")
	myPackage.DisplayBlocks(Blockchain)

	// Change the transaction of the second block.
	concernedBlock := Blockchain[1]
	changedTransaction := "Modified Transaction String"
	myPackage.ChangeBlock(concernedBlock, changedTransaction)
	fmt.Print("Block on index 1 has been changed!\n\n")

	// Print blockchain after making changes in block present on index 1 of blockchain
	fmt.Print("Blockchain after changing block:\n")
	myPackage.DisplayBlocks(Blockchain)

	// Verify the integrity of the blockchain.
	isTrue := myPackage.VerifyChain(Blockchain)
	if isTrue {
		fmt.Println("Blockchain is valid.")
	} else {
		fmt.Println("Blockchain is invalid.")
	}
}
