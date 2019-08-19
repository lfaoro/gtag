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
	Usage: "delete the last created tag, pass --all to delete all tags at once.\n" +
		"providing a tag as argument, will delete that specific tag",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "all, a",
			Usage: "removes all tags",
		}, cli.BoolFlag{
			Name:  "remote, r",
			Usage: "removes tag from remote",
		},
	},
	Action: func(c *cli.Context) error {
		out, err := shellCmd("git --no-pager tag -l")
		if err != nil {
			return err
		}
		s := strings.TrimSpace(out)
		tags := strings.Split(s, "\n")

		if c.Bool("all") {
			input := bufio.NewReader(os.Stdin)
			fmt.Print("You're about to delete all tags, sure? ")

			b, err := input.ReadString('\n')
			exitIfError(err)
			if strings.Contains(b, "n") {
				fmt.Println("Deletion cancelled.")
				return nil
			}

			for _, tag := range tags {
				err := deleteTag(tag, false)
				if err != nil {
					return err
				}

				fmt.Printf("deleted tag %v\n", tag)
			}
			return nil
		}

		if c.NArg() > 0 {
			for _, tag := range c.Args() {
				err = deleteTag(tag, false)
				if err != nil {
					return err
				}
				fmt.Printf("deleted tag %v\n", tag)
				return nil
			}
		}

		tag := tags[len(tags)-1:][0]
		err = deleteTag(tag, false)
		if err != nil {
			return err
		}
		fmt.Printf("deleted tag %v\n", tag)

		return nil
	},
}

func deleteTag(tag string, remote bool) error {
	cmd := fmt.Sprintf("git tag -d %v", tag)
	out, err := exec.Command("sh", "-c",
		cmd).CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return err
	}
	return nil
}
