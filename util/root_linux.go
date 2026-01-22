//go:build linux

package util

import "os/exec"

func RunAsRoot(program string) *exec.Cmd {
	return nil
}
