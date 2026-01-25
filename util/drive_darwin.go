//go:build darwin

package util

import (
	"errors"
	"os/exec"
)

func GetDrivePath(name string) string {
	return "/dev/r" + name
}

func Eject(drivePath string) error {
	err := exec.Command("diskutil", "eject", drivePath).Run()
	var exitErr *exec.ExitError
	if err != nil && errors.As(err, &exitErr) {
		err := exec.Command("diskutil", "eject", "force", drivePath).Run()
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}

func PrepareDrive(drivePath string) error {
	return nil
}
