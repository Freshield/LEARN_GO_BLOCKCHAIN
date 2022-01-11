/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: block.py
@Time: 2022-01-06 16:34
@Last_update: 2022-01-06 16:34
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
	"encoding/gob"
	"log"
	"time"
)

type Block struct {
	// 区块的基本结构
	Timestamp int64
	Transactions []*Transaction
	PrevBlockHash []byte
	Hash []byte
	Nonce int
}

func (b *Block) Serialize() []byte {
	// 序列化一个区块
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func (b *Block) HashTransactions() []byte {
	// 把一个交易hash化 转为byte
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
}

func NewBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
	// 新建一个区块
	block := &Block{
		time.Now().Unix(), transactions, prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func NewGenesisBlock(coinbase *Transaction) *Block {
	// 新的创世区块
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

func DeserializeBlock(d []byte) *Block {
	// 反序列化区块
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}