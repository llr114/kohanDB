package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	catalog := &Catalog{Tables: make(map[string]*CreateTableStmt)}
	catalog.LoadFromFile("kohan.db")

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
				switch s := stmt.(type) {
					case *CreateTableStmt:
						err = catalog.CreateTable(s)
						if err != nil {
							fmt.Println("Error:", err)
						} else {
							catalog.SaveToFile("kohan.db")
							fmt.Println("Table created.")
						}
					case *InsertStmt:
						fmt.Println(s)
				}
			}
		}
	}
}
