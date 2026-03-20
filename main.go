package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)

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
				fmt.Println(stmt)
			}
		}
	}
}
