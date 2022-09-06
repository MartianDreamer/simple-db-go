package backend

import "os"

type Pager struct {
	fileDescriptor *os.File
	fileLength     int64
	Pages          [MaxPage]*Page
}

func OpenPager(fileName string) (rs *Pager) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic("Can not create page")
	}
	fileInfo, err := file.Stat()
	if err != nil {
		panic("Failed to get file info")
	}
	rs = &Pager{file, fileInfo.Size(), [MaxPage]*Page{}}
	return rs
}

func (pager *Pager) rowSlot(rowNumber uint32) []byte {
	pageNumber := rowNumber / uint32(RowPerPage)
	rowOffset := rowNumber % uint32(RowPerPage)
	byteOffset := rowOffset * uint32(RowSize)
	if pager.Pages[pageNumber] == nil {
		pager.Pages[pageNumber] = &Page{[PageSize]byte{}}
	}
	page := pager.Pages[pageNumber]
	return page.Bytes[byteOffset:]
}
