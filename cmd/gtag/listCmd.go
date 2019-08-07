package main

import (
	"fmt"
	"os/exec"

	"github.com/urfave/cli"
)

var listCmd = cli.Command{
	Name:    "list",
	Aliases: []string{"l", "lis"},
	Usage:   "list all git tags",
	Flags:   []cli.Flag{},
	Action: func(c *cli.Context) error {
		out, err := exec.Command("sh", "-c",
			"git --no-pager tag -l").CombinedOutput()
		exitIfError(err)
		fmt.Println(string(out))
		return nil
	},
}
