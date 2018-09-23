package core

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type BlockChain struct {
	Blocks []*Block
}

func (bc *BlockChain) appendBlock(newblock Block) {
	newblock.validate(bc)
	if newblock.Valid == true {
		fmt.Println(newblock)
		bc.Blocks = append(bc.Blocks, &newblock)
	} else {
		log.Fatal("Invalid block.")
	}
}

func (bc *BlockChain) AddBlock(data string) {
	newblock := Block{}
	newblock.Data = data
	newblock.new(bc.Blocks[len(bc.Blocks)-1])
	bc.appendBlock(newblock)
}

func (bc *BlockChain) HttpServer() {
	http.HandleFunc("/blockchain/get", bc.getBlockChain)
	http.HandleFunc("/blockchain/put", bc.putBlockChain)
	http.ListenAndServe("localhost:8080", nil)
}

func (bc *BlockChain) getBlockChain(w http.ResponseWriter, r *http.Request) {
	fmt.Println(bc.Blocks)
	data, err := json.Marshal(bc.Blocks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(data))
}

func (bc *BlockChain) putBlockChain(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query().Get("data")
	bc.AddBlock(data)
	bc.getBlockChain(w, r)
}

func (bc *BlockChain) New() {
	genesis := Block{}
	genesis.Timestamp = time.Now().Unix()
	genesis.Data = "Genesis Block."
	bc.Blocks = []*Block{&genesis}
}
