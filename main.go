package main

import (
	"fmt"
	"github.com/english-learning/cmd"
	"runtime/debug"
)

func main() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("main error recover %s\n", p)
			fmt.Printf("main error recover %s\n", string(debug.Stack()))
		}
	}()

	cmd.Run()
}
