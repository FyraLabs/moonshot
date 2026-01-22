package lib

import (
	"io"
	"os"

	"github.com/ncw/directio"
)

func Flash(filePath string, drivePath string, progressCh chan int) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	drive, err := directio.OpenFile(drivePath, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer drive.Close()

	block := directio.AlignedBlock(directio.BlockSize * 256)

	for {
		_, err := io.ReadFull(file, block)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			break
		}
		if err != nil {
			return err
		}

		n, err := drive.Write(block)
		if err != nil {
			return err
		}

		progressCh <- n
	}

	return nil
}
