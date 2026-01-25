//go:build windows

package util

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"golang.org/x/sys/windows"
)

func GetDrivePath(name string) string {
	return name
}

func Eject(drivePath string) error {
	return nil
}

const (
	FSCTL_LOCK_VOLUME     = uint32(0x90018)
	FSCTL_DISMOUNT_VOLUME = uint32(0x90020)
)

func PrepareDrive(driveFile *os.File) error {
	driveNumber, err := strconv.Atoi(strings.TrimPrefix(strings.ToUpper(driveFile.Name()), `\\.\PHYSICALDRIVE`))
	if err != nil {
		return err
	}

	cmd := exec.Command("diskpart")
	cmd.Stdin = bytes.NewBufferString(fmt.Sprintf("select disk %d\nclean\nrescan", driveNumber))
	if err := cmd.Run(); err != nil {
		return err
	}

	handle := windows.Handle(driveFile.Fd())

	var bytesReturned uint32
	err = windows.DeviceIoControl(
		handle,
		FSCTL_LOCK_VOLUME,
		nil, 0,
		nil, 0,
		&bytesReturned,
		nil,
	)
	if err != nil {
		return err
	}

	err = windows.DeviceIoControl(
		handle,
		FSCTL_DISMOUNT_VOLUME,
		nil, 0,
		nil, 0,
		&bytesReturned,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}
