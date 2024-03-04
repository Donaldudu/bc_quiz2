package main

import (
    "fmt"
)

type Block struct {
    Data string
}

type Blockchain struct {
    Blocks []*Block
}


func (bc *Blockchain) DisplayAllBlocks() {
    fmt.Println("Blocks in the blockchain:")
    for _, block := range bc.Blocks {
        fmt.Println(block.Data)
    }
}


func (bc *Blockchain) NewBlock(data string) {
    newBlock := &Block{Data: data}
    bc.Blocks = append(bc.Blocks, newBlock)
    fmt.Println("New block added to the blockchain with data:", data)
}

func (bc *Blockchain) ModifyBlock(index int, newData string) {
    if index >= 0 && index < len(bc.Blocks) {
        bc.Blocks[index].Data = newData
        fmt.Println("Block at index", index, "modified with new data:", newData)
    } else {
        fmt.Println("Invalid index")
    }
}
