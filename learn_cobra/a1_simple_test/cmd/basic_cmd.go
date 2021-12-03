/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: basic_cmd.py
@Time: 2021-12-04 18:07
@Last_update: 2021-12-04 18:07
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

var Verbose bool
var Source string

var rootCmd = &cobra.Command{
	Use: "hugo",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Here is HUGO", args)
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(
		&Verbose, "verbose", "v", false, "verbose output")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}