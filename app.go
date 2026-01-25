package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"moonshot/util"
	"os"

	"github.com/diskfs/go-diskfs/partition/gpt"
	"github.com/jaypipes/ghw"
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

type SourceFile struct {
	Path     string `json:"path"`
	Basename string `json:"basename"`
	Size     int64  `json:"size"`
	ValidGPT bool   `json:"validGPT"`
}

func (a *App) SelectFile(filepath *string) (*SourceFile, error) {
	if filepath == nil {
		fp, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
			Filters: []runtime.FileFilter{
				{
					DisplayName: "Disk Images (*.img, *.img.*, *.iso, *.raw)",
					Pattern:     "*.img;*.iso;*.raw",
				},
			},
		})
		if err != nil {
			return nil, err
		}
		filepath = &fp
	}

	file, err := os.Open(*filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	validGpt := false

	table, err := gpt.Read(file, 512, 512)
	if err != nil {
		log.Println(err)
	}

	if table != nil {
		table, err := gpt.Read(file, 512, 512)
		if err != nil {
			log.Println(err)
		} else if table != nil {
			validGpt = true
		}
	}

	return &SourceFile{
		Path:     *filepath,
		Basename: stat.Name(),
		Size:     stat.Size(),
		ValidGPT: validGpt,
	}, nil
}

type Drive struct {
	Name      string `json:"name"`
	Model     string `json:"model"`
	Capacity  uint64 `json:"capacity"`
	Path      string `json:"path"`
	Removable bool   `json:"removable"`
}

func (a *App) ListDrives() ([]Drive, error) {
	block, err := ghw.Block()
	if err != nil {
		fmt.Printf("Error getting block storage info: %v", err)
	}

	drives := []Drive{}

	for _, disk := range block.Disks {
		drives = append(drives, Drive{
			Name:      disk.Name,
			Model:     disk.Model,
			Capacity:  disk.SizeBytes,
			Path:      disk.BusPath,
			Removable: disk.IsRemovable,
		})
	}

	return drives, nil
}

func (a *App) FlashDrive(filePath string, driveName string, eject bool) error {
	exe, err := os.Executable()
	if err != nil {
		return err
	}

	command := []string{exe, "flash"}
	if eject {
		command = append(command, "--eject")
	}

	command = append(command, filePath, util.GetDrivePath(driveName))

	cmd, cleanup, err := util.RunAsRoot(command)
	defer cleanup()
	if err != nil {
		return err
	}
	cmd.Stderr = os.Stderr

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	scanner := bufio.NewScanner(stdout)

	for scanner.Scan() {
		runtime.EventsEmit(a.ctx, "progress", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}
