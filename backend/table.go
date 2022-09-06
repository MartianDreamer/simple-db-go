package backend

const (
	PageSize   uint16 = 4096
	MaxPage    uint8  = 100
	RowSize    uint16 = 295
	RowPerPage uint8  = 4096 / 295
)

type Page struct {
	Bytes [PageSize]byte
}

type Table struct {
	RowNumber uint32
	Pager     *Pager
}

func DbOpen(dbName string) (rs *Table) {
	pager := OpenPager(dbName)
	rowNumber := pager.fileLength / int64(RowSize)
	rs = &Table{uint32(rowNumber), pager}
	return rs
}

func (t *Table) DbClose() {
	pageCount := t.RowNumber / uint32(RowPerPage)
	for i := 0; i < int(pageCount); i++ {
		if t.Pager.Pages[i] == nil {
			continue
		}
		t.Pager.flushPage(uint8(i), RowPerPage)
		t.Pager.Pages[i] = nil
	}
	partialPageRow := t.RowNumber % uint32(RowPerPage)
	if partialPageRow > 0 {
		lastPage := pageCount
		if t.Pager.Pages[lastPage] != nil {
			t.Pager.flushPage(uint8(lastPage), uint8(partialPageRow))
			t.Pager.Pages[lastPage] = nil
		}
	}
	err := t.Pager.fileDescriptor.Close()
	if err != nil {
		panic("Failed to close db")
	}
	t = nil
}

func (t *Table) RowSlot(rowNumber uint32) []byte {
	return t.Pager.rowSlot(rowNumber)
}
