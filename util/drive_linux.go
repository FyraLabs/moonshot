//go:build linux

package util

import (
	"errors"
	"os/exec"
	"syscall"
)

func GetDrivePath(name string) string {
	return "/dev/" + name
}

func Eject(drivePath string) error {
	// Attempt to eject the drive, if that fails, try to unmount it
	error := exec.Command("eject", drivePath).Run()
	if errors.Is(error, &exec.ExitError{}) {
		return syscall.Unmount(drivePath, syscall.MNT_DETACH)
	} else if error != nil {
		return error
	}

	return nil
}
