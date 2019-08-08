package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/urfave/cli"
)

var deleteCmd = cli.Command{
	Name:    "delete",
	Aliases: []string{"d", "del"},
	Usage:   "deletes the last created tag, pass --all to delete all tags at once.",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "all",
			Usage: "removes all tags",
		},
	},
	Action: func(c *cli.Context) error {
		out, err := shellCmd("git --no-pager tag -l")
		if err != nil {
			return err
		}
		s := strings.TrimSpace(string(out))
		tags := strings.Split(s, "\n")

		if c.Bool("all") {
			input := bufio.NewReader(os.Stdin)
			fmt.Print("You're about to delete all tags, sure? ")
			_, err := input.ReadString('\n')
			exitIfError(err)
			for _, tag := range tags {
				err := deleteTag(tag)
				if err != nil {
					return err
				}

				fmt.Printf("deleted tag %v\n", tag)
			}
			return nil
		}

		if c.NArg() > 0 {
			for _, tag := range c.Args() {
				err = deleteTag(tag)
				if err != nil {
					return err
				}
				fmt.Printf("deleted tag %v\n", tag)
				return nil
			}
		}

		tag := tags[len(tags)-1:][0]
		err = deleteTag(tag)
		if err != nil {
			return err
		}
		fmt.Printf("deleted tag %v\n", tag)

		return nil
	},
}

func deleteTag(tag string) error {
	cmd := fmt.Sprintf("git tag -d %v", tag)
	out, err := exec.Command("sh", "-c",
		cmd).CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return err
	}
	return nil
}
