package frontend

type StatementType int8

const (
	InsertStatement StatementType = iota
	SelectStatement
)

type Statement struct {
	Content       string
	StatementType StatementType
}
