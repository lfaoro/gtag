package main

import (
	"fmt"
	"os/exec"

	"github.com/urfave/cli"
)

var zeroCmd = cli.Command{
	Name:    "zero",
	Aliases: []string{"z", "ze", "zer"},
	Usage:   "create zero tag v0.0.0",
	Flags:   []cli.Flag{},
	Action: func(c *cli.Context) error {
		commit, err := exec.Command("sh", "-c",
			"git rev-parse HEAD").CombinedOutput()
		if err != nil {
			return err
		}
		cmd2 := exec.Command("sh", "-c",
			"git tag -a 0.0.0 -m \"the zero tag\"")
		outErr, err := cmd2.CombinedOutput()
		if err != nil {
			fmt.Println(string(outErr))
			return err
		}
		fmt.Printf("tag 0.0.0 created on commit %v\n", string(commit[:8]))
		return nil
	},
}
