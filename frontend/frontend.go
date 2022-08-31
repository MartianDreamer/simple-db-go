package frontend

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strings"
)

func PrepareStatement(command string) (result Statement, Ok bool) {
	if strings.HasPrefix(command, "insert") {
		return Statement{command, InsertStatement}, true
	} else if strings.HasPrefix(command, "select") {
		return Statement{command, SelectStatement}, true
	}
	return Statement{}, false
}

func PrepareInsertStatement(statement Statement) (result Row, err error) {
	if statement.StatementType != InsertStatement {
		return result, errors.New("the statement is not an insert statement")
	}
	var Username, Email string
	fmt.Sscanf(statement.Content, "insert %v %v %v", &result.Id, &Username, &Email)
	copy(result.Username[:], []byte(Username))
	copy(result.Email[:], []byte(Email))
	return result, nil
}

func RowToBytes(row Row) (rs [295]byte) {
	binary.LittleEndian.PutUint64(rs[:8], row.Id)
	copy(rs[8:40], row.Username[:])
	copy(rs[40:], row.Email[:])
	return rs
}

func BytesToRow(bytes [295]byte) (rs Row) {
	rs.Id = binary.LittleEndian.Uint64(bytes[:8])
	copy(rs.Username[:], bytes[8:40])
	copy(rs.Email[:], bytes[40:])
	return rs
}
