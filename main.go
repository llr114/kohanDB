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
			fmt.Println(input_text)
		}
	}
}
