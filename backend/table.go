package backend

const (
	PageSize   uint16 = 4096
	MaxPage    uint8  = 100
	RowSize    uint16 = 291
	RowPerPage uint8  = 4096 / 291
)

type Page struct {
	Bytes [PageSize]byte
}

type Table struct {
	RowNumber uint32
	Pager     *Pager
}

func NewTable(tableName string) (rs *Table) {
	rs.RowNumber = 0
	rs.Pager = OpenPager(tableName)
	return rs
}

func (t *Table) RowSlot(rowNumber uint32) []byte {
	return t.Pager.rowSlot(rowNumber)
}
