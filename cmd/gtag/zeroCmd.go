// Copyright (c) 2019 Leonardo Faoro. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"errors"
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
			if debugFlag {
				fmt.Println(err)
			}
			return errors.New("no commits in this repo to apply a tag")
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
