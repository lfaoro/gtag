package main

import (
	"fmt"

	"github.com/urfave/cli"
)

var zeroCmd = cli.Command{
	Name:    "zero",
	Aliases: []string{"z"},
	Usage:   "create zero tag v0.0.0",
	Flags:   []cli.Flag{},
	Action: func(c *cli.Context) error {
		commit, err := shellCmd("git rev-parse HEAD")
		if err != nil {
			return err
		}

		outErr, err := shellCmd("git tag -a v0.0.0 -m \"the zero tag\"")
		if err != nil {
			fmt.Println(outErr)
			return err
		}
		fmt.Printf("tag v0.0.0 created on commit %v\n", commit[:8])
		return nil
	},
}
