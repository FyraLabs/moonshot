package main

import (
	"context"
	"fmt"
	"moonshot/util"
	"os"

	"github.com/jaypipes/ghw"
	"github.com/jaypipes/ghw/pkg/block"
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

func (a *App) FlashDrive(filePath string, drivePath string) error {
	exe, err := os.Executable()
	if err != nil {
		return err
	}

	cmdStr := fmt.Sprintf("sudo '%s' flash %s", exe, fmt.Sprintf("\"%s\" \"%s\"", filePath, drivePath))
	println(cmdStr)
	cmd := util.RunAsRoot(cmdStr)

	b, err := cmd.CombinedOutput()
	if err != nil {
		println(string(b))
		println("Error:", err.Error())
		return err
	}
	println(string(b))

	// stdout, err := cmd.StdoutPipe()
	// if err != nil {
	// 	return err
	// }

	// if err := cmd.Start(); err != nil {
	// 	return err
	// }

	// go func() {
	// 	io.Copy(os.Stdout, stdout)
	// 	// println("owo")
	// 	// scanner := bufio.NewScanner(stdout)

	// 	// for scanner.Scan() {
	// 	// 	println(scanner.Text())
	// 	// 	runtime.EventsEmit(a.ctx, "progress", scanner.Text())
	// 	// }

	// 	// if err := scanner.Err(); err != nil {
	// 	// 	panic(err)
	// 	// }
	// }()

	// if err := cmd.Wait(); err != nil {
	// 	return err
	// }

	return nil
}
