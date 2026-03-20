package main

import "fmt"

type ColumnDef struct {
	Name string
	ColType string
}

type CreateTableStmt struct {
	TableName string
	Columns []ColumnDef
}

func parse(tokens []Token) (*CreateTableStmt, error) {
	pos := 0

	if tokens[pos].Value != "CREATE" {
		return nil, fmt.Errorf("expected CREATE")
	}
	pos++

	if tokens[pos].Value != "TABLE" {
		return nil, fmt.Errorf("expected TABLE")
	}
	pos++

	tableName := tokens[pos].Value
	pos++

	if tokens[pos].Value != "(" {
		return nil, fmt.Errorf("expected (")
	}
	pos++

	var columns []ColumnDef
	for tokens[pos].Value != ")" {
		colName := tokens[pos].Value
		pos++
		colType := tokens[pos].Value
		pos++
		columns = append(columns, ColumnDef{Name: colName, ColType: colType})
		
		if tokens[pos].Value == "," {
			pos++
		}
	}

	if tokens[pos].Value != ")" {
		return nil, fmt.Errorf("expected )")
	}
	pos++

	return &CreateTableStmt{TableName: tableName, Columns: columns}, nil
}
