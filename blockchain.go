package main

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBloc := bc.blocks[len(bc.blocks)-1]
	newBloc := NewBlock(data, prevBloc.Hash)
	bc.blocks = append(bc.blocks, newBloc)
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
