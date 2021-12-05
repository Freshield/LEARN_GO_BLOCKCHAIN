/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: server.py
@Time: 2021-12-07 15:18
@Last_update: 2021-12-07 15:18
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

var ServerCmd = &cobra.Command{
	Use: "server",
	Short: "A test server",
	Long: `Server is a test application for print things`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Server")
	},
}

