package main

import (
	"encoding/json"
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
	filePath := os.Args[2]
	drivePath := os.Args[3]

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

	if err := util.Eject(drivePath); err != nil {
		return err
	}

	return nil
}
