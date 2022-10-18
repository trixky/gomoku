package main

import (
	"os"
	"fmt"
	"github.com/trixky/gomoku/server"
)
func main() {
	api_mode := len(os.Args) == 1

	if api_mode {
		server.Start()
	} else {
		debug.Start()
	}
}
