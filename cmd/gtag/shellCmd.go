package main

import "os/exec"

func shellCmd(cmd string) (output string, err error) {
	out, err := exec.Command("sh", "-c",
		cmd).CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
