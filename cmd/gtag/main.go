// Copyright (c) 2019 Leonardo Faoro. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var (
	// vars are injected during the release process.
	// ldflags are set by Goreleaser
	//
	// ref: https://goreleaser.com/environment/
	version = "latest-dev"
	commit  = "none"
	date    = "unknown"
)

var debugFlag bool

func main() {
	app := cli.NewApp()
	app.Name = "gTag"
	app.Usage = "gTag is a $(git tag) workflow tool"
	app.Version = fmt.Sprintf("%s \ncommit %s \nbuilt on %s", version, commit, date)
	app.EnableBashCompletion = true
	app.Authors = []cli.Author{
		{
			Name:  "Leonardo Faoro",
			Email: "lfaoro@gmail.com",
		},
	}
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug, d",
			Usage:       "displays debug information.",
			EnvVar:      "GTAG_DEBUG",
			Destination: &debugFlag,
			Hidden:      true,
		},
	}
	app.Action = func(c *cli.Context) error {
		cmd := c.App.Command("list")
		err := cmd.Run(c)
		if err != nil {
			return err
		}
		return nil
	}

	app.Before = func(c *cli.Context) error {
		exitIfError(isGitRepo())
		return nil
	}

	app.Commands = []cli.Command{
		listCmd,
		zeroCmd,
		semCmd("patch"),
		semCmd("minor"),
		semCmd("major"),
		deleteCmd,
	}

	err := app.Run(os.Args)
	exitIfError(err)
}

func exitIfError(err error) {
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		os.Exit(1)
	}
}
