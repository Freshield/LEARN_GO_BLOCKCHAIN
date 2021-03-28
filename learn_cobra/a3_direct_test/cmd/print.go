/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: print.py
@Time: 2021-12-04 19:28
@Last_update: 2021-12-04 19:28
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

func init() {
	rootCmd.AddCommand(printCmd)

	printCmd.Flags().StringVarP(
		&Msg, "message", "m", "default message", "message to be printed")
}

var Msg string
var printCmd = &cobra.Command{
	Use: "print",
	Short: "print something",
	Long: "print something, xxxx",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("execute print cmd", Msg)
	},
}