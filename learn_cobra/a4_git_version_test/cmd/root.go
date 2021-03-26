/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: root.py
@Time: 2021-12-04 21:07
@Last_update: 2021-12-04 21:07
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
	"errors"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "git",
	Short: "Git is a distributed version control system.",
	Long: `Git is a free and open source distributed version control system
designed to handle everything from small to very large projects 
with speed and efficiency.`,
	Run: func(cmd *cobra.Command, args []string) {
		Error(cmd, args, errors.New("unrecognized command"))
	},
}

func Execute() {
	rootCmd.Execute()
}