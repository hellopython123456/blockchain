package main

type BlockChain struct{
    blocks []*Block
}
func NewBlockChain() *BlockChain{
    block := GenesisBlock()
    return &BlockChain{blocks:[]*Block{block}}
}
func (bc *BlockChain)AddBlock(data string){
    prevblockhash := bc.blocks[len(bc.blocks)-1].Hash
    block := NewBlock(data string,prevblockhash)
    bc.blocks = append(bc.blocks,block)
}
