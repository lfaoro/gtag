// Copyright (c) 2019 Leonardo Faoro. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/lfaoro/semver"
	"github.com/urfave/cli"
)

func semCmd(cmd string) cli.Command {
	return cli.Command{
		Name:    cmd,
		Aliases: []string{cmd[0:2]},
		Usage:   fmt.Sprintf("increases %v semantic version", cmd),
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "message, m",
				Usage: "adds a message to the new tag",
			},
			cli.BoolFlag{
				Name:   "push, p",
				Usage:  "pushes tags upstream",
				EnvVar: "GTAG_PUSH",
			},
		},
		Action: func(c *cli.Context) error {
			return SemVersion(c, cmd)
		},
	}
}

func SemVersion(c *cli.Context, field string) error {
	commit, err := shellCmd("git rev-parse HEAD")
	if err != nil {
		return err
	}
	// shorten commit string without causing a possible panic.
	if len(commit) >= 8 {
		commit = commit[:8]
	}

	lastTag, err := shellCmd("git describe --abbrev=0")
	if err != nil {
		return err
	}

	tag := strings.TrimSpace(lastTag)
	tag = strings.TrimPrefix(tag, "v")

	vNew, err := semver.Make(tag)
	if err != nil {
		return err
	}
	vOld := "v" + vNew.String()

	switch field {
	case "patch":
		err = vNew.IncrementPatch()
	case "minor":
		err = vNew.IncrementMinor()
	case "major":
		err = vNew.IncrementMajor()
	}
	if err != nil {
		return err
	}

	vNewS := "v" + vNew.String()

	var message string
	if msg := c.String("message"); msg != "" {
		message = msg
	} else {
		message = fmt.Sprintf("incremented tag %v -> %v on commit %v\n", vOld, vNewS, commit)
	}
	fmt.Printf(message)

	cmd := fmt.Sprintf("git tag -a %v -m \"%v\"", vNewS, message)
	newTag, err := shellCmd(cmd)
	if err != nil {
		fmt.Println(newTag)
		return err
	}

	yesPush := c.Bool("yesPush")
	if yesPush {
		err := pushTag(vNewS)
		if err != nil {
			return err
		}
	} else {
		// ask the user
		input := bufio.NewReader(os.Stdin)
		fmt.Print("Would you like to push this tag to the origin? ")

		b, err := input.ReadString('\n')
		exitIfError(err)

		// return if the user doesn't want to push
		if strings.Contains(b, "n") {
			fmt.Printf("tag %s not pushed to the origin\n", tag)
			return nil
		}

		// push to the origin
		err = pushTag(vNewS)
		if err != nil {
			return err
		}
	}

	return nil
}

func pushTag(tag string) error {
	cmd := fmt.Sprintf("git push origin %v", tag)
	push, err := shellCmd(cmd)
	if err != nil {
		fmt.Println(push)
		return err
	}
	fmt.Println(push)
	return nil
}
