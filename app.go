package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"moonshot/util"
	"os"
	"os/exec"

	"github.com/diskfs/go-diskfs/partition/gpt"
	"github.com/jaypipes/ghw"
	"github.com/wailsapp/wails/v3/pkg/application"
)

type AppService struct {
	app *application.App
}

func NewAppService(app *application.App) *AppService {
	return &AppService{
		app: app,
	}
}

type SourceFile struct {
	Path     string `json:"path"`
	Basename string `json:"basename"`
	Size     int64  `json:"size"`
	ValidGPT bool   `json:"validGPT"`
}

func (s *AppService) SelectFile(filepath *string) (*SourceFile, error) {
	if filepath == nil {
		fp, err := s.app.Dialog.OpenFile().
			SetTitle("Select Disk Image").
			AddFilter("Disk Images", "*.img;*.iso;*.raw").
			PromptForSingleSelection()
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

func (s *AppService) ListDrives() ([]Drive, error) {
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

func (s *AppService) FlashDrive(filePath string, driveName string, eject bool) error {
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

	var stderrBuf bytes.Buffer

	cmd.Stderr = &stderrBuf

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	scanner := bufio.NewScanner(stdout)

	for scanner.Scan() {
		s.app.Event.Emit("progress", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	if err := cmd.Wait(); err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			return errors.New(stderrBuf.String())
		}
		return err
	}

	return nil
}
