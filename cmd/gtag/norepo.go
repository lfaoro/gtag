package main

import (
	"github.com/pkg/errors"
	"os/exec"
	"strings"
)

func isGitRepo() error {
	out, err := exec.Command("sh", "-c",
		"git --no-pager tag -l").CombinedOutput()
	if err != nil {
		yes := strings.Contains(string(out), "not a git repository")
		if yes {
			return errors.Wrap(err, "not a git repository")
		}
		return err
	}
	return nil
}
