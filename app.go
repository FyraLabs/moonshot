package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/jaypipes/ghw"
	"github.com/jaypipes/ghw/pkg/block"
	"github.com/ncw/directio"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

type SourceFile struct {
	Path     string `json:"path"`
	Basename string `json:"basename"`
	Size     int64  `json:"size"`
}

func (a *App) SelectFile(filepath *string) (*SourceFile, error) {
	if filepath == nil {
		fp, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{})
		if err != nil {
			return nil, err
		}
		filepath = &fp
	}

	stat, err := os.Stat(*filepath)
	if err != nil {
		return nil, err
	}

	return &SourceFile{
		Path:     *filepath,
		Basename: stat.Name(),
		Size:     stat.Size(),
	}, nil
}

type Drive struct {
	Name      string `json:"name"`
	Capacity  uint64 `json:"capacity"`
	Path      string `json:"path"`
	Removable bool   `json:"removable"`
}

func (a *App) ListDrives() ([]*block.Disk, error) {
	block, err := ghw.Block()
	if err != nil {
		fmt.Printf("Error getting block storage info: %v", err)
	}

	drives := []Drive{}

	for _, disk := range block.Disks {
		drives = append(drives, Drive{
			Name:      disk.Name,
			Capacity:  disk.SizeBytes,
			Path:      disk.BusPath,
			Removable: disk.IsRemovable,
		})
	}

	return block.Disks, nil
}

func (a *App) FlashDrive(drivePath string, filePath string) error {
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

	block := directio.AlignedBlock(directio.BlockSize)

	for {
		println("owo")
		_, err := io.ReadFull(file, block)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		_, err = drive.Write(block)
		if err != nil {
			return err
		}
	}

	return nil
}
