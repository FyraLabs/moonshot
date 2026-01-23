//go:build darwin

package util

func GetDrivePath(name string) string {
	return "/dev/r" + name
}

func Eject(drivePath string) error {
	return nil
}
