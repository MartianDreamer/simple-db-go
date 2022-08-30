package frontend

import "strings"

func PrepareStatement(command string) (result Statement, Ok bool) {
	if strings.HasPrefix(command, "insert") {
		return Statement{command, InsertStatement}, true
	} else if strings.HasPrefix(command, "select") {
		return Statement{command, SelectStatement}, true
	}
	return Statement{}, false
}
