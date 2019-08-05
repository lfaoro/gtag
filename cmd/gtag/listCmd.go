package main

import (
	"fmt"
	"os/exec"

	"github.com/urfave/cli"
)

var listCmd = cli.Command{
	Name:    "list",
	Aliases: []string{"l", "li", "lis"},
	Usage:   "lists all tags",
	Flags:   []cli.Flag{},
	Action: func(c *cli.Context) error {
		cmd := exec.Command("sh", "-c",
			"git --no-pager tag -l")
		outErr, err := cmd.CombinedOutput()
		exitIfError(err)
		fmt.Println(string(outErr))
		return nil
	},
}
