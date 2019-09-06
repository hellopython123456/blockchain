package main
import "fmt"
func main() {
    bc := NewBlockChain()
    bc.AddBlock("A send B 1 BTC")
    bc.AddBlock("B send C 1 BTC")
    for _,block := range bc.blocks{
        fmt.Printf("Version: %d\n",block.Version)
        fmt.Printf("PrevBlockHash: %x\n",block.PrevBlockHash)
        fmt.Printf("Hash: %x\n", block.Hash)
        fmt.Printf("TimeStamp: %d\n", block.TimeStamp)
        fmt.Printf("Bits: %d\n", block.Bits)
        fmt.Printf("PrevBlockHash: %d\n", block.prevBlockHash)
    }
}
