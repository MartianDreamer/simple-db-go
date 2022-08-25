package main

import (
	"bufio"
	"fmt"
	"os"
	"simple-db-go/command"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var cmdb []byte
	for {
		cmdb, _, _ = reader.ReadLine()
		cmds := string(cmdb)
		if cmd, ok := command.COMMAND_MAP[cmds]; ok {
			cmd()
		} else {
			fmt.Printf("%s is not recognized", cmds)
		}
	}
}
