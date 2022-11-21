package main

type BlockChain struct {
	blocks []*Block
}

func (bc *BlockChain) NewGensisBlock() {
	gensis := NewBlock("Here We Go BTC", []byte{}, 0, []byte{})
	bc.blocks = append(bc.blocks, gensis)
}
