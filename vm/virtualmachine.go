package vm

import (
	"fmt"
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
		whereToWrite := TABLE.RowSlot(TABLE.RowNumber)
		copy(whereToWrite[:], rowBytes[:])
		TABLE.RowNumber += 1
	case frontend.SelectStatement:
		fmt.Println(statement.Content)
		lastRowBytes := TABLE.RowSlot(TABLE.RowNumber - 1)
		var lastRow [295]byte
		copy(lastRow[:], lastRowBytes[:295])
		convertedRow := frontend.BytesToRow(lastRow)
		fmt.Printf("%v %v %v \n", convertedRow.Id, convertedRow.Username, convertedRow.Email)
	}
}
