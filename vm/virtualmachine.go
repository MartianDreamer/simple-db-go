package vm

import (
	"fmt"
	"simple-db-go/backend"
	"simple-db-go/frontend"
)

func ExecuteStatement(statement frontend.Statement) {
	switch statement.StatementType {
	case frontend.InsertStatement:
		if backend.TABLE == nil {
			fmt.Print("Use a databse\n")
			break
		}
		insertRow, err := frontend.PrepareInsertStatement(statement)
		if err != nil {
			break
		}
		cursor := backend.TableEnd(backend.TABLE)
		Serialize(insertRow, cursor)
		fmt.Println("db > Executed.")
		fmt.Printf("db > (%v, %v, %v)\n", insertRow.Id, string(insertRow.Username[:]), string(insertRow.Email[:]))
	case frontend.SelectStatement:
		if backend.TABLE == nil {
			fmt.Print("Use a databse\n")
			break
		}
		cursor := backend.TableStart(backend.TABLE)
		fmt.Println("db > Executed.")
		for !cursor.EndOfTable {
			curRow := cursor.GetValue()
			convertedRow := Deserialize(curRow)
			fmt.Printf("db > (%v %v %v)\n", convertedRow.Id, string(convertedRow.Username[:]), string(convertedRow.Email[:]))
			cursor.Advance()
		}
	case frontend.UseStatement:
		fileName, _ := frontend.PrepareUseStatement(statement)
		backend.TABLE = backend.DbOpen(fileName + ".db")
		fmt.Println("db > Executed.")
		fmt.Printf("db > use %v\n", fileName)
	}
}
