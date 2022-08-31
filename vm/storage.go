package vm

var TABLE = initTable()

func initTable() *Table {
	rs := Table{}
	rs.RowNumber = 0
	rs.Pages = (*[MaxPage]Page)(make([]Page, MaxPage))
	for i := 0; i < MaxPage; i += 1 {
		rs.Pages[i].Bytes = (*[PageSize]byte)(make([]byte, PageSize))
	}
	return &rs
}
