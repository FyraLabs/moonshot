//go:build darwin

package util

import (
	"fmt"
	"os/exec"
	"strings"
)

func RunAsRoot(command []string) *exec.Cmd {
	script := fmt.Sprintf(`do shell script "%s" with administrator privileges`, strings.Join(command, " "))
	return exec.Command("osascript", "-e", script)
}
