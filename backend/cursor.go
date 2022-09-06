package backend

type Cursor struct {
	Table      *Table
	RowNumber  uint32
	EndOfTable bool
}

func TableStart(table *Table) (cur *Cursor) {
	return &Cursor{table, 0, table.RowNumber == 0}
}

func TableEnd(table *Table) (cur *Cursor) {
	return &Cursor{table, table.RowNumber, true}
}

func (cur *Cursor) GetValue() []byte {
	return cur.Table.RowSlot(cur.RowNumber)
}

func (cur *Cursor) Advance() {
	cur.RowNumber += 1
	if cur.RowNumber == cur.Table.RowNumber {
		cur.EndOfTable = true
	}
}
