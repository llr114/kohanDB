package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	catalog := &Catalog{Tables: make(map[string]*CreateTableStmt)}

	for {
		fmt.Print("kohanDB> ")
		scanner.Scan()
		input_text := scanner.Text()
		if input_text == ".exit" {
			break
		} else {
			tokens := tokenize(input_text)
			stmt, err := parse(tokens)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				err = catalog.CreateTable(stmt)
				if err != nil {
					fmt.Println("Error:", err)
				} else {
					fmt.Println("Table created.")
				}
			}
		}
	}
}
