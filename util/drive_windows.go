//go:build windows

package util

func GetDrivePath(name string) string {
	return name
}

func Eject(drivePath string) error {
	return nil
}
