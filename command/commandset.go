package command

import (
	"fmt"
	"os"
)

const EXIT string = ".exit"

var COMMAND_MAP = map[string]func(){
	".exit": func() {
		os.Exit(0)
	},
	".command": func() {
		fmt.Println(".command")
	},
}
