/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: version.py
@Time: 2021-12-04 21:09
@Last_update: 2021-12-04 21:09
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

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "version subcommand show git version info.",

	Run: func(cmd *cobra.Command, args []string) {
		output, err := ExecCommand("git", "version", args...)
		if err != nil {
			Error(cmd, args, err)
		}

		fmt.Fprint(os.Stdout, output)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}