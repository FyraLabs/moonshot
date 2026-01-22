package main

import (
	"fmt"
	"moonshot/lib"
	"os"
)

func main() {
	if err := lib.Flash(os.Args[1], os.Args[2]); err != nil {
		fmt.Println("Error flashing drive: " + err.Error())
	}
}
