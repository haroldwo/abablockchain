package main

import (
	"fmt"

	"github.com/haroldwo/abablockchain/core"
)

func main() {
	blockchain := &core.BlockChain{}
	blockchain.New()

	blockchain.AddBlock("Data1")
	blockchain.AddBlock("Data2")

	for _, block := range blockchain.Blocks {
		fmt.Println(block)
		//pow := core.POW{}
		//pow.New(block)
		//fmt.Println(strconv.FormatBool(pow.Validate()))
	}

	blockchain.HttpServer()
}
