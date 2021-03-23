/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: helpers.py
@Time: 2021-12-04 20:58
@Last_update: 2021-12-04 20:58
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
	"os"
	"os/exec"
	"github.com/spf13/cobra"
)

func ExecCommand(name string, subname string, args ...string) (string, error) {
	args = append([]string{subname}, args...)
	cmd := exec.Command(name, args...)
	bytes, err := cmd.CombinedOutput()

	return string(bytes), err
}

func Error(command *cobra.Command, args []string, err error) {
	fmt.Fprintf(
		os.Stderr, "execute %s args: %v error: %v\n", command, args, err)
	os.Exit(1)
}