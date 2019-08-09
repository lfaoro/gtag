package main

import (
	"fmt"
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

	lastTag, err := shellCmd("git describe --abbrev=0")
	if err != nil {
		return err
	}

	tag := strings.TrimSpace(lastTag)
	tag = strings.TrimPrefix(tag, "v")
	fmt.Println(tag)

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
		message = fmt.Sprintf("incremented %v -> %v on commit %v\n", vOld, vNewS, commit[:8])
	}
	fmt.Printf(message)

	cmd := fmt.Sprintf("git tag -a %v -m \"%v\"", vNewS, message)
	newTag, err := shellCmd(cmd)
	if err != nil {
		fmt.Println(newTag)
		return err
	}

	if c.Bool("push") {
		cmd := fmt.Sprintf("git push origin %v", vNewS)
		push, err := shellCmd(cmd)
		if err != nil {
			fmt.Println(push)
			return err
		}
		fmt.Println(push)
	}

	return nil
}
