//go:build windows

package util

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func GetDrivePath(name string) string {
	return name
}

func Eject(drivePath string) error {
	return nil
}

func PrepareDrive(drivePath string) error {
	driveNumber, err := strconv.Atoi(strings.TrimPrefix(drivePath, `\\.\PhysicalDrive`))
	if err != nil {
		return err
	}

	cmd := exec.Command("diskpart")
	cmd.Stdin = bytes.NewBufferString(fmt.Sprintf("select disk %d\nclean\nrescan", driveNumber))
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
