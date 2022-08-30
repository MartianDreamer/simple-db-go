package main

import (
	"fmt"
	"simple-db-go/frontend"
	"simple-db-go/metacommand"
	"simple-db-go/vm"
	"strings"
)

func isMetaCommand(command string) bool {
	return strings.HasPrefix(command, ".")
}

func executeMetaCommand(cmds string) {
	if executeCmd, ok := metacommand.COMMAND_MAP[cmds]; ok {
		executeCmd()
	} else {
		fmt.Printf("%s is not recognized\n", cmds)
	}
}

func executeStatement(cmds string) {
	byteCode := frontend.CompileByteCode(cmds)
	vm.ExecuteStatement(byteCode)
}
