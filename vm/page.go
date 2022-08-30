package vm

const (
	PageSize         = 4096
	MaxPage          = 100
	RowSize          = 291
	RowPerPage uint8 = 4096 / 291
)

type Page struct {
	Bytes [PageSize]byte
}

type Table struct {
	RowNumber uint32
	Pages     [MaxPage]Page
}

var TABLE = Table{0, [MaxPage]Page{}}
