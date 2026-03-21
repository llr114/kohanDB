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

type InsertStmt struct {
	TableName string
	Columns []string
	Values []string
}

func parse(tokens []Token) (interface{}, error) {
	if tokens[0].Value == "CREATE" {
		return parseCreateTable(tokens)
	}
	if tokens[0].Value == "INSERT" {
		return parseInsert(tokens)
	}
	return nil, fmt.Errorf("unknown statement: %s", tokens[0].Value)
}

func parseCreateTable(tokens []Token) (*CreateTableStmt, error) {
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
	for pos < len(tokens) && tokens[pos].Value != ")" {
		colName := tokens[pos].Value
		pos++
		colType := tokens[pos].Value
		pos++
		columns = append(columns, ColumnDef{Name: colName, ColType: colType})
		
		if tokens[pos].Value == "," {
			pos++
		}
	}
	pos++

	return &CreateTableStmt{TableName: tableName, Columns: columns}, nil
}

func parseInsert(tokens []Token) (*InsertStmt, error) {
	pos := 0

	if tokens[pos].Value != "INSERT" {
		return nil, fmt.Errorf("expected INSERT")
	}
	pos++

	if tokens[pos].Value != "INTO" {
		return nil, fmt.Errorf("expected INTO")
	}
	pos++

	tableName := tokens[pos].Value
	pos++

	if tokens[pos].Value != "(" {
		return nil, fmt.Errorf("expected (")
	}
	pos++

	var columns []string
	for pos < len(tokens) && tokens[pos].Value != ")" {
		columns = append(columns, tokens[pos].Value)
		pos++
		if tokens[pos].Value == "," {
			pos++
		}
	}
	pos++

	if tokens[pos].Value != "VALUES" {
		return nil, fmt.Errorf("expected VALUES")
	}
	pos++
	
	if tokens[pos].Value != "(" {
		return nil, fmt.Errorf("expected (")
	}
	pos++

	var values []string
	for pos < len(tokens) && tokens[pos].Value != ")" {
		values = append(values, tokens[pos].Value)
		pos++
		if tokens[pos].Value == "," {
			pos++
		}
	}
	pos++

	return &InsertStmt{TableName: tableName, Columns: columns, Values: values}, nil
}
