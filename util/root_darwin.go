//go:build darwin

package util

import (
	"fmt"
	"os/exec"
	"strings"
)

func RunAsRoot(command []string) *exec.Cmd {
	escaped := strings.ReplaceAll(strings.Join(command, " "), "\"", "\\\"")
	script := fmt.Sprintf("do shell script \"%s\" with administrator privileges", escaped)
	cmd := exec.Command("osascript", "-e", script)
	return cmd
}
