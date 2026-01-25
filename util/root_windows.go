//go:build windows

package util

import "os/exec"

func RunAsRoot(program []string) (*exec.Cmd, func() error, error) {
	cmd := exec.Command(program[0], program[1:]...)
	return cmd, func() error { return nil }, nil
}
