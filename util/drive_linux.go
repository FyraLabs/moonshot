//go:build linux

package util

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/ncw/directio"
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

func OpenDriveForFlash(drivePath string) (*os.File, error) {
	matches, err := filepath.Glob(drivePath + "*")
	if err != nil {
		return nil, err
	}

	for _, match := range matches {
		_ = exec.Command("umount", match).Run()
	}

	drive, err := directio.OpenFile(drivePath, os.O_WRONLY|os.O_EXCL, 0666)
	if err != nil {
		return nil, err
	}

	return drive, nil
}
