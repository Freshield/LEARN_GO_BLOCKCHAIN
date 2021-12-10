/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: main.py
@Time: 2021-12-10 16:39
@Last_update: 2021-12-10 16:39
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
	"chp3/block"
	"fmt"
	"strconv"
)

func main() {
	bc := block.NewBlockchain()
	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, b := range bc.Blocks {
		fmt.Printf("Prev. hash:%64x\n", b.PrevBlockHash)
		fmt.Printf("Data:      %s\n", b.Data)
		fmt.Printf("Hash:      %64x\n", b.Hash)

		pow := block.NewProofOfWork(b)
		fmt.Printf("POW: %s\n", strconv.FormatBool(pow.Validatae()))
		fmt.Println()
	}
}