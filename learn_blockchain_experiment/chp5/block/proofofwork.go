/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: proofofwork.py
@Time: 2022-01-06 16:48
@Last_update: 2022-01-06 16:48
@Desc: None
@==============================================@
@      _____             _   _     _   _       @
@     |   __|___ ___ ___| |_|_|___| |_| |      @
@     |   __|  _| -_|_ -|   | | -_| | . |      @
@     |__|  |_| |___|___|_|_|_|___|_|___|      @
@                                    Freshield @
@==============================================@
*/
package block

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

var (
	// 最大尝试次数
	maxNonce = math.MaxInt64
)

// 挖矿难度
const targetBits = 24

type ProofOfWork struct {
	// pow的基本结构
	block *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	// 新建一个pow
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	// 准备数据，也就是把相应的数据转为字节
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.HashTransactions(),
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{})

	return data[:]
}

func (pow *ProofOfWork) Run() (int, []byte) {
	// 运行pow，也就是找到符合的nonce
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining a new block")
	for nonce < maxNonce {
		data := pow.prepareData(nonce)

		hash = sha256.Sum256(data)
		fmt.Printf("\r%64x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce ++
		}
	}
	fmt.Print("\n\n")

	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	// 判断pow的结果是否正确
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}