package lib

import (
	"io"
	"os"

	"github.com/ncw/directio"
	"github.com/schollz/progressbar/v3"
)

func Flash(filePath string, drivePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		return err
	}

	drive, err := directio.OpenFile(drivePath, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer drive.Close()

	block := directio.AlignedBlock(directio.BlockSize * 256)

	bar := progressbar.DefaultBytes(fileStat.Size(), "flashing")

	for {
		_, err := io.ReadFull(file, block)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			break
		}
		if err != nil {
			return err
		}

		_, err = io.MultiWriter(drive, bar).Write(block)
		if err != nil {
			return err
		}
	}

	return nil

}
