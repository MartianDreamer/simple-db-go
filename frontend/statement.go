package frontend

type StatementType int8

const (
	InsertStatement StatementType = iota
	SelectStatement
	UseStatement
)

type Statement struct {
	Content       string
	StatementType StatementType
}

type Row struct {
	Id       uint64
	Username [32]byte
	Email    [255]byte
}
