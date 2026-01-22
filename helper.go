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

	if err := lib.Flash(os.Args[2], os.Args[3], ch); err != nil {
		println("Error:", err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
