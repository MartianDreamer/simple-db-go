package main

import (
	"bufio"
	"fmt"
	"os"
	"simple-db-go/frontend"
	"simple-db-go/metacommand"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var cmdb []byte
	for {
		cmdb, _, _ = reader.ReadLine()
		cmds := string(cmdb)
		if isMetaCommand(cmds) {
			if executeCmd, ok := metacommand.COMMAND_MAP[cmds]; ok {
				executeCmd()
			} else {
				fmt.Printf("%s is not recognized\n", cmds)
			}
			continue
		}
		bytes := frontend.CompileByteCode(cmds)
		fmt.Println(bytes)
	}
}

func isMetaCommand(command string) bool {
	return strings.HasPrefix(command, ".")
}
