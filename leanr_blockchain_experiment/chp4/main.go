/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: main.py
@Time: 2021-12-21 16:09
@Last_update: 2021-12-21 16:09
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

import "chp4/block"

func main() {
	bc := block.NewBlockchain()
	defer bc.DB.Close()

	cli := block.CLI{bc}
	cli.Run()
}