package main

import (
	"encoding/json"
	"fmt"
	"moonshot/lib"
	"os"
)

type Message struct {
	Stage   string `json:"stage"`
	Written int    `json:"written"`
	End     bool   `json:"end"`
}

func flash() error {
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

	stat, err := os.Stat(os.Args[2])
	if err != nil {
		return err
	}

	hash, err := lib.Flash(os.Args[2], os.Args[3], ch)
	if err != nil {
		return err
	}

	stage = "verify"
	if ok, err := lib.Verify(hash, uint64(stat.Size()), os.Args[3], ch); !ok {
		return err
	}

	return nil
}
