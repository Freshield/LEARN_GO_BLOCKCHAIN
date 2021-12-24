/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: main.py
@Time: 2021-12-09 16:04
@Last_update: 2021-12-09 16:04
@Desc: None
@==============================================@
@      _____             _   _     _   _       @
@     |   __|___ ___ ___| |_|_|___| |_| |      @
@     |   __|  _| -_|_ -|   | | -_| | . |      @
@     |__|  |_| |___|___|_|_|_|___|_|___|      @
@                                    Freshield @
@==============================================@
*/
package main

import (
	"chp2/block"
	"crypto/sha256"
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	data1 := []byte("I like donuts")
	data2 := []byte("I like donutsca07ca")
	targetBits := 24
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	fmt.Printf("%x\n", sha256.Sum256(data1))
	fmt.Printf("%64x\n", target)
	fmt.Printf("%x\n", sha256.Sum256(data2))

	bc := block.NewBlockchain()
	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, b := range bc.Blocks {
		fmt.Printf("Prev. hash:%64x\n", b.PrevBlockHash)
		fmt.Printf("Data:      %s\n", b.Data)
		fmt.Printf("Hash:      %64x\n", b.Hash)

		pow := block.NewProofOfWork(b)
		fmt.Printf("POW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}