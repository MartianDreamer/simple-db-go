package vm

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
