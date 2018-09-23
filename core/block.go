package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index        int64
	Timestamp    int64
	PreBlockHash string
	Hash         string
	Data         string
	Valid        bool
	//Nonce        int
}

func (b *Block) getHash() {
	data := string(b.Index) + string(b.Timestamp) + string(b.PreBlockHash) + string(b.Data)
	hashByte := sha256.Sum256([]byte(data))
	b.Hash = hex.EncodeToString(hashByte[:])
}

func (b *Block) new(pre *Block) {
	b.Index = pre.Index + 1
	b.Timestamp = time.Now().Unix()
	b.PreBlockHash = pre.Hash
	b.getHash()
	//pow := POW{}
	//pow.New(b)
	//b.nonce, b.hash = pow.RUN()
}

func (b *Block) validate(bc *BlockChain) {
	oriHash := b.Hash
	b.getHash()
	trueHash := b.Hash
	if oriHash == trueHash &&
		b.PreBlockHash == bc.Blocks[len(bc.Blocks)-1].Hash &&
		b.Index == int64(len(bc.Blocks)) {
		b.Valid = true
	} else {
		b.Valid = false
	}
	b.Hash = oriHash
}
