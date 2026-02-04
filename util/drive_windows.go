//go:build windows

package util

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"structs"
	"syscall"
	"unsafe"

	"github.com/ncw/directio"
	"golang.org/x/sys/windows"
)

const (
	FSCTL_LOCK_VOLUME                    = uint32(0x90018)
	FSCTL_DISMOUNT_VOLUME                = uint32(0x90020)
	IOCTL_STORAGE_EJECT_MEDIA            = uint32(0x2d4808)
	IOCTL_VOLUME_GET_VOLUME_DISK_EXTENTS = uint32(0x560000)
	FSCTL_ALLOW_EXTENDED_DASD_IO         = uint32(0x90083)
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

func OpenDriveForFlash(drivePath string) (*os.File, error) {
	driveNumber, err := strconv.Atoi(strings.TrimPrefix(strings.ToUpper(drivePath), "\\\\.\\PHYSICALDRIVE"))
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

	drive, err := directio.OpenFile(drivePath, os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	handle := windows.Handle(drive.Fd())

	volumePaths, err := listVolumes()
	if err != nil {
		return nil, err
	}

	for _, volumePath := range volumePaths {
		diskNumbers, err := findDisksForVolume(volumePath)
		if err != nil {
			return nil, err
		}

		if len(diskNumbers) == 0 || int(diskNumbers[0]) != driveNumber {
			continue
		} else if len(diskNumbers) > 1 {
			return nil, errors.New("Drive contains volume that spans multiple disks, refusing to continue")
		}

		handle, err := windows.Open(volumePath, windows.GENERIC_READ, windows.FILE_SHARE_READ|windows.FILE_SHARE_WRITE|windows.FILE_SHARE_DELETE)
		if err != nil {
			return nil, err
		}

		// May or may not work depending on the underlying filesystem, I think...
		_ = windows.DeviceIoControl(handle,
			FSCTL_ALLOW_EXTENDED_DASD_IO,
			nil, 0,
			nil, 0,
			nil,
			nil,
		)

		err = windows.DeviceIoControl(
			handle,
			FSCTL_LOCK_VOLUME,
			nil, 0,
			nil, 0,
			nil,
			nil,
		)
		if err != nil {
			windows.Close(handle)
			return nil, err
		}

		err = windows.DeviceIoControl(
			handle,
			FSCTL_DISMOUNT_VOLUME,
			nil, 0,
			nil, 0,
			nil,
			nil,
		)
		if err != nil {
			windows.Close(handle)
			return nil, err
		}
	}

	err = windows.DeviceIoControl(
		handle,
		FSCTL_LOCK_VOLUME,
		nil, 0,
		nil, 0,
		nil,
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
		nil,
		nil,
	)
	if err != nil {
		drive.Close()
		return nil, err
	}

	return drive, nil
}

func listVolumes() ([]string, error) {
	var volumeNames []string
	var volumeNameBuffer [windows.MAX_PATH]uint16

	handle, err := windows.FindFirstVolume(&volumeNameBuffer[0], uint32(len(volumeNameBuffer)))
	if err != nil {
		return nil, err
	}
	defer windows.FindVolumeClose(handle)

	for {
		volumeNames = append(volumeNames, strings.TrimSuffix(windows.UTF16ToString(volumeNameBuffer[:]), "\\"))

		err = windows.FindNextVolume(handle, &volumeNameBuffer[0], uint32(len(volumeNameBuffer)))
		if err != nil {
			if err == windows.ERROR_NO_MORE_FILES {
				break
			}
			return nil, err
		}
	}

	return volumeNames, nil
}

type DiskExtent struct {
	DiskNumber     uint32
	StartingOffset int64
	ExtentLength   int64
	_              structs.HostLayout
}

type VolumeDiskExtents struct {
	NumberOfDiskExtents uint32
	Extents             [1]DiskExtent
	_                   structs.HostLayout
}

func findDisksForVolume(volumeName string) ([]uint32, error) {
	handle, err := windows.Open(volumeName, windows.GENERIC_READ, windows.FILE_SHARE_READ|windows.FILE_SHARE_WRITE|windows.FILE_SHARE_DELETE)
	if err != nil {
		return nil, err
	}

	defer windows.CloseHandle(handle)

	buffer := make([]byte, unsafe.Sizeof(VolumeDiskExtents{}))
	volumeDiskExtents := (*VolumeDiskExtents)(unsafe.Pointer(&buffer[0]))

	if err := windows.DeviceIoControl(handle, IOCTL_VOLUME_GET_VOLUME_DISK_EXTENTS, nil, 0, &buffer[0], uint32(len(buffer)), nil, nil); err != nil {
		if err == windows.ERROR_MORE_DATA {
			buffer = make([]byte, unsafe.Sizeof(VolumeDiskExtents{})+unsafe.Sizeof(DiskExtent{})*(uintptr(volumeDiskExtents.NumberOfDiskExtents)-1))
			volumeDiskExtents = (*VolumeDiskExtents)(unsafe.Pointer(&buffer[0]))

			if err := windows.DeviceIoControl(handle, IOCTL_VOLUME_GET_VOLUME_DISK_EXTENTS, nil, 0, &buffer[0], uint32(len(buffer)), nil, nil); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	extents := unsafe.Slice(&volumeDiskExtents.Extents[0], volumeDiskExtents.NumberOfDiskExtents)

	var diskNumbers []uint32
	for _, extent := range extents {
		diskNumbers = append(diskNumbers, extent.DiskNumber)
	}

	return diskNumbers, nil
}
