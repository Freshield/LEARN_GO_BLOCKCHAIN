/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: clone.py
@Time: 2021-12-04 21:20
@Last_update: 2021-12-04 21:20
@Desc: None
@==============================================@
@      _____             _   _     _   _       @
@     |   __|___ ___ ___| |_|_|___| |_| |      @
@     |   __|  _| -_|_ -|   | | -_| | . |      @
@     |__|  |_| |___|___|_|_|_|___|_|___|      @
@                                    Freshield @
@==============================================@
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	Use: "clone url [destination]",
	Short: "Clone a repository into a new directory",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("clone", args)
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)
}