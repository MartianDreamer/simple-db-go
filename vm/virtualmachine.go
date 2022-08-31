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
		fmt.Println("db > Executed.")
		fmt.Printf("db > (%v, %v, %v)\n", insertRow.Id, string(insertRow.Username[:]), string(insertRow.Email[:]))
	case frontend.SelectStatement:
		lastRowBytes := TABLE.RowSlot(TABLE.RowNumber - 1)
		var lastRow [295]byte
		copy(lastRow[:], lastRowBytes[:295])
		convertedRow := frontend.BytesToRow(lastRow)
		fmt.Printf("%v %v %v \n", convertedRow.Id, string(convertedRow.Username[:]), string(convertedRow.Email[:]))
	}
}
