package lib

import (
	"errors"
	"hash"
	"hash/crc64"
	"io"
	"moonshot/util"
	"os"
	"syscall"

	"github.com/ncw/directio"
)

type ProgressWriter struct {
	Channel chan int
}

func (p *ProgressWriter) Write(b []byte) (int, error) {
	p.Channel <- len(b)
	return len(b), nil
}

func Flash(filePath string, drivePath string, progressCh chan int) (hash.Hash64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	drive, err := directio.OpenFile(drivePath, os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}
	defer drive.Close()

	if err := util.PrepareDrive(drive); err != nil {
		return nil, err
	}

	block := directio.AlignedBlock(directio.BlockSize * 256)

	hash := crc64.New(crc64.MakeTable(crc64.ISO))
	progress := &ProgressWriter{Channel: progressCh}
	auxWriter := io.MultiWriter(hash, progress)

	for {
		readN, err := io.ReadFull(file, block)
		if err == io.EOF {
			break
		} else if err != nil && err != io.ErrUnexpectedEOF {
			return nil, err
		}

		if readN < len(block) {
			clear(block[readN:])
		}

		// If we read a partial block, we should round down to the minimum required block size to handle an edge case where the end of the disk is smaller than our default buffer size
		writeSize := readN
		if readN%directio.BlockSize != 0 {
			writeSize = ((readN / directio.BlockSize) + 1) * directio.BlockSize
		}

		// We must write the full block to the drive, whereas the hash and progress only get the valid data
		_, err = drive.Write(block[:writeSize])
		if err != nil {
			return nil, err
		}

		_, err = auxWriter.Write(block[:readN])
		if err != nil {
			return nil, err
		}
	}

	// On macOS, sync on a rdisk can return ENOTTY, which should be ignored
	err = drive.Sync()
	if !errors.Is(err, syscall.ENOTTY) && err != nil {
		return nil, err
	}

	return hash, nil
}

func Verify(hash hash.Hash64, size uint64, drivePath string, progressCh chan int) (bool, error) {
	drive, err := directio.OpenFile(drivePath, os.O_RDONLY, 0666)
	if err != nil {
		return false, err
	}
	defer drive.Close()

	block := directio.AlignedBlock(directio.BlockSize * 256)

	diskHash := crc64.New(crc64.MakeTable(crc64.ISO))
	progress := &ProgressWriter{Channel: progressCh}
	writer := io.MultiWriter(diskHash, progress)

	readCount := uint64(0)

	for {
		readN, err := io.ReadFull(drive, block)
		if err == io.EOF {
			break
		} else if err != nil && err != io.ErrUnexpectedEOF {
			return false, err
		}

		readCount += uint64(readN)
		if readCount > size {
			readN -= int(readCount - size)
		}

		_, err = writer.Write(block[:readN])
		if err != nil {
			return false, err
		}

		if readCount >= size {
			break
		}
	}

	if diskHash.Sum64() != hash.Sum64() {
		return false, nil
	}

	return true, nil
}
