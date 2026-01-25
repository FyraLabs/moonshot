package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"moonshot/lib"
	"moonshot/util"
	"os"
)

type Message struct {
	Stage   string `json:"stage"`
	Written int    `json:"written"`
	End     bool   `json:"end"`
}

func flash() error {
	fs := flag.NewFlagSet("flash", flag.ContinueOnError)
	eject := fs.Bool("eject", false, "Eject drive after flashing")
	fs.Parse(os.Args[2:])

	filePath := fs.Arg(0)
	drivePath := fs.Arg(1)

	stage := "flash"
	ch := make(chan int)
	go func() {
		for n := range ch {
			msg := Message{Stage: stage, Written: n, End: false}
			bytes, err := json.Marshal(msg)
			if err != nil {
				os.Exit(1)
			}
			fmt.Println(string(bytes))
		}
	}()

	stat, err := os.Stat(filePath)
	if err != nil {
		return err
	}

	hash, err := lib.Flash(filePath, drivePath, ch)
	if err != nil {
		return err
	}

	stage = "verify"
	if ok, err := lib.Verify(hash, uint64(stat.Size()), drivePath, ch); !ok {
		return err
	}

	if *eject {
		if err := util.Eject(drivePath); err != nil {
			return err
		}
	}

	return nil
}
