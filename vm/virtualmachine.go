package vm

import (
	"fmt"
	"simple-db-go/backend"
	"simple-db-go/frontend"
)

func ExecuteStatement(statement frontend.Statement) {
	switch statement.StatementType {
	case frontend.InsertStatement:
		insertRow, err := frontend.PrepareInsertStatement(statement)
		if err != nil {
			return
		}
		rowBytes := frontend.RowToBytes(insertRow)
		whereToWrite := backend.TABLE.RowSlot(backend.TABLE.RowNumber)
		copy(whereToWrite[:], rowBytes[:])
		backend.TABLE.RowNumber += 1
		fmt.Println("db > Executed.")
		fmt.Printf("db > (%v, %v, %v)\n", insertRow.Id, string(insertRow.Username[:]), string(insertRow.Email[:]))
	case frontend.SelectStatement:
		lastRowBytes := backend.TABLE.RowSlot(backend.TABLE.RowNumber - 1)
		var lastRow [295]byte
		copy(lastRow[:], lastRowBytes[:295])
		convertedRow := frontend.BytesToRow(lastRow)
		fmt.Println("db > Executed.")
		fmt.Printf("db > (%v %v %v)\n", convertedRow.Id, string(convertedRow.Username[:]), string(convertedRow.Email[:]))
	}
}
