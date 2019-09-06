package main

type Block struct{
    Version int64
    TimeStamp int64
    Bits  int64
    Nonce int64
    MerkelRoot  []byte
    PrevBlockHash []byte
    Data []byte
}

func NewBlock(data string,prevblockhash []byte) *Block{
    var block Block
    block = Block{
        Version:1,
        TimeStamp:time.Now().Unix(),
        Bits:1,
        Nonce:1,
        MerkelRoot: []byte{},
        PrevBlockHash:prevblockhash,
        Data:[]byte(data)
    }
    block.SetHash()
    return *block
}
func (block *Block)SetHash(){
    tmp := [][]byte{
        IntToByte(block.Version),
        IntToByte(block.TimeStamp),
        IntToByte(block.Bits),
        IntToByte(block.Nonce),
        block.MerkelRoot,
        block.PrevBlockHash,
        block.Data}
    data := bytes.Join(tmp,[]byte)
    hash := sha256.SUM256(data)
    block.Hash = hash[:]
}
func GenesisBlock(){
    NewBlock("genesis is a block",[]byte{})
}
