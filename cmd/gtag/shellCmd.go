package main

import (
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

func shellCmd(cmd string) (output string, err error) {
	out, err := exec.Command("sh", "-c",
		cmd).CombinedOutput()
	if err != nil {
		return "", errors.Wrap(err, string(out))
	}
	return string(out), nil
}

func isGitRepo() error {
	out, err := shellCmd("git --no-pager tag -l")
	if err != nil {
		yes := strings.Contains(out, "not a git repository")
		if yes {
			return errors.Wrap(err, "not a git repository")
		}
		return errors.Wrap(err, out)
	}
	return nil
}
