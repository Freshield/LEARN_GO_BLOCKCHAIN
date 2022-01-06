/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: main.py
@Time: 2022-01-13 16:53
@Last_update: 2022-01-13 16:53
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
	"chp5/block"
	"fmt"
)

func main() {
	//cli := block.CLI{}
	//cli.Run()
	wallet := block.NewWallet()
	address := wallet.GetAddress()
	fmt.Println(address)
	fmt.Println(string(address))
	fmt.Println(wallet.PublicKey)
	fmt.Println(block.Base58Decode(address))
}