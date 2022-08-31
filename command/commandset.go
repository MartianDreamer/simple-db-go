package command

import (
	"fmt"
	"os"
	"simple-db-go/globconst"
)

const EXIT string = ".exit"

var COMMAND_MAP = map[string]func(args ...string){
	".exit": func(args ...string) {
		os.Exit(0)
	},
	".command": func(args ...string) {
		fmt.Println(".command")
	},
	".version": func(args ...string) {
		fmt.Println(globconst.Version)
	},
}
