/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: version.py
@Time: 2021-12-04 18:38
@Last_update: 2021-12-04 18:38
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

//func init() {
//	rootCmd.AddCommand(versionCmd)
//}
//
//var versionCmd = &cobra.Command{
//	Use:   "version",
//	Short: "Print the version number of Hugo",
//	Long:  `All software has versions. This is Hugo's`,
//	Run: func(cmd *cobra.Command, args []string) {
//		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
//	},
//}
func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Print the version number of Hugo",
	Long: `All software has version. This is Hugo's'`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Statisc Site Generator v0.9 -- HEAD")
	},
}