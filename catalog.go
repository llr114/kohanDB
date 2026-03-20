package main

import (
	"fmt"
	"encoding/json"
	"os"
)

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

func (c *Catalog) SaveToFile(path string) error {
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func (c *Catalog) LoadFromFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, c)
}
