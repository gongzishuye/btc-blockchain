package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type CLI struct{}

func (cli *CLI) checkSanity() {
	if len(os.Args) < 2 {
		fmt.Println(os.Args)
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  createwallet - Generates a new key-pair and saves it into the wallet file")
	fmt.Println("  startnode -miner ADDRESS - Start a node with ID specified in NODE_ID env. var. -miner enables mining")
}

func (cli *CLI) createWallet() {
	address := newWallet()
	fmt.Printf("New wallet address: %s", address)
}

func (cli *CLI) runBlockChain() {
	bc := BlockChain{[]*Block{}}
	bc.NewGensisBlock()
	for {
		block := &Block{time.Now().Unix(), []byte("HHH"), bc.blocks[len(bc.blocks)-1].Hash, []byte{}, 0}
		pow := NewProofOfWork(block)
		nonce, hash := pow.Run()
		block.Nonce = int64(nonce)
		block.Hash = hash

		bc.blocks = append(bc.blocks, block)
		fmt.Println(block)
	}
}

func (cli *CLI) run() {
	cli.checkSanity()

	createWalletCmd := flag.NewFlagSet("createwallet", flag.ExitOnError)
	runBlockChainCmd := flag.NewFlagSet("runblockchain", flag.ExitOnError)

	err := createWalletCmd.Parse(os.Args[2:])
	if err != nil {
		log.Panic(err)
	}
	err = runBlockChainCmd.Parse(os.Args[2:])
	if err != nil {
		log.Panic(err)
	}

	if createWalletCmd.Parsed() {
		cli.createWallet()
	} else if runBlockChainCmd.Parsed() {
		cli.runBlockChain()
	}

}
