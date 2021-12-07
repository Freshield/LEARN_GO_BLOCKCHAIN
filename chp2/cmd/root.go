/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: root.py
@Time: 2021-12-01 22:03
@Last_update: 2021-12-01 22:03
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
	"chp2/imp"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	name string
	age int
)

func init() {
	RootCmd.Flags().StringVarP(&name, "name", "n", "", "person's name")
	RootCmd.Flags().IntVarP(&age, "age", "a", 0, "person's age")
	RootCmd.AddCommand(ServerCmd)
}

var RootCmd = &cobra.Command{
	Use: "demo",
	Short: "A test demo",
	Long: `Demo is a test application for print things`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(name) == 0 {
			cmd.Help()
			return
		}
		imp.Show(name, age)
	},
}

func Execute()  {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}