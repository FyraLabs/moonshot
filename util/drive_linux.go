//go:build linux

package util

import (
	"os"
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

func PrepareDrive(driveFile *os.File) error {
	return nil
}
