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
		bytes := TABLE.RowSlot(TABLE.RowNumber)
		copy(bytes[:], rowBytes[:])
	case frontend.SelectStatement:
		fmt.Println(statement.Content)
	}
}

func (t Table) RowSlot(rowNumber uint32) []byte {
	pageNumber := rowNumber / uint32(RowPerPage)
	rowOffset := rowNumber % uint32(RowPerPage)
	page := t.Pages[pageNumber]
	return page.Bytes[rowOffset:]
}
