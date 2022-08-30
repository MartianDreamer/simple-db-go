package vm

import (
	"fmt"
	"simple-db-go/frontend"
)

func ExecuteStatement(statement frontend.Statement) {
	switch statement.StatementType {
	case frontend.InsertStatement:
		fmt.Println(statement.Content)
	case frontend.SelectStatement:
		fmt.Println(statement.Content)
	}
}
