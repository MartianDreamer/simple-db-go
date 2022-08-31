package main

import (
	"fmt"
	"simple-db-go/command"
	"simple-db-go/frontend"
	"simple-db-go/vm"
	"strings"
)

func isMetaCommand(command string) bool {
	return strings.HasPrefix(command, ".")
}

func executeMetaCommand(cmds string) {
	if executeCmd, ok := command.COMMAND_MAP[cmds]; ok {
		executeCmd()
	} else {
		fmt.Printf("%s is not recognized\n", cmds)
	}
}

func executeStatement(cmds string) {
	if statement, ok := frontend.PrepareStatement(cmds); ok {
		vm.ExecuteStatement(statement)
	} else {
		fmt.Printf("%s is not recognized\n", cmds)
	}
}
