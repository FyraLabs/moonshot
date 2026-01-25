//go:build linux

package util

import (
	"os/exec"
)

func RunAsRoot(program []string) (*exec.Cmd, func() error, error) {
	cmd := exec.Command("pkexec", program...)
	return cmd, func() error { return nil }, nil
}
