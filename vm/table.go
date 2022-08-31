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

const (
	PageSize         = 4096
	MaxPage          = 100
	RowSize          = 291
	RowPerPage uint8 = 4096 / 291
)

type Page struct {
	Bytes *[PageSize]byte
}

type Table struct {
	RowNumber uint32
	Pages     *[MaxPage]Page
}

func (t Table) RowSlot(rowNumber uint32) *[]byte {
	pageNumber := rowNumber / uint32(RowPerPage)
	rowOffset := rowNumber % uint32(RowPerPage)
	page := t.Pages[pageNumber]
	rs := page.Bytes[rowOffset:]
	return &rs
}
