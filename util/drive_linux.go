//go:build linux

package util

import (
	"os/exec"
)

func GetDrivePath(name string) string {
	return "/dev/" + name
}

func Eject(drivePath string) error {
	err := exec.Command("eject", drivePath).Run()
	if err != nil {
		return err
	}

	return nil
}

func PrepareDrive(drivePath string) error {
	return nil
}
