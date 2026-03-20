package main

import "fmt"

type Catalog struct {
	Tables map[string]*CreateTableStmt
}

func (c *Catalog) CreateTable(stmt *CreateTableStmt) error {
	if _, exists := c.Tables[stmt.TableName]; exists {
		return fmt.Errorf("table %s already exists", stmt.TableName)
	}
	c.Tables[stmt.TableName] = stmt
	return nil
}
