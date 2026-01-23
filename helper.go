package main

import (
	"encoding/json"
	"fmt"
	"moonshot/lib"
	"os"
)

type Message struct {
	Written int  `json:"written"`
	End     bool `json:"end"`
}

func flash() {
	ch := make(chan int)
	go func() {
		for n := range ch {
			msg := Message{Written: n, End: false}
			bytes, err := json.Marshal(msg)
			if err != nil {
				println("Error:", err.Error())
				os.Exit(1)
			}
			fmt.Println(string(bytes))
		}
	}()

	stat, err := os.Stat(os.Args[2])
	if err != nil {
		println("Error:", err.Error())
		os.Exit(1)
	}

	hash, err := lib.Flash(os.Args[2], os.Args[3], ch)
	if err != nil {
		println("Error:", err.Error())
		os.Exit(1)
	}

	if ok, err := lib.Verify(hash, uint64(stat.Size()), os.Args[3], ch); !ok {
		if err != nil {
			println("Error:", err.Error())
		} else {
			println("Hash mismatch")
		}
		os.Exit(1)
	}

	os.Exit(0)
}
