package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/blang/semver"
	"github.com/urfave/cli"
)

func semCmd(cmd string) cli.Command {
	return cli.Command{
		Name:    cmd,
		Aliases: []string{cmd[0:0], cmd[0:1], cmd[0:2]},
		Usage:   fmt.Sprintf("increases %v semantic version", cmd),
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "message, m",
				Usage: "adds message to the new tag",
			},
			cli.BoolFlag{
				Name:  "push, p",
				Usage: "push tags to origin",
				EnvVar: "GTAG_PUSH",
			},
		},
		Action: func(c *cli.Context) error {
			return SemVersion(c, cmd)
		},
	}
}

func SemVersion(c *cli.Context, field string) error {
	commit, err := exec.Command("sh", "-c",
		"git rev-parse HEAD").CombinedOutput()
	if err != nil {
		return err
	}

	lastTag, err := exec.Command("sh", "-c",
		"git describe --abbrev=0").CombinedOutput()
	if err != nil {
		return err
	}
	tag := strings.TrimSuffix(string(lastTag), "\n")
	v, err := semver.Make(tag)
	if err != nil {
		return err
	}
	v0 := v

	switch field {
	case "patch":
		err = v.IncrementPatch()
	case "minor":
		err = v.IncrementMinor()
	case "major":
		err = v.IncrementMajor()
	}
	if err != nil {
		return err
	}

	var message string
	if msg := c.String("message"); msg != "" {
		message = msg
	} else {
		message = fmt.Sprintf("incremented %v -> %v on commit %v\n", v0, v, string(commit[:8]))
	}
	fmt.Printf(message)
	cmd := fmt.Sprintf("git tag -a %v -m \"%v\"", v, message)
	newTag, err := exec.Command("sh", "-c",
		cmd).CombinedOutput()
	if err != nil {
		fmt.Println(string(newTag))
		return err
	}

	if c.Bool("push") {
		cmd := fmt.Sprintf("git push origin %v", v)
		push, err := exec.Command("sh", "-c",
			cmd).CombinedOutput()
		if err != nil {
			fmt.Println(push)
			return err
		}
		fmt.Println(push)
	}

	return nil
}
