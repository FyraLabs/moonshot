//go:build windows

package util

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"

	"github.com/ncw/directio"
	"golang.org/x/sys/windows"
)

const (
	FSCTL_LOCK_VOLUME         = uint32(0x90018)
	FSCTL_DISMOUNT_VOLUME     = uint32(0x90020)
	IOCTL_STORAGE_EJECT_MEDIA = uint32(0x2d4808)
)

func GetDrivePath(name string) string {
	return name
}

func Eject(drivePath string) error {
	driveFile, err := os.Open(drivePath)
	if err != nil {
		return err
	}
	defer driveFile.Close()

	handle := windows.Handle(driveFile.Fd())

	var bytesReturned uint32
	err = windows.DeviceIoControl(
		handle,
		IOCTL_STORAGE_EJECT_MEDIA,
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

func OpenDriveForFlash(driveFile string) (*os.File, error) {
	// I'm not 100% sure if this is correct, but cleaning the disk using diskpart, then locking and dismounting the drive seems to allow the write calls to succeed.
	// Thanks Gemini for saving me from shoveling through Microsoft's shitty documentation.

	driveNumber, err := strconv.Atoi(strings.TrimPrefix(strings.ToUpper(driveFile), `\\.\PHYSICALDRIVE`))
	if err != nil {
		return nil, err
	}

	cmd := exec.Command("diskpart")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: 0x08000000,
	}

	cmd.Stdin = bytes.NewBufferString(fmt.Sprintf("select disk %d\nclean\nrescan", driveNumber))

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	drive, err := directio.OpenFile(driveFile, os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}

	handle := windows.Handle(drive.Fd())

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
		drive.Close()
		return nil, err
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
		drive.Close()
		return nil, err
	}

	return drive, nil
}
