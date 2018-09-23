package core

import (
	"math"
	"math/big"
)

var (
	maxNonce = math.MaxInt64
)

const targetBits = 20

type POW struct {
	block  *Block
	target *big.Int
}

func (pow *POW) New(b *Block) {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	pow = &POW{b, target}
}

func (pow *POW) data() {

}
