package backend

import (
	"fmt"
	"io"
	"os"
)

type Pager struct {
	fileDescriptor *os.File
	fileLength     int64
	Pages          [MaxPage]*Page
}

func OpenPager(fileName string) (rs *Pager) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)
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
	page := pager.getPage(pageNumber)
	return page.Bytes[byteOffset:]
}

func (pager *Pager) getPage(pageNumber uint32) (rs *Page) {
	if pageNumber > uint32(MaxPage) {
		panic("Try to get out of bound page")
	}
	if pager.Pages[pageNumber] == nil {
		allocatedPage := &Page{[PageSize]byte{}}
		savedPageNumber := pager.fileLength / int64(PageSize)
		if pager.fileLength%int64(PageSize) != 0 {
			savedPageNumber += 1
		}
		if pageNumber < uint32(savedPageNumber) {
			_, err := pager.fileDescriptor.ReadAt(allocatedPage.Bytes[:], int64(pageNumber)*int64(PageSize))
			if err != nil && err != io.EOF {
				panic(fmt.Sprintf("Error reading : %v\n", err))
			}
		}
		pager.Pages[pageNumber] = allocatedPage
	}
	return pager.Pages[pageNumber]
}

func (pager *Pager) flushPage(pageIndex uint8, rowCount uint8) {
	_, err := pager.fileDescriptor.WriteAt(pager.Pages[pageIndex].Bytes[:uint32(rowCount)*uint32(RowSize)], int64(pageIndex)*int64(PageSize))
	if err != nil {
		panic(fmt.Sprintf("Error writing %v\n", err))
	}
}
