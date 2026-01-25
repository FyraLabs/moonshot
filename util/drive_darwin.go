//go:build darwin

package util

import (
	"errors"
	"os"
	"os/exec"

	"github.com/ncw/directio"
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

func OpenDriveForFlash(drivePath string) (*os.File, error) {
	cmd := exec.Command("diskutil", "unmountDisk", drivePath)
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	drive, err := directio.OpenFile(drivePath, os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}

	return drive, err
}
