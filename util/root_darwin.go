//go:build darwin

package util

import (
	_ "embed"
	"os"
	"os/exec"
)

//go:embed askpass.js
var askpassScript []byte

func RunAsRoot(command []string) (*exec.Cmd, func() error, error) {
	cleanup := func() error { return nil }
	scriptFile, err := os.CreateTemp("", "askpass")
	if err != nil {
		return nil, cleanup, err
	}
	cleanup = func() error {
		return os.Remove(scriptFile.Name())
	}

	_, err = scriptFile.Write(askpassScript)
	if err != nil {
		return nil, cleanup, err
	}

	if err := scriptFile.Chmod(0700); err != nil {
		return nil, cleanup, err
	}

	cmd := exec.Command("sudo", "-A")
	cmd.Args = append(cmd.Args, command...)
	cmd.Env = append(os.Environ(), "SUDO_ASKPASS="+scriptFile.Name())

	return cmd, cleanup, nil
}
