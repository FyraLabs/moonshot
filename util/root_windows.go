//go:build windows

package util

import "os/exec"

func RunAsRoot(program []string) (*exec.Cmd, func() error, error) {
	return nil, func() error { return nil }, nil
}
