/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: root.py
@Time: 2021-12-04 19:25
@Last_update: 2021-12-04 19:25
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
	"os"
)

var rootCmd = &cobra.Command{
	Use: "demo",
	Short: "A cobra demo",
	Long: "A cobra demo, xxxx",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("execute root cmd")
	},
}

func Exectue() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}