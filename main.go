package main

import (
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var cmdb []byte
	for {
		cmdb, _, _ = reader.ReadLine()
		cmds := string(cmdb)
		if isMetaCommand(cmds) {
			executeMetaCommand(cmds)
			continue
		}
		executeStatement(cmds)
	}
}
